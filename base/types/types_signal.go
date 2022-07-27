package types

import (
	"fmt"
	"strings"
	"syscall"
)

type Signal int

func (s Signal) Int() int { return int(s) }

func (s Signal) String() string { return shorts[s] }

func (s Signal) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *Signal) UnmarshalText(data []byte) error {
	if sig, ok := signals[strings.ToUpper(string(data))]; ok {
		*s = sig
		return nil
	}
	return fmt.Errorf("unknown signal")
}

func (s Signal) Error() string { return asErrors[s] }

const (
	SIGHUP    = Signal(syscall.SIGHUP)
	SIGINT    = Signal(syscall.SIGINT)
	SIGQUIT   = Signal(syscall.SIGQUIT)
	SIGILL    = Signal(syscall.SIGILL)
	SIGTRAP   = Signal(syscall.SIGTRAP)
	SIGABRT   = Signal(syscall.SIGABRT)
	SIGEMT    = Signal(syscall.SIGEMT)
	SIGFPE    = Signal(syscall.SIGFPE)
	SIGKILL   = Signal(syscall.SIGKILL)
	SIGBUS    = Signal(syscall.SIGBUS)
	SIGSEGV   = Signal(syscall.SIGSEGV)
	SIGSYS    = Signal(syscall.SIGSYS)
	SIGPIPE   = Signal(syscall.SIGPIPE)
	SIGALRM   = Signal(syscall.SIGALRM)
	SIGTERM   = Signal(syscall.SIGTERM)
	SIGURG    = Signal(syscall.SIGURG)
	SIGSTOP   = Signal(syscall.SIGSTOP)
	SIGTSTP   = Signal(syscall.SIGTSTP)
	SIGCONT   = Signal(syscall.SIGCONT)
	SIGCHLD   = Signal(syscall.SIGCHLD)
	SIGTTIN   = Signal(syscall.SIGTTIN)
	SIGTTOU   = Signal(syscall.SIGTTOU)
	SIGIO     = Signal(syscall.SIGIO)
	SIGXCPU   = Signal(syscall.SIGXCPU)
	SIGXFSZ   = Signal(syscall.SIGXFSZ)
	SIGVTALRM = Signal(syscall.SIGVTALRM)
	SIGPROF   = Signal(syscall.SIGPROF)
	SIGWINCH  = Signal(syscall.SIGWINCH)
	SIGINFO   = Signal(syscall.SIGINFO)
	SIGUSR1   = Signal(syscall.SIGUSR1)
	SIGUSR2   = Signal(syscall.SIGUSR2)
)

var asErrors = [...]string{
	SIGHUP:    "hangup",
	SIGINT:    "interrupt",
	SIGQUIT:   "quit",
	SIGILL:    "illegal instruction",
	SIGTRAP:   "trace/BPT trap",
	SIGABRT:   "abort trap",
	SIGEMT:    "EMT trap",
	SIGFPE:    "floating point exception",
	SIGKILL:   "killed",
	SIGBUS:    "bus error",
	SIGSEGV:   "segmentation fault",
	SIGSYS:    "bad system call",
	SIGPIPE:   "broken pipe",
	SIGALRM:   "alarm clock",
	SIGTERM:   "terminated",
	SIGURG:    "urgent I/O condition",
	SIGSTOP:   "suspended (signal)",
	SIGTSTP:   "suspended",
	SIGCONT:   "continued",
	SIGCHLD:   "child exited",
	SIGTTIN:   "stopped (tty input)",
	SIGTTOU:   "stopped (tty output)",
	SIGIO:     "I/O possible",
	SIGXCPU:   "cputime limit exceeded",
	SIGXFSZ:   "filesize limit exceeded",
	SIGVTALRM: "virtual timer expired",
	SIGPROF:   "profiling timer expired",
	SIGWINCH:  "window size changes",
	SIGINFO:   "information request",
	SIGUSR1:   "user defined signal 1",
	SIGUSR2:   "user defined signal 2",
}

var shorts = [...]string{
	SIGHUP:    "HUP",
	SIGINT:    "INT",
	SIGQUIT:   "QUIT",
	SIGILL:    "ILL",
	SIGTRAP:   "TRAP",
	SIGABRT:   "ABRT",
	SIGEMT:    "EMT",
	SIGFPE:    "FPE",
	SIGKILL:   "KILL",
	SIGBUS:    "BUS",
	SIGSEGV:   "SEGV",
	SIGSYS:    "SYS",
	SIGPIPE:   "PIPE",
	SIGALRM:   "ALRM",
	SIGTERM:   "TERM",
	SIGURG:    "URG",
	SIGSTOP:   "STOP",
	SIGTSTP:   "TSTP",
	SIGCONT:   "CONT",
	SIGCHLD:   "CHLD",
	SIGTTIN:   "TTIN",
	SIGTTOU:   "TTOU",
	SIGIO:     "IO",
	SIGXCPU:   "XCPU",
	SIGXFSZ:   "XFSZ",
	SIGVTALRM: "VTALARM",
	SIGPROF:   "PROF",
	SIGWINCH:  "WINCH",
	SIGINFO:   "INFO",
	SIGUSR1:   "USR1",
	SIGUSR2:   "USR2",
}

var signals = map[string]Signal{}

func init() {
	for sig, short := range shorts {
		signals[short] = Signal(sig)
	}
}
