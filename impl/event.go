package impl

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

const (
	EventEntryCreate = "entry.create"
	EventEntryUpdate = "entry.update"
	EventEntryDelete = "entry.delete"
)

var (
	// Hugo site dir
	siteDir string

	// git commit message, leave blank to ignore
	gitCommitMsg string
)

// Sets commands on message
func SetSiteDir(value string) {
	fmt.Println("Site dir:", value)
	siteDir = value
}

// Sets git commit message, leave blank to ignore `gitCommit & gitPush`
func SetGitCommitMsg(value string) {
	gitCommitMsg = value
}

// Runs a command
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = siteDir
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Builds the site
func hugoBuild() error {
	if err := runCommand("rm", "-rf", "public"); err != nil {
		return err
	}

	return runCommand("hugo", "--gc", "--minify")
}

// Commits the changes
func gitCommit(message string) error {
	if message == "" {
		return errors.New("git commit message required")
	}

	if err := runCommand("git", "add", "."); err != nil {
		return err
	}

	return runCommand("git", "commit", "-m", message)
}

// Pushs the changes
func gitPush() error {
	return runCommand("git", "push")
}
