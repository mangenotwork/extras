package command

import (
	"context"
	"log"
	"os/exec"
	"syscall"
	"time"
)

func Cmd(timeOut time.Duration, cmdName string, arg ...string) (bool, error) {

	ctxt, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	cmd := exec.CommandContext(ctxt, cmdName, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return false, err
	}

	cmd.Stderr = cmd.Stdout
	if err = cmd.Start(); err != nil {
		return false, err
	}

	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		log.Println(string(tmp)) // 从管道中实时获取输出并打印到终端
		if err != nil {
			break
		}
	}

	if err = cmd.Wait(); err != nil {
		return false, err
	}

	log.Println(cmd.Process.Pid)
	err = syscall.Kill(cmd.Process.Pid, syscall.SIGKILL)
	log.Println("kill err = ", err)
	return true, nil
}

type SysCmd struct {
	Command string
	Pid int
	OutString chan string
	End chan bool
}

func NewCmd(command string) *SysCmd {
	return &SysCmd{
		Command: command,
		OutString: make(chan string, 0),
		End: make(chan bool, 0),
	}
}

func (syscmd *SysCmd) Run(arg ...string) {
	go func() {
		cmd := exec.Command(syscmd.Command, arg...)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			//return false, err
			log.Println(false, err)
			syscmd.End <- true
		}

		cmd.Stderr = cmd.Stdout
		if err = cmd.Start(); err != nil {
			//return false, err
			log.Println(false, err)
			syscmd.End <- true
		}

		for {
			tmp := make([]byte, 1024)
			_, err := stdout.Read(tmp)
			//fmt.Print(string(tmp)) // 从管道中实时获取输出并打印到终端
			if err != nil {
				break
			}
			syscmd.OutString <- string(tmp)
		}

		if err = cmd.Wait(); err != nil {
			//return false, err
			log.Println(false, err)
			syscmd.End <- true
		}

		syscmd.Pid = cmd.Process.Pid

		//log.Println(cmd.Process.Pid)
		//err = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		//log.Println("kill err = ", err)
		//return true, nil
		log.Println(false, err)
		syscmd.End <- true
	}()
}

func (syscmd *SysCmd) Kill() error {
	close(syscmd.OutString)
	close(syscmd.End)
	return syscall.Kill(syscmd.Pid, syscall.SIGKILL)
}

func (syscmd *SysCmd) ListenerOut() {
	go func() {
		for{
			select {
			case str := <-syscmd.OutString:
				log.Println("str = ", str)
			}
		}
	}()
}

func (syscmd *SysCmd) Listener() {
	for{
		select {
		case end := <-syscmd.End:
			log.Println("end = ", end)
		}
	}
}