// +build linux

package logrus

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TCGETS
