package cmdutil

import (
	"bytes"
	"os/exec"
	"runtime"
)

// ExecCommand 执行cmd命令
// use shell /bin/bash -c to execute command
func ExecCommand(command string) (stdout, stderr string, err error) {
	var out bytes.Buffer
	var errout bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd")
	}
	cmd.Stdout = &out
	cmd.Stderr = &errout
	err = cmd.Run()

	if err != nil {
		stderr = errout.String()
	}
	stdout = out.String()
	return
}

// ExecCmd 命令和返回输出。
//
// Usage:
//
//	ExecCmd("ls", []string{"-al"})
func ExecCmd(binName string, args []string, workDir ...string) (string, error) {
	// create a new Cmd instance
	cmd := exec.Command(binName, args...)
	if len(workDir) > 0 {
		cmd.Dir = workDir[0]
	}

	bs, err := cmd.Output()
	return string(bs), err
}
