package build

import (
	"flag"
	"fmt"

	"github.com/josehbez/pm"
)

// Command ..
type Command struct {
	label  bool
	major  bool
	remove bool
}

const buildShortHelp = ``
const buildLongHelp = ``

// Name ...
func (cmd *Command) Name() string { return "build" }

// Args ...
func (cmd *Command) Args() string { return "" }

// ShortHelp ...
func (cmd *Command) ShortHelp() string { return buildShortHelp }

// LongHelp ...
func (cmd *Command) LongHelp() string { return buildLongHelp }

// Hidden ...
func (cmd *Command) Hidden() bool { return false }

// Register ...
func (cmd *Command) Register(fs *flag.FlagSet) {
	fs.BoolVar(&cmd.label, "label", false, "Set label build metadata")
	fs.BoolVar(&cmd.major, "major", false, "Increase the major")
	fs.BoolVar(&cmd.remove, "remove", false, "Remove build metadata")
}

// Run ...
func (cmd *Command) Run(ctx *pm.Ctx, args []string) error {
	if cmd.label {
		return cmd.runLabel(ctx, args)
	} else if cmd.major {
		return cmd.runMajor(ctx, args)
	} else if cmd.remove {
		return cmd.runRemove(ctx, args)
	}
	return nil

}
func (cmd *Command) runRemove(ctx *pm.Ctx, args []string) error {
	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	ctx.Manifest.Viper.Set("build.label", "")
	ctx.Manifest.Viper.Set("build.major", 0)
	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func (cmd *Command) runMajor(ctx *pm.Ctx, args []string) error {
	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	label := ctx.Manifest.Viper.GetString("build.label")
	if len(label) == 0 {
		return fmt.Errorf("warng: first run semver build -label")
	}
	major := ctx.Manifest.Viper.GetInt("build.major") + 1
	ctx.Manifest.Viper.Set("build.major", major)
	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}
func (cmd *Command) runLabel(ctx *pm.Ctx, args []string) error {

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
		ctx.Manifest.Viper.Set("build.major", 0)
		if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
			return err
		}
	}

	return nil
}
