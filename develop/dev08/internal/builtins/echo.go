package builtins

import (
	"log"
	"sync"
	"syscall"
	"wb_l2/develop/dev08/internal"
)

type Echo struct {
	internal.Command
}

func (e *Echo) Run(wg *sync.WaitGroup) error {
	var (
		wasFlagN bool
		x        = 1
	)
	defer wg.Done()
	wasFlagN = isFlagN(e.GetArgs())
	if wasFlagN {
		x++
	}
	if _, err := syscall.Write(e.GetWriter(), []byte(e.GetArgs()[x])); err != nil {
		log.Fatal(err)
	}
	if !wasFlagN {
		if _, err := syscall.Write(e.GetWriter(), []byte{'\n'}); err != nil {
			log.Fatal(err)
		}
	}
	e.CloseFds()

	return nil
}

func isFlagN(args []string) bool {
	var t, b bool
	for _, v := range args {
		for _, r := range v {
			if r == '-' {
				t = true
			} else if r == 'n' && t {
				b = true
			} else if !t || !b {
				t, b = false, false
			}
		}
	}
	return t && b
}
