package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/josehbez/semver"
)

func main() {
	os.Exit(Run())
}

var (
	successExitCode = 0
	errorExitCode   = 1
)

type command interface {
	Name() string           // "foobar"
	Args() string           // "<baz> [quux...]"
	ShortHelp() string      // "Foo the first bar"
	LongHelp() string       // "Foo the first bar meeting the following conditions..."
	Register(*flag.FlagSet) // command-specific flags
	Hidden() bool           // indicates whether the command should be hidden from help output
	Run(*semver.Ctx, []string) error
}

// Run executes a configuration and returns an exit code.
func Run() int {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to get working directory", err)
		os.Exit(errorExitCode)
	}

	commands := commandList()
	cmdName, printCommandHelp, exit := parseArgs(os.Args)

	if exit {
		fprintUsage(os.Stderr)
		return errorExitCode
	}
	outLogger := log.New(os.Stdout, "", 0)
	errLogger := log.New(os.Stderr, "", 0)

	for _, cmd := range commands {
		if cmd.Name() == cmdName {
			flags := flag.NewFlagSet(cmdName, flag.ContinueOnError)
			cmd.Register(flags)
			if printCommandHelp {
				flags.Usage()
				return errorExitCode
			}
			if err := flags.Parse(os.Args[2:]); err != nil {
				return errorExitCode
			}
			ctx := &semver.Ctx{
				Out:        outLogger,
				Err:        errLogger,
				WorkingDir: workingDir,
				Manifest:   semver.NewManifest(),
			}
			if err := cmd.Run(ctx, flags.Args()); err != nil {
				errLogger.Printf("%v\n", err)
				return errorExitCode
			}
			return successExitCode
		}
	}
	errLogger.Printf("sev: %s: no such command\n", cmdName)
	return errorExitCode
}

func commandList() []command {
	return []command{
		&semver.InitCommand{},
	}

}

func parseArgs(args []string) (cmdName string, printCmdUsage bool, exit bool) {
	isHelpArg := func() bool {
		return strings.Contains(strings.ToLower(args[1]), "help") || strings.ToLower(args[1]) == "-h"
	}

	switch len(args) {
	case 0, 1:
		exit = true
	case 2:
		if isHelpArg() {
			exit = true
		} else {
			cmdName = args[1]
		}
	default:
		if isHelpArg() {
			cmdName = args[2]
			printCmdUsage = true
		} else {
			cmdName = args[1]
		}
	}
	return cmdName, printCmdUsage, exit
}

func fprintUsage(w io.Writer) {
	fmt.Fprintln(w, "semver is a tool for managing semantic-versioning")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Usage: \"semver [command]\"")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Commands:")
	fmt.Fprintln(w)
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)

	commands := commandList()
	for _, cmd := range commands {
		if !cmd.Hidden() {
			fmt.Fprintf(tw, "\t%s\t%s\n", cmd.Name(), cmd.ShortHelp())
		}
	}
	tw.Flush()
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Examples:")
	for _, example := range examples {
		fmt.Fprintf(tw, "\t%s\t%s\n", example[0], example[1])
	}
	tw.Flush()
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Use \"semver help [command]\" for more information about a command.")
}

var examples = [...][2]string{
	{
		"semver init",
		"set up a new sematic-versioning version-release+build",
	},
	{
		"semver release",
		"Set up label release -beta.0.0.0",
	},
	{
		"semver build",
		"set up build metadata +build.0.0.0",
	},
	{
		"semver changelog",
		"Create file CHANGELOG.rst and append lastet version header",
	},
}
