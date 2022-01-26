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

	EventMediaCreate = "media.create"
	EventMediaUpdate = "media.update"
	EventMediaDelete = "media.delete"
)

var (
	// Strapi server
	strapiAddr string

	// Hugo site dir
	siteDir string

	// Default locale
	localeDefault string

	// git commit message, leave blank to ignore
	gitCommitMsg string
)

// Sets commands on message
func SetStrapiAddr(value string) {
	fmt.Println("Strapi:", value)
	strapiAddr = value
}

// Sets commands on message
func SetSiteDir(value string) {
	fmt.Println("Site dir:", value)
	siteDir = value
}

// Sets default locale
func SetDefaultLocale(value string) {
	localeDefault = value
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
func gitSync(message string) error {
	if message == "" {
		return errors.New("git commit message required")
	}

	if err := runCommand("git", "pull"); err != nil {
		return err
	}

	if err := runCommand("git", "add", "."); err != nil {
		return err
	}

	if err := runCommand("git", "commit", "-m", message); err != nil {
		return err
	}

	return runCommand("git", "push")

}

// Calls `hugoBuild` and `gitSync`
func buildAndSync(message string) error {
	if err := hugoBuild(); err != nil {
		return err
	}

	if message != "" {
		return gitSync(message)
	}

	return nil
}
