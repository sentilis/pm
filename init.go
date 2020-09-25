package semver

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"

	"github.com/pkg/errors"
)

// InitCommand ..
type InitCommand struct {
}

const initShortHelp = `Set up a new SEV project`
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
func (cmd *InitCommand) Register(fs *flag.FlagSet) {

}

// Run ...
func (cmd *InitCommand) Run(ctx *Ctx, args []string) error {
	if len(args) > 1 {
		return errors.Errorf("too many args (%d)", len(args))
	}
	if err := ctx.Manifest.Viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if errDV := ctx.Manifest.DefaultVersion(); errDV != nil {

				return errDV
			}
			ctx.Out.Printf("Initialized semver in %s", ctx.WorkingDir)
			return nil
		}
		return err
	}
	return errors.New(fmt.Sprintf("Reinitialized semver in %s", ctx.WorkingDir))

}
