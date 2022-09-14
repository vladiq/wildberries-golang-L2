package builtins

import (
	"os"
	"sync"
	"syscall"
	"wb_l2/develop/dev08/internal"
)

type Pwd struct {
	internal.Command
}

func (p Pwd) Run(wg *sync.WaitGroup) error {
	if wd, err := os.Getwd(); err != nil {
		return err
	} else if _, err = syscall.Write(p.GetWriter(), []byte(wd+"\n")); err != nil {
		return err
	}
	p.CloseFds()
	wg.Done()
	return nil
}
