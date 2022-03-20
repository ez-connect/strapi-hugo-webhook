package impl

import (
	"fmt"
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

var (
	// Strapi server
	strapiAddr string

	// Hugo site dir
	siteDir string

	// Default locale
	localeDefault string

	// Git commit message, leave blank to ignore
	gitCommitMsg string

	// Git timeout
	gitTimeout int64
	debounced  func(f func())
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
func SetGit(msg string, timeout int64) {
	gitCommitMsg = msg
	gitTimeout = timeout

	debounced = Debouncer(time.Duration(timeout * int64(time.Second)))
}

// Runs a command
func runCommand(name string, args ...string) error {
	fmt.Println("command:", name, strings.Join(args, " "))
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
func gitSync() {
	if gitCommitMsg == "" {
		fmt.Println("git commit message required")
	}

	if err := runCommand("git", "pull"); err != nil {
		fmt.Println(err)
	}

	if err := runCommand("git", "add", "."); err != nil {
		fmt.Println(err)
	}

	if err := runCommand("git", "commit", "-m", gitCommitMsg); err != nil {
		fmt.Println(err)
	}

	runCommand("git", "push")
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
