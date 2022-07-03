package impl

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	EventEntryCreate  = "entry.create"
	EventEntryUpdate  = "entry.update"
	EventEntryDelete  = "entry.delete"
	EventEntryPublish = "entry.publish"

	EventMediaCreate = "media.create"
	EventMediaUpdate = "media.update"
	EventMediaDelete = "media.delete"
)

var debounced func(f func())

// Sets git commit message, leave blank to ignore `gitCommit & gitPush`
func SetGit(msg string, timeout int64) {
	gitCommitMsg = msg
	gitTimeout = timeout

	debounced = Debouncer(time.Duration(timeout * int64(time.Second)))
}

// Runs a command
func runCommand(name string, args ...string) error {
	logger.Infow("run command", "command", name, "args", strings.Join(args, " "))
	cmd := exec.Command(name, args...)
	cmd.Dir = siteDir
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Builds the site
func hugoBuild() error {
	// if err := runCommand("rm", "-rf", "public"); err != nil {
	// 	return err
	// }

	return runCommand("hugo", "--gc", "--minify")
}

// Commits the changes
func gitSync() {
	if gitCommitMsg == "" {
		logger.Warnw("git commit message required")
	}

	if err := runCommand("git", "pull"); err != nil {
		logger.Errorw("git pull", "err", err)
	}

	if err := runCommand("git", "add", "."); err != nil {
		logger.Errorw("git add", "err", err)
	}

	if err := runCommand("git", "commit", "-m", gitCommitMsg); err != nil {
		logger.Errorw("git commit", "err", err)
	}

	if err := runCommand("git", "push"); err != nil {
		logger.Errorw("git push", "err", err)

	}
}

// Calls `hugoBuild` and `gitSync`
func buildAndSync() error {
	if err := hugoBuild(); err != nil {
		return err
	}

	if gitCommitMsg != "" {
		debounced(func() { gitSync() })
	}

	return nil
}
