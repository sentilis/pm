package main

import (
	"fmt"
	"os"
	"path"

	"github.com/josehbez/pm"
	"github.com/josehbez/pm/changelog"
	"github.com/josehbez/pm/version"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	"github.com/spf13/viper"
)

func main() {
	doc.GenMarkdownTree(run(), "./docs")
}

type command interface {
	Run(*pm.Ctx) *cobra.Command
}

func run() *cobra.Command {

	var rootCmd = &cobra.Command{Use: "pm"}

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get working directory")
		os.Exit(1)
	}

	var ctx = pm.NewCtx()

	ctx.WorkingDir = workingDir
	ctx.Manifest.AddConfigPath(path.Join(ctx.WorkingDir, ctx.PMDir))
	if err := ctx.Manifest.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//ctx.Err.Fatalln("pm is not initialized, first run: pm init")
		}
	}
	for _, cmd := range commandList() {
		rootCmd.AddCommand(cmd.Run(ctx))
	}
	return rootCmd
}

func commandList() []command {
	return []command{
		pm.InitCommand{},
		version.Command{},
		changelog.Command{},
	}
}
