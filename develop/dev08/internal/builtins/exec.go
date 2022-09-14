package builtins

import (
	"log"
	"os"
	"sync"
	"syscall"
	"wb_l2/develop/dev08/internal"
)

type Exec struct {
	internal.Command
}

func (e *Exec) Run(wg *sync.WaitGroup) error {
	e.CloseFds()
	if len(e.GetArgs()) > 1 {
		if err := syscall.Exec(e.GetArgs()[1], e.GetArgs()[1:], os.Environ()); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
