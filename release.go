package semver

import (
	"flag"
	"fmt"
	"strings"
)

// ReleaseCommand ..
type ReleaseCommand struct {
	label  bool
	patch  bool
	remove bool
}

const releaseShortHelp = `Software release cycle life [alfa < beta < release-candidate < release | release-tls < discontinued]`
const releaseLongHelp = ``

// Name ...
func (cmd *ReleaseCommand) Name() string { return "release" }

// Args ...
func (cmd *ReleaseCommand) Args() string { return "" }

// ShortHelp ...
func (cmd *ReleaseCommand) ShortHelp() string { return releaseShortHelp }

// LongHelp ...
func (cmd *ReleaseCommand) LongHelp() string { return releaseLongHelp }

// Hidden ...
func (cmd *ReleaseCommand) Hidden() bool { return false }

// Register ...
func (cmd *ReleaseCommand) Register(fs *flag.FlagSet) {
	fs.BoolVar(&cmd.label, "label", false, "Set label software release cycle life [alfa < beta < release-candidate < release | release-tls < discontinued]")
	fs.BoolVar(&cmd.patch, "patch", false, "Increase the patch")
	fs.BoolVar(&cmd.remove, "rm", false, "Remove release")
}

// Run ...
func (cmd *ReleaseCommand) Run(ctx *Ctx, args []string) error {
	if cmd.label {
		return cmd.runLabel(ctx, args)
	} else if cmd.patch {
		return cmd.runPatch(ctx, args)
	} else if cmd.remove {
		return cmd.runRemove(ctx, args)
	}
	return nil

}
func (cmd *ReleaseCommand) runRemove(ctx *Ctx, args []string) error {
	ctx.Manifest.Viper.Set("release.label", "")
	ctx.Manifest.Viper.Set("release.patch", 0)
	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func (cmd *ReleaseCommand) runPatch(ctx *Ctx, args []string) error {
	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	label := ctx.Manifest.Viper.GetString("release.label")
	if len(label) == 0 {
		return fmt.Errorf("warng: first run semver pre-release -label")
	}
	patch := ctx.Manifest.Viper.GetInt("release.patch") + 1
	ctx.Manifest.Viper.Set("release.patch", patch)
	if err := ctx.Manifest.Viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}
func (cmd *ReleaseCommand) runLabel(ctx *Ctx, args []string) error {

	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	cycleLife := SoftwareReleaseCycleLife("simplified")

	label := ctx.Manifest.Viper.GetString("release.label")
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
		ctx.Manifest.Viper.Set("release.label", labelNew)
		ctx.Manifest.Viper.Set("release.patch", 0)
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
func WorkingOnPreRelease(ctx *Ctx) error {
	releaseLabel := ctx.Manifest.Viper.GetString("release.label")
	if len(releaseLabel) > 0 {
		for _, l := range SoftwareReleaseCycleLife("simplified-pre-release") {
			if l == releaseLabel {
				return fmt.Errorf("Warning: You are working on a pre-release.\n\nUsege: \n\n semver pre-release -patch")
			}
		}
	}
	return nil
}
