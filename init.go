package pm

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"

	"github.com/pkg/errors"
)

// InitCommand ..
type InitCommand struct {
}

const initShortHelp = `Set up a new Project Metadata Management`
const initLongHelp = ``

// Name ...
func (cmd *InitCommand) Name() string { return "init" }

// Args ...
func (cmd *InitCommand) Args() string { return "[root]" }

// ShortHelp ...
func (cmd *InitCommand) ShortHelp() string { return initShortHelp }

// LongHelp ...
func (cmd *InitCommand) LongHelp() string { return initLongHelp }

// Hidden ...
func (cmd *InitCommand) Hidden() bool { return false }

// Register ...
func (cmd *InitCommand) Register(fs *flag.FlagSet) {}

// Run ...
func (cmd *InitCommand) Run(ctx *Ctx, args []string) error {
	if err := ctx.Manifest.Viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if errDV := ctx.Manifest.DefaultVersion(); errDV != nil {
				return errDV
			}
			ctx.Out.Printf("Initialized in %s", ctx.WorkingDir)
			return nil
		}
		return err
	}
	return errors.New(fmt.Sprintf("Reinitialized in %s", ctx.WorkingDir))

}
