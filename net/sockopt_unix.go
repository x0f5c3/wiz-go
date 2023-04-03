//go:build !windows

package net

import "syscall"

func setSockOpt(fd uintptr, level, opt, value int) error {
	return syscall.SetsockoptInt(int(fd), level, opt, value)
}
