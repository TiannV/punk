// @CreateTime: Oct 31, 2019 6:03 PM
// @Author: tianwei
// @Contact: tianwei@langnal.com
// @Last Modified By: tianwei
// @Last Modified Time: Oct 31, 2019 6:30 PM
// @Description: Modify Here, Please

package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			os.Stdout.Write(d)
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

// CmdExec 执行shell命令
func CmdExec(command string) error {
	cmd := exec.Command("/bin/sh", "-c", command)
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	cmd.Start()
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
	}()
	go func() {
		stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)
	}()
	err := cmd.Wait()
	if err != nil {
		color.Red("cmd.Run() failed with %s\n", err)
		return err
	}
	if errStdout != nil || errStderr != nil {
		color.Red("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	return errStderr
}
