package builtins

import (
	"errors"
	"fmt"
	ps "github.com/mitchellh/go-ps"
	"sync"
	"syscall"
	"wb_l2/develop/dev08/internal"
)

type Ps struct {
	internal.Command
}

func (p Ps) Run(wg *sync.WaitGroup) (err error) {
	var processes []ps.Process
	if processes, err = ps.Processes(); err != nil {
		return errors.New("ps: " + err.Error())
	}

	if _, err = syscall.Write(p.GetWriter(), []byte(fmt.Sprintf("pid\tproc\n"))); err != nil {
		return errors.New("ps: " + err.Error())
	}
	for _, proc := range processes {
		if _, err = syscall.Write(p.GetWriter(), []byte(fmt.Sprintf("%d\t%s\n", proc.Pid(), proc.Executable()))); err != nil {
			return errors.New("ps: " + err.Error())
		}
	}
	wg.Done()
	p.CloseFds()
	return nil
}
