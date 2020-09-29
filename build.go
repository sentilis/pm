package semver

import (
	"flag"
	"fmt"
)

// BuildCommand ..
type BuildCommand struct {
	label  bool
	patch  bool
	remove bool
}

const BuildShortHelp = ``
const BuildLongHelp = ``

// Name ...
func (cmd *BuildCommand) Name() string { return "build" }

// Args ...
func (cmd *BuildCommand) Args() string { return "" }

// ShortHelp ...
func (cmd *BuildCommand) ShortHelp() string { return BuildShortHelp }

// LongHelp ...
func (cmd *BuildCommand) LongHelp() string { return BuildLongHelp }

// Hidden ...
func (cmd *BuildCommand) Hidden() bool { return false }

// Register ...
func (cmd *BuildCommand) Register(fs *flag.FlagSet) {
	fs.BoolVar(&cmd.label, "label", false, "Set label build metadata")
	fs.BoolVar(&cmd.patch, "patch", false, "Increase the patch")
	fs.BoolVar(&cmd.remove, "rm", false, "Remove build metadata")
}

// Run ...
func (cmd *BuildCommand) Run(ctx *Ctx, args []string) error {
	if cmd.label {
		return cmd.runLabel(ctx, args)
	} else if cmd.patch {
		return cmd.runPatch(ctx, args)
	} else if cmd.remove {
		return cmd.runRemove(ctx, args)
	}
	return nil

}
func (cmd *BuildCommand) runRemove(ctx *Ctx, args []string) error {
	ctx.Manifest.Viper.Set("build.label", nil)
	ctx.Manifest.Viper.Set("build.patch", 0)
	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func (cmd *BuildCommand) runPatch(ctx *Ctx, args []string) error {
	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	label := ctx.Manifest.Viper.GetString("build.label")
	if len(label) == 0 {
		return fmt.Errorf("warng: first run semver build -label")
	}
	patch := ctx.Manifest.Viper.GetInt("build.patch") + 1
	ctx.Manifest.Viper.Set("build.patch", patch)
	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}
func (cmd *BuildCommand) runLabel(ctx *Ctx, args []string) error {

	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}

	label := ctx.Manifest.Viper.GetString("build.label")

	labelNew := ""

	if len(label) > 0 {
		if len(args) == 0 {
			return fmt.Errorf("fatal: semver build -label \t  set build label")
		}
		labelNew = args[0]
	} else {
		labelNew = "build"
	}

	if len(labelNew) > 0 {
		ctx.Manifest.Viper.Set("build.label", labelNew)
		ctx.Manifest.Viper.Set("build.patch", 0)
		if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
			return err
		}
	}

	return nil
}
