package impl

import (
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	EventEntryCreate = "entry.create"
	EventEntryUpdate = "entry.update"
	EventEntryDelete = "entry.delete"

	EventMediaCreate = "media.create"
	EventMediaUpdate = "media.update"
	EventMediaDelete = "media.delete"
)

var debounced func(f func())

// Sets commands on message
func SetStrapiAddr(value string) {
	strapiAddr = value
}

// Sets commands on message
func SetSiteDir(value string) {
	siteDir = value
}

// Sets default locale
func SetDefaultLocale(value string) {
	localeDefault = value
}

// Sets git commit message, leave blank to ignore `gitCommit & gitPush`
func SetGit(msg string, timeout int64) {
	gitCommitMsg = msg
	gitTimeout = timeout

	debounced = Debouncer(time.Duration(timeout * int64(time.Second)))
}

// Runs a command
func runCommand(name string, args ...string) error {
	GetLogger().Infow("run command", "command", name, "args", strings.Join(args, " "))
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
		GetLogger().Warnw("git commit message required")
	}

	if err := runCommand("git", "pull"); err != nil {
		GetLogger().Errorw("git pull", "err", err)
	}

	if err := runCommand("git", "add", "."); err != nil {
		GetLogger().Errorw("git add", "err", err)
	}

	if err := runCommand("git", "commit", "-m", gitCommitMsg); err != nil {
		GetLogger().Errorw("git commit", "err", err)
	}

	if err := runCommand("git", "push"); err != nil {
		GetLogger().Errorw("git push", "err", err)

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
