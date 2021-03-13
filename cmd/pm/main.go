package main

import (
	"fmt"
	"os"

	"github.com/josehbez/pm"
	"github.com/josehbez/pm/author"
	"github.com/josehbez/pm/changelog"
	"github.com/josehbez/pm/license"
	"github.com/josehbez/pm/version"
	"github.com/spf13/cobra"
)

func main() {
	run().Execute()
}

type command interface {
	Run(*pm.Ctx) *cobra.Command
}

func run() *cobra.Command {

	var rootCmd = &cobra.Command{
		Use: "pm",
	}

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get working directory")
		os.Exit(1)
	}

	var ctx = pm.NewCtx(workingDir)

	initCommand := pm.InitCommand{}
	if initCommand.Exceuted(ctx) {
		if err := ctx.PreLoad(); err != nil {
			fmt.Println("failed to pre-load files")
			os.Exit(1)
		}
		for _, cmd := range commandList() {
			rootCmd.AddCommand(cmd.Run(ctx))
		}
	} else {
		rootCmd.AddCommand(initCommand.Run(ctx))
	}

	return rootCmd
}

func commandList() []command {
	return []command{
		pm.InitCommand{},
		version.Command{},
		changelog.Command{},
		license.Command{},
		author.Command{},
		author.CommandMaintainer{},
	}
}
