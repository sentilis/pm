package semver

import (
	"flag"
)

// StatusCommand ..
type StatusCommand struct {
}

const statusShortHelp = ``
const statusLongHelp = ``

// Name ...
func (cmd *StatusCommand) Name() string { return "status" }

// Args ...
func (cmd *StatusCommand) Args() string { return "" }

// ShortHelp ...
func (cmd *StatusCommand) ShortHelp() string { return statusShortHelp }

// LongHelp ...
func (cmd *StatusCommand) LongHelp() string { return statusLongHelp }

// Hidden ...
func (cmd *StatusCommand) Hidden() bool { return false }

// Register ...
func (cmd *StatusCommand) Register(fs *flag.FlagSet) {

}

// Run ...
func (cmd *StatusCommand) Run(ctx *Ctx, args []string) error {
	if _, err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	ctx.Out.Printf("%d.%d.%d",
		ctx.Manifest.Viper.Get("version.major"),
		ctx.Manifest.Viper.Get("version.minor"),
		ctx.Manifest.Viper.Get("version.patch"),
	)
	return nil

}
