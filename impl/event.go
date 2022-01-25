package impl

import (
	"os"
	"os/exec"
)

const (
	EventEntryCreate = "entry.create"
	EventEntryUpdate = "entry.update"
	EventEntryDelete = "entry.delete"

	hugoBuildCmd = "rm -rf public && hugo --gc --minify"
	gitAddCmd    = "git status" //"git add ."
)

var siteDir string

// Sets commands on message
func SetSiteDir(value string) {
	siteDir = value
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
	return runCommand(hugoBuildCmd)
}

// Commits the changes
func GitCommit() error {
	if err := runCommand(gitAddCmd); err != nil {
		return err
	}

	return nil
}

// Pushs the changes
func GitPush() error {
	return nil
}
