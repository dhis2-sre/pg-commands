package pgcommands

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
)

type ExecOptions struct {
	StreamPrint       bool
	StreamDestination io.Writer
}

func streamExecOutput(out io.ReadCloser, options ExecOptions) (string, error) {
	output := ""
	reader := bufio.NewReader(out)
	line, err := reader.ReadString('\n')
	log.Println("line")
	log.Println(line)
	output += line
	for err == nil {
		if options.StreamPrint {
			//nolint: staticcheck
			fmt.Printf(line)
		}
		line, err = reader.ReadString('\n')
		output += line
	}

	return output, nil
}

func streamExecOutput2(out io.ReadCloser, options ExecOptions) (string, error) {
	output := ""
	reader := bufio.NewReader(out)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				return output, nil
			}
			return output, err
		}

		if options.StreamPrint {
			_, err = fmt.Fprintln(options.StreamDestination, line)
			if err != nil {
				return output, err
			}
		}

		output += line
	}
}

func CommandExist(command string) bool {
	_, err := exec.LookPath(command)

	return err == nil
}
