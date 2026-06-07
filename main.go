package main

import (
	"os"
)

func main() {
	_ = os.Args[0]
	command := os.Args[1]
	var commandArgs []string
	var gameDirIndex int
	var gameDir string

	for i, arg := range os.Args {
		if i < 2 {
			continue
		}

		if arg == "--gameDir" {
			gameDirIndex = i
		}

		if i == gameDirIndex+1 {
			gameDir = arg
		}
		commandArgs = append(commandArgs, arg)
	}

	starter := ClientStarter{}
	cmd := starter.run(command, commandArgs, gameDir)

	watcher := FileWatcher{}
	watcher.run(gameDir, cmd)
	cmd.Wait()
}
