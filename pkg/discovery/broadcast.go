// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package discovery

import (
	"fmt"
	"net"
	"time"
	"github.com/Sirupsen/logrus"
	"github.com/syncthing/syncthing/lib/sync"
	"github.com/thejerf/suture"
)

type Broadcast struct {
	*suture.Supervisor
	port   int
	inbox  chan []byte
	outbox chan recvPacket
	br     *broadcastReader
	bw     *broadcastWriter
}

func NewBroadcast(port int) *Broadcast {
	b := &Broadcast{
		Supervisor: suture.New("broadcastDiscovery", suture.Spec{
			// Don't retry too frenetically: an error to open a socket or
			// whatever is usually something that is either permanent or takes
			// a while to get solved...
			FailureThreshold: 2,
			FailureBackoff: 60*time.Second,
			// Only log restarts in debug mode.
			Log: func(line string) {
				logrus.Debugln(line)
			},
		}),
		port:   port,
		inbox:  make(chan []byte),
		outbox: make(chan recvPacket, 16),
	}

	b.br = &broadcastReader{
		port:   port,
		outbox: b.outbox,
		mut:    sync.NewMutex(),
	}
	b.Add(b.br)
	b.bw = &broadcastWriter{
		port:    port,
		inbox:   b.inbox,
		connMut: sync.NewMutex(),
	}
	b.Add(b.bw)

	return b
}

func (b *Broadcast) Send(data []byte) {
	b.inbox <- data
}

func (b *Broadcast) Recv() ([]byte, net.Addr) {
	recv := <-b.outbox
	return recv.data, recv.src
}

func (b *Broadcast) Error() error {
	if err := b.br.Error(); err != nil {
		return err
	}
	return b.bw.Error()
}

type broadcastWriter struct {
	port    int
	inbox   chan []byte
	conn    *net.UDPConn
	connMut sync.Mutex
	errorHolder
}

func (w *broadcastWriter) Serve() {
	logrus.Debugln(w, "starting")
	defer logrus.Debugln(w, "stopped")

	conn, err := net.ListenUDP("udp4", nil)
	if err != nil {
		logrus.Debugln(err)
		w.setError(err)
		return
	}
	defer conn.Close()

	w.connMut.Lock()
	w.conn = conn
	w.connMut.Unlock()

	for bs := range w.inbox {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			logrus.Debugln(err)
			w.setError(err)
			continue
		}

		var dsts []net.IP
		for _, addr := range addrs {
			iaddr, ok := addr.(*net.IPNet)
			if ok && iaddr.IP.IsGlobalUnicast() && iaddr.IP.To4() != nil {
				baddr := bcast(iaddr)
				dsts = append(dsts, baddr.IP)
			}
		}

		if len(dsts) == 0 {
			// Fall back to the general IPv4 broadcast address
			dsts = append(dsts, net.IP{255, 255, 255, 255})
		}

		logrus.Debugln("addresses:", dsts)

		success := 0
		for _, ip := range dsts {
			dst := &net.UDPAddr{IP: ip, Port: w.port}

			conn.SetWriteDeadline(time.Now().Add(time.Second))
			_, err := conn.WriteTo(bs, dst)
			conn.SetWriteDeadline(time.Time{})

			if err, ok := err.(net.Error); ok {
				if err.Timeout() {
					// Write timeouts should not happen. We treat it as a fatal
					// error on the socket.
					logrus.Debugln(err)
					w.setError(err)
					return
				} else if err.Temporary() {
					// A transient error. Lets hope for better luck in the future.
					logrus.Debugln(err)
					continue
				}
			}

			if err != nil {
				// Some other error that we don't expect. Bail and retry.
				logrus.Debugln(err)
				w.setError(err)
				return
			}

			logrus.Debugf("sent %d bytes to %s", len(bs), dst)
			success++
		}

		if success > 0 {
			w.setError(nil)
		}
	}
}

func (w *broadcastWriter) Stop() {
	w.connMut.Lock()
	if w.conn != nil {
		w.conn.Close()
	}
	w.connMut.Unlock()
}

func (w *broadcastWriter) String() string {
	return fmt.Sprintf("broadcastWriter@%p", w)
}

type broadcastReader struct {
	port   int
	outbox chan recvPacket
	conn   *net.UDPConn
	mut    sync.Mutex
	errorHolder
}

func (r *broadcastReader) Serve() {
	logrus.Debugln(r, "starting")
	defer logrus.Debugln(r, "stopped")

	conn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: r.port})
	if err != nil {
		logrus.Debugln(err)
		r.setError(err)
		return
	}
	defer conn.Close()

	r.mut.Lock()
	r.conn = conn
	r.mut.Unlock()

	bs := make([]byte, 65536)
	for {
		n, addr, err := conn.ReadFrom(bs)
		if err != nil {
			logrus.Debugln(err)
			r.setError(err)
			return
		}

		r.setError(nil)

		logrus.Debugf("recvPacket %d bytes from %s", n, addr)

		c := make([]byte, n)
		copy(c, bs)
		select {
		case r.outbox <- recvPacket{c, addr}:
		default:
			logrus.Debugln("UDP packet dropped")
		}
	}
}

func (r *broadcastReader) Stop() {
	r.mut.Lock()
	if r.conn != nil {
		r.conn.Close()
	}
	r.mut.Unlock()
}

func (r *broadcastReader) String() string {
	return fmt.Sprintf("broadcastReader@%p", r)
}

func bcast(ip *net.IPNet) *net.IPNet {
	var bc = &net.IPNet{}
	bc.IP = make([]byte, len(ip.IP))
	copy(bc.IP, ip.IP)
	bc.Mask = ip.Mask

	offset := len(bc.IP) - len(bc.Mask)
	for i := range bc.IP {
		if i-offset >= 0 {
			bc.IP[i] = ip.IP[i] | ^ip.Mask[i-offset]
		}
	}
	return bc
}
