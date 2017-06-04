package osutil

import (
	"path/filepath"
	"runtime"
	"os"
	"strings"
	"fmt"
	"errors"
	"github.com/Sirupsen/logrus"
)

func GetDefaultConfigDir(appName string) string {
	switch runtime.GOOS {
	case "windows":
		appName = strings.Title(appName)
		if p := os.Getenv("LOCALAPPDATA"); p != "" {
			return filepath.Join(p, appName)
		}
		return filepath.Join(os.Getenv("APPDATA"), appName)

	case "darwin":
		appName = strings.Title(appName)
		dir, err := ExpandTilde("~/Library/Application Support/"+appName)
		if err != nil {
			logrus.Fatalln(err)
		}
		return dir

	default:
		appName = strings.ToLower(appName)
		if xdgCfg := os.Getenv("XDG_CONFIG_HOME"); xdgCfg != "" {
			return filepath.Join(xdgCfg, appName)
		}
		dir, err := ExpandTilde("~/.config/"+appName)
		if err != nil {
			logrus.Fatalln(err)
		}
		return dir
	}
}

func GetHomeDir() (string, error) {
	var home string
	switch runtime.GOOS {
	case "windows":
		home = filepath.Join(os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"))
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	default:
		home = os.Getenv("HOME")
	}

	if home == "" {
		return "", errors.New("No home directory found (platform specific env should be set)")
	}
	return home, nil
}

func ExpandTilde(path string) (string, error) {
	if path == "~" {
		return GetHomeDir()
	}
	if strings.HasPrefix(path, fmt.Sprintf("~%c", os.PathSeparator)) {
		homeDir, err := GetHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(homeDir, path[2:]), nil
	}
	return filepath.FromSlash(path), nil
}
