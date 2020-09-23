package semver

import (
	"flag"

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
	//root := ctx.WorkingDir
	if len(args) > 1 {
		return errors.Errorf("too many args (%d)", len(args))
	}
	return nil
}
