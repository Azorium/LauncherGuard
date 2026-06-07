package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type ClientStarter struct {
}

func (c *ClientStarter) run(command string, commandArgs []string, gameDir string) *exec.Cmd {
	cmd := exec.Command(command, commandArgs...)

	setProcessFlags(cmd)

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	cmd.Dir = gameDir
	cmd.Env = os.Environ()

	err := cmd.Start()

	if err != nil {
		fmt.Println(fmt.Errorf("ошибка выполнения: %w", err).Error())
	}

	go func() {
		scanner := bufio.NewScanner(stdoutPipe)
		for scanner.Scan() {
			text := scanner.Text()
			fmt.Println("[STDOUT] " + text)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(fmt.Sprintf("stdout scan error: %v", err))
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderrPipe)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	return cmd
}
