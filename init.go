package pm

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//InitCommand ...
type InitCommand struct {
}

// Run ...
func (command InitCommand) Run(ctx *Ctx) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create an empty pm or reinitialize an existing one",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			if err := command.Initialized(ctx); err != nil {
				ctx.Err.Fatal(err)
			}
			ctx.Out.Println(fmt.Sprintf("Initialized in %s", ctx.WorkingDir))
		},
	}
	return cmd
}

// Initialized ..
func (command InitCommand) Initialized(ctx *Ctx) error {
	if _, err := os.Stat(ctx.PMDir); os.IsNotExist(err) {
		if err := os.MkdirAll(ctx.PMDir, 0755); err != nil {
			return err
		}
	}
	if err := ctx.PreLoad(); err != nil {
		return err
	}
	return nil
}

//Exceuted ...
func (command InitCommand) Exceuted(ctx *Ctx) bool {
	if _, err := os.Stat(ctx.PMDir); os.IsNotExist(err) {
		return false
	}
	return true
}
