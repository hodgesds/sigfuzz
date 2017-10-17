package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

/*
   SIGABRT   = Signal(0x6)
   SIGALRM   = Signal(0xe)
   SIGBUS    = Signal(0x7)
   SIGCHLD   = Signal(0x11)
   SIGCLD    = Signal(0x11)
   SIGCONT   = Signal(0x12)
   SIGFPE    = Signal(0x8)
   SIGHUP    = Signal(0x1)
   SIGILL    = Signal(0x4)
   SIGINT    = Signal(0x2)
   SIGIO     = Signal(0x1d)
   SIGIOT    = Signal(0x6)
   SIGKILL   = Signal(0x9)
   SIGPIPE   = Signal(0xd)
   SIGPOLL   = Signal(0x1d)
   SIGPROF   = Signal(0x1b)
   SIGPWR    = Signal(0x1e)
   SIGQUIT   = Signal(0x3)
   SIGSEGV   = Signal(0xb)
   SIGSTKFLT = Signal(0x10)
   SIGSTOP   = Signal(0x13)
   SIGSYS    = Signal(0x1f)
   SIGTERM   = Signal(0xf)
   SIGTRAP   = Signal(0x5)
   SIGTSTP   = Signal(0x14)
   SIGTTIN   = Signal(0x15)
   SIGTTOU   = Signal(0x16)
   SIGUNUSED = Signal(0x1f)
   SIGURG    = Signal(0x17)
   SIGUSR1   = Signal(0xa)
   SIGUSR2   = Signal(0xc)
   SIGVTALRM = Signal(0x1a)
   SIGWINCH  = Signal(0x1c)
   SIGXCPU   = Signal(0x18)
   SIGXFSZ   = Signal(0x19)

*/

func getSignals(sigs []string) ([]os.Signal, error) {
	signals := []os.Signal{}
	for _, sigStr := range sigs {
		sig, err := getSignal(sigStr)
		if err != nil {
			return signals, err
		}
		signals = append(signals, sig)
	}

	return signals, nil
}

func getSignal(sig string) (os.Signal, error) {
	switch strings.ToUpper(sig) {
	case syscall.SIGABRT.String():
		return syscall.SIGABRT, nil
	case syscall.SIGALRM.String():
		return syscall.SIGALRM, nil
	case syscall.SIGBUS.String():
		return syscall.SIGBUS, nil
	case syscall.SIGCHLD.String():
		return syscall.SIGCHLD, nil
	case syscall.SIGCLD.String():
		return syscall.SIGCLD, nil
	case syscall.SIGCONT.String():
		return syscall.SIGCONT, nil
	case syscall.SIGFPE.String():
		return syscall.SIGFPE, nil
	case syscall.SIGHUP.String():
		return syscall.SIGHUP, nil
	case syscall.SIGILL.String():
		return syscall.SIGILL, nil
	case syscall.SIGINT.String():
		return syscall.SIGINT, nil
	case syscall.SIGIO.String():
		return syscall.SIGIO, nil
	case syscall.SIGIOT.String():
		return syscall.SIGIOT, nil
	case syscall.SIGKILL.String():
		return syscall.SIGKILL, nil
	case syscall.SIGPIPE.String():
		return syscall.SIGPIPE, nil
	case syscall.SIGPOLL.String():
		return syscall.SIGPOLL, nil
	case syscall.SIGPROF.String():
		return syscall.SIGPROF, nil
	case syscall.SIGPWR.String():
		return syscall.SIGPWR, nil
	case syscall.SIGQUIT.String():
		return syscall.SIGQUIT, nil
	case syscall.SIGSEGV.String():
		return syscall.SIGSEGV, nil
	case syscall.SIGSTKFLT.String():
		return syscall.SIGSTKFLT, nil
	case syscall.SIGSTOP.String():
		return syscall.SIGSTOP, nil
	case syscall.SIGSYS.String():
		return syscall.SIGSYS, nil
	case syscall.SIGTERM.String():
		return syscall.SIGTERM, nil
	case syscall.SIGTRAP.String():
		return syscall.SIGTRAP, nil
	case syscall.SIGTSTP.String():
		return syscall.SIGTSTP, nil
	case syscall.SIGTTIN.String():
		return syscall.SIGTTIN, nil
	case syscall.SIGTTOU.String():
		return syscall.SIGTTOU, nil
	case syscall.SIGUNUSED.String():
		return syscall.SIGUNUSED, nil
	case "SIGURG":
		return syscall.SIGURG, nil
	case syscall.SIGUSR1.String():
		return syscall.SIGUSR1, nil
	case syscall.SIGUSR2.String():
		return syscall.SIGUSR2, nil
	case syscall.SIGVTALRM.String():
		return syscall.SIGVTALRM, nil
	case syscall.SIGWINCH.String():
		return syscall.SIGWINCH, nil
	case syscall.SIGXCPU.String():
		return syscall.SIGXCPU, nil
	case syscall.SIGXFSZ.String():
		return syscall.SIGXFSZ, nil
	default:
		sigInt, err := strconv.Atoi(sig)
		if err != nil {
			return nil, fmt.Errorf("unknown signal: %s", sig)
		}

		return syscall.Signal(sigInt), nil
	}
}

func getProcesses(pidStrings []string) ([]*os.Process, error) {
	processes := []*os.Process{}
	for _, pidString := range pidStrings {
		pid, err := strconv.Atoi(pidString)
		if err != nil {
			return processes, fmt.Errorf("pid is not a int, got: %s", pidString)
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			return processes, err
		}

		processes = append(processes, process)
	}

	return processes, nil
}
