package semv

import (
	"flag"
)

// MajorCommand ..
type MajorCommand struct {
}

const majorShortHelp = `Version when you make incompatible API changes`
const majorLongHelp = ``

// Name ...
func (cmd *MajorCommand) Name() string { return "major" }

// Args ...
func (cmd *MajorCommand) Args() string { return "" }

// ShortHelp ...
func (cmd *MajorCommand) ShortHelp() string { return majorShortHelp }

// LongHelp ...
func (cmd *MajorCommand) LongHelp() string { return majorLongHelp }

// Hidden ...
func (cmd *MajorCommand) Hidden() bool { return false }

// Register ...
func (cmd *MajorCommand) Register(fs *flag.FlagSet) {

}

// Run ...
func (cmd *MajorCommand) Run(ctx *Ctx, args []string) error {
	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	if err := WorkingOnPreRelease(ctx); err != nil {
		return err
	}
	major := ctx.Manifest.Viper.GetInt("version.major") + 1
	ctx.Manifest.Viper.Set("version.major", major)
	ctx.Manifest.Viper.Set("version.minor", 0)
	ctx.Manifest.Viper.Set("version.patch", 0)

	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil

}
