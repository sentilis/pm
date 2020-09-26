package semver

import (
	"flag"
)

// PatchCommand ..
type PatchCommand struct {
}

const patchShortHelp = `Version when you make backwards compatible bug fixes`
const patchLongHelp = ``

// Name ...
func (cmd *PatchCommand) Name() string { return "patch" }

// Args ...
func (cmd *PatchCommand) Args() string { return "" }

// ShortHelp ...
func (cmd *PatchCommand) ShortHelp() string { return patchShortHelp }

// LongHelp ...
func (cmd *PatchCommand) LongHelp() string { return patchLongHelp }

// Hidden ...
func (cmd *PatchCommand) Hidden() bool { return false }

// Register ...
func (cmd *PatchCommand) Register(fs *flag.FlagSet) {

}

// Run ...
func (cmd *PatchCommand) Run(ctx *Ctx, args []string) error {
	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	patch := ctx.Manifest.Viper.GetInt("version.patch") + 1
	ctx.Manifest.Viper.Set("version.patch", patch)

	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil

}
