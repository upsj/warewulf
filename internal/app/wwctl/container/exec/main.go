// +build linux

package exec

import (
	"fmt"
	"github.com/hpcng/warewulf/internal/pkg/container"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"syscall"
)

func CobraRunE(cmd *cobra.Command, args []string) error {

	containerName := args[0]

	if container.ValidSource(containerName) == false {
		wwlog.Printf(wwlog.ERROR, "Unknown Warewulf container: %s\n", containerName)
		os.Exit(1)
	}

	c := exec.Command("/proc/self/exe", append([]string{"container", "exec", "__child"}, args...)...)

	//c := exec.Command("/bin/sh")
	c.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	if err := c.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	return nil
}
