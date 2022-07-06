package main

import (
	"strapiwebhook/base/cmd"
	"strapiwebhook/impl"
)

func main() {
	// You can add new command here
	// Call `cmd.GetRootCmd()` to get the root command
	impl.AddCmd(cmd.GetServeCmd())

	cmd.Execute()
}
