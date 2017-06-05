// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package discovery

import (
	"errors"
	"fmt"
	"net"
	"time"
	"golang.org/x/net/ipv6"
	"github.com/thejerf/suture"
	"github.com/Sirupsen/logrus"
)

type Multicast struct {
	*suture.Supervisor
	inbox  chan []byte
	outbox chan recvPacket
	mr     *multicastReader
	mw     *multicastWriter
}

func NewMulticast(addr string) *Multicast {
	m := &Multicast{
		Supervisor: suture.New("multicastBeacon", suture.Spec{
			// Don't retry too frenetically: an error to open a socket or
			// whatever is usually something that is either permanent or takes
			// a while to get solved...
			FailureThreshold: 2,
			FailureBackoff:   60*time.Second,
			// Only log restarts in debug mode.
			Log: func(line string) {
				logrus.Debugln(line)
			},
		}),
		inbox:  make(chan []byte),
		outbox: make(chan recvPacket, 16),
	}

	m.mr = &multicastReader{
		addr:   addr,
		outbox: m.outbox,
		stop:   make(chan struct{}),
	}
	m.Add(m.mr)

	m.mw = &multicastWriter{
		addr:  addr,
		inbox: m.inbox,
		stop:  make(chan struct{}),
	}
	m.Add(m.mw)

	return m
}

func (m *Multicast) Send(data []byte) {
	m.inbox <- data
}

func (m *Multicast) Recv() ([]byte, net.Addr) {
	recv := <-m.outbox
	return recv.data, recv.src
}

func (m *Multicast) Error() error {
	if err := m.mr.Error(); err != nil {
		return err
	}
	return m.mw.Error()
}

type multicastWriter struct {
	addr  string
	inbox <-chan []byte
	errorHolder
	stop chan struct{}
}

func (w *multicastWriter) Serve() {
	logrus.Debugln(w, "starting")
	defer logrus.Debugln(w, "stopped")

	gaddr, err := net.ResolveUDPAddr("udp6", w.addr)
	if err != nil {
		logrus.Debugln(err)
		w.setError(err)
		return
	}

	conn, err := net.ListenPacket("udp6", ":0")
	if err != nil {
		logrus.Debugln(err)
		w.setError(err)
		return
	}

	pconn := ipv6.NewPacketConn(conn)

	wcm := &ipv6.ControlMessage{
		HopLimit: 1,
	}

	for bs := range w.inbox {
		intfs, err := net.Interfaces()
		if err != nil {
			logrus.Debugln(err)
			w.setError(err)
			return
		}

		success := 0
		for _, intf := range intfs {
			wcm.IfIndex = intf.Index
			pconn.SetWriteDeadline(time.Now().Add(time.Second))
			_, err = pconn.WriteTo(bs, wcm, gaddr)
			pconn.SetWriteDeadline(time.Time{})

			if err != nil {
				logrus.Debugln("%s when writing to %v on %s", err, gaddr, intf.Name)
				w.setError(err)
				continue
			}

			logrus.Debugf("sent %d bytes to %v on %s", len(bs), gaddr, intf.Name)

			success++
		}

		if success > 0 {
			w.setError(nil)
		} else {
			logrus.Debugln(err)
			w.setError(err)
		}
	}
}

func (w *multicastWriter) Stop() {
	close(w.stop)
}

func (w *multicastWriter) String() string {
	return fmt.Sprintf("multicastWriter@%p", w)
}

type multicastReader struct {
	addr   string
	outbox chan<- recvPacket
	errorHolder
	stop chan struct{}
}

func (r *multicastReader) Serve() {
	logrus.Debugln(r, "starting")
	defer logrus.Debugln(r, "stopped")

	gaddr, err := net.ResolveUDPAddr("udp6", r.addr)
	if err != nil {
		logrus.Debugln(err)
		r.setError(err)
		return
	}

	conn, err := net.ListenPacket("udp6", r.addr)
	if err != nil {
		logrus.Debugln(err)
		r.setError(err)
		return
	}

	intfs, err := net.Interfaces()
	if err != nil {
		logrus.Debugln(err)
		r.setError(err)
		return
	}

	pconn := ipv6.NewPacketConn(conn)
	joined := 0
	for _, intf := range intfs {
		err := pconn.JoinGroup(&intf, &net.UDPAddr{IP: gaddr.IP})
		if err != nil {
			logrus.Debugf("%s when joining group %v on %s", err, gaddr.IP, intf.Name)
		} else {
			logrus.Debugf("joined group %v on %s", gaddr.IP, intf.Name)
		}
		joined++
	}

	if joined == 0 {
		logrus.Debugln("no multicast interfaces available")
		r.setError(errors.New("no multicast interfaces available"))
		return
	}

	bs := make([]byte, 65536)
	for {
		n, _, addr, err := pconn.ReadFrom(bs)
		if err != nil {
			logrus.Debugln(err)
			r.setError(err)
			continue
		}
		logrus.Debugf("recvPacket %d bytes from %s", n, addr)

		c := make([]byte, n)
		copy(c, bs)
		select {
		case r.outbox <- recvPacket{c, addr}:
		default:
			logrus.Debugln("UDP6 packet dropped")
		}
	}
}

func (r *multicastReader) Stop() {
	close(r.stop)
}

func (r *multicastReader) String() string {
	return fmt.Sprintf("multicastReader@%p", r)
}
