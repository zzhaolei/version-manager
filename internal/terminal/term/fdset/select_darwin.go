//go:build darwin

package fdset

import "syscall"

func Select(nfd int, r *syscall.FdSet, w *syscall.FdSet, e *syscall.FdSet, timeout *syscall.Timeval) error {
	return syscall.Select(nfd, r, w, e, timeout)
}
