package semver

import (
	"flag"
)

// MinorCommand ..
type MinorCommand struct {
}

const minorShortHelp = ``
const minorLongHelp = ``

// Name ...
func (cmd *MinorCommand) Name() string { return "minor" }

// Args ...
func (cmd *MinorCommand) Args() string { return "..." }

// ShortHelp ...
func (cmd *MinorCommand) ShortHelp() string { return minorShortHelp }

// LongHelp ...
func (cmd *MinorCommand) LongHelp() string { return minorLongHelp }

// Hidden ...
func (cmd *MinorCommand) Hidden() bool { return false }

// Register ...
func (cmd *MinorCommand) Register(fs *flag.FlagSet) {

}

// Run ...
func (cmd *MinorCommand) Run(ctx *Ctx, args []string) error {
	if _, err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	minor := ctx.Manifest.Viper.GetInt("version.minor") + 1
	ctx.Manifest.Viper.Set("version.minor", minor)
	ctx.Manifest.Viper.Set("version.patch", 0)

	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil

}
