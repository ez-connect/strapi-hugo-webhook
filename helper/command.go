package helper

import (
	"os"
	"os/exec"
	"time"

	"strapiwebhook/helper/zlog"
)

var (
	debounced func(f func())
)

func InitCommand(workingDir string, timeout int64) {
	debounced = Debouncer(time.Duration(timeout * int64(time.Second)))
}

// Runs multiple commands in a single shell instance
func RunCommand(dir, commands string) {
	zlog.Infow("shell", "commands", commands)
	cmd := exec.Command("/bin/sh", "-c", commands)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		zlog.Errorw("run command", "error", err)
	}
}

// Commits the changes
func RunDebouncedCommand(dir, commands string) {
	zlog.Infow("shell (debounced)", "commands", commands)
	debounced(func() {
		RunCommand(dir, commands)
	})
}
