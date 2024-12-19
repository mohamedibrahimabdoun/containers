package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("/bin/bash")
	// The statements below refer to the input, output and error streams of the process created (cmd)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//setting an environment variable
	cmd.Env = []string{"name=Mohamed"}
	// the command below creates a UTS namespace for the process
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET |
			syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}
	//  UidMappings and GidMappings. We have a field called ContainerID ,
	// which we are setting to 0. This means we are mapping the uid and gid 0 within the container to
	//the uid and gid of the user who launched the process.

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running the /bin/bash command - %s\n", err)
		os.Exit(1)
	}
}
