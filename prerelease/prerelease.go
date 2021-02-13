package pm

import (
	"flag"
	"fmt"
	"strings"

	"github.com/josehbez/pm"
)

// Command ..
type Command struct {
	label  bool
	major  bool
	remove bool
}

const releaseShortHelp = `Software release cycle life [alfa < beta < release-candidate < release | release-tls < discontinued]`
const releaseLongHelp = ``

// Name ...
func (cmd *Command) Name() string { return "pre-release" }

// Args ...
func (cmd *Command) Args() string { return "" }

// ShortHelp ...
func (cmd *Command) ShortHelp() string { return releaseShortHelp }

// LongHelp ...
func (cmd *Command) LongHelp() string { return releaseLongHelp }

// Hidden ...
func (cmd *Command) Hidden() bool { return false }

// Register ...
func (cmd *Command) Register(fs *flag.FlagSet) {
	fs.BoolVar(&cmd.label, "label", false, "Set label software release cycle life [alfa < beta < release-candidate < release | release-tls < discontinued]")
	fs.BoolVar(&cmd.major, "major", false, "Increase the major")
	fs.BoolVar(&cmd.remove, "remove", false, "Remove release")
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

	ctx.Manifest.Viper.Set("prerelease.label", "")
	ctx.Manifest.Viper.Set("prerelease.major", 0)
	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func (cmd *Command) runMajor(ctx *pm.Ctx, args []string) error {
	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	label := ctx.Manifest.Viper.GetString("prerelease.label")
	if len(label) == 0 {
		return fmt.Errorf("warng: first run semver pre-release -label")
	}
	major := ctx.Manifest.Viper.GetInt("prerelease.major") + 1
	ctx.Manifest.Viper.Set("prerelease.major", major)
	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}
func (cmd *Command) runLabel(ctx *pm.Ctx, args []string) error {

	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	cycleLife := SoftwareReleaseCycleLife("simplified")

	label := ctx.Manifest.Viper.GetString("prerelease.label")
	labelID := -1

	labelNew := ""
	labelNewID := -1
	if len(label) > 0 {
		if len(args) == 0 {
			return fmt.Errorf("fatal: semver pre-release -label \t Select one of the tags %s", cycleLife)
		}
		labelNew = args[0]
		for i, l := range cycleLife {
			if label == l {
				labelID = i
			}
			if labelNew == l {
				labelNewID = i
			}
		}
		if labelNewID < labelID {
			return fmt.Errorf("fatal: Use this Software Release Cycle Life %s", cycleLife)
		} else if labelNewID == labelID {
			return fmt.Errorf("warng: Select one tag after %s - %s", label, cycleLife)
		}
	} else {
		labelNew = "alfa"
		if len(args) > 0 {
			srcl := false
			for _, l := range cycleLife {
				if l == args[0] {
					srcl = true
					labelNew = l
				}
			}
			if !srcl {
				return fmt.Errorf("fatal: only lebel allowed - %s ", cycleLife)
			}
		}
	}

	if len(labelNew) > 0 {
		ctx.Manifest.Viper.Set("prerelease.label", labelNew)
		ctx.Manifest.Viper.Set("prerelease.mejor", 0)
		if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
			return err
		}
	}

	return nil
}

// SoftwareReleaseCycleLife ...
func SoftwareReleaseCycleLife(kind string) []string {
	var cycle, preRelease, release []string

	if strings.Contains(kind, "simplified") {
		preRelease = []string{
			"alfa",
			"beta",
			"release-candidate",
		}
		release = []string{
			"release",
			"release-tls",
			"discontinued",
		}
		if kind == "simplified-pre-release" {
			cycle = preRelease
		} else if kind == "simplified-release" {
			cycle = release
		} else {
			cycle = append(preRelease, release...)
		}
	}
	return cycle
}

// WorkingOnPreRelease ...
func WorkingOnPreRelease(ctx *pm.Ctx) error {
	releaseLabel := ctx.Manifest.Viper.GetString("prerelease.label")
	if len(releaseLabel) > 0 {
		for _, l := range SoftwareReleaseCycleLife("simplified-pre-release") {
			if l == releaseLabel {
				return fmt.Errorf("Warning: You are working on a pre-release.\n\nUsege: \n\n semver pre-release -major")
			}
		}
	}
	return nil
}
