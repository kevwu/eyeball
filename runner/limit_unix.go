// +build !windows

package runner

import (
	"syscall"
)

func initLimit() {
	var rLimit syscall.Rlimit
	rLimit.Max = 10000
	rLimit.Cur = 10000
	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		mainLog("Error setting RLIMIT_NOFILE. Try running as root.")
	}
}
