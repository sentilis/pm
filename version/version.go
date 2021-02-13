package version

import (
	"flag"
	"fmt"

	"github.com/josehbez/pm"
)

// Command ..
type Command struct {
	label bool
	major bool
	minor bool
	patch bool
}

const majorShortHelp = `Version when you make incompatible API changes`
const majorLongHelp = ``

// Name ...
func (cmd *Command) Name() string { return "major" }

// Args ...
func (cmd *Command) Args() string { return "" }

// ShortHelp ...
func (cmd *Command) ShortHelp() string { return majorShortHelp }

// LongHelp ...
func (cmd *Command) LongHelp() string { return majorLongHelp }

// Hidden ...
func (cmd *Command) Hidden() bool { return false }

// Register ...
func (cmd *Command) Register(fs *flag.FlagSet) {
	fs.BoolVar(&cmd.label, "label", false, "")
	fs.BoolVar(&cmd.major, "major", false, "")
	fs.BoolVar(&cmd.minor, "minor", false, "")
	fs.BoolVar(&cmd.patch, "patch", false, "")
}

// Run ...
func (cmd *Command) Run(ctx *pm.Ctx, args []string) error {
	if cmd.label {
		//return cmd.runLabel(ctx, args)
	} else if cmd.major {
		return cmd.runMajor(ctx, args)
	} else if cmd.minor {
		return cmd.runMinor(ctx, args)
	} else if cmd.patch {
		return cmd.runPatch(ctx, args)
	}
	return nil
}

func (cmd *Command) runLabel(ctx *pm.Ctx, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("fatal: pm version -label NAME")
	}
	labelNew := args[0]
	if len(labelNew) > 0 {
		ctx.Manifest.Viper.Set("version.label", labelNew)
		ctx.Manifest.Viper.Set("version.mejor", 0)
		if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
			return err
		}
	}
	return nil
}
func (cmd *Command) runMajor(ctx *pm.Ctx, args []string) error {
	major := ctx.Manifest.Viper.GetInt("version.major") + 1
	ctx.Manifest.Viper.Set("version.major", major)
	ctx.Manifest.Viper.Set("version.minor", 0)
	ctx.Manifest.Viper.Set("version.patch", 0)

	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func (cmd *Command) runMinor(ctx *pm.Ctx, args []string) error {

	minor := ctx.Manifest.Viper.GetInt("version.minor") + 1
	ctx.Manifest.Viper.Set("version.minor", minor)
	ctx.Manifest.Viper.Set("version.patch", 0)

	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil

}

func (cmd *Command) runPatch(ctx *pm.Ctx, args []string) error {
	patch := ctx.Manifest.Viper.GetInt("version.patch") + 1
	ctx.Manifest.Viper.Set("version.patch", patch)

	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil

}
