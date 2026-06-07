//go:build !windows

package main

import (
	"os/exec"
)

func setProcessFlags(cmd *exec.Cmd) {
}
