package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckFlag ensures that flag is not empty
func CheckFlag(cmd *cobra.Command, flagVar string) {
	if flagVar == "" {
		cmd.Help()
		os.Exit(0)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m[Info]: %s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[33;1m[Warning]: %s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Error should be used to display an error and exit
func Error(format string, args ...interface{}) {
	fmt.Printf("\x1b[31;1m[Error]: %s\x1b[0m\n", fmt.Sprintf(format, args...))
	os.Exit(1)
}
