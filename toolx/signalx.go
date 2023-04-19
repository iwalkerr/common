package toolx

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/logx"
)

// kill -9 强杀不能捕捉
func WaitExit(exit func()) {
	for i := range NewShutdownSignal() {
		switch i {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			logx.Infof("receive exit gignal %s,exit...", i.String())
			exit()
			os.Exit(0)
		}
	}
}

func NewShutdownSignal() chan os.Signal {
	c := make(chan os.Signal)
	// SIGHUP: terminal closed
	// SIGINT: Ctrl+C
	// SIGTERM: program exit
	// SIGQUIT: Ctrl+/
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	return c
}
