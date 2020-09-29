package semver

import (
	"flag"
	"fmt"
)

// StatusCommand ..
type StatusCommand struct {
}

const statusShortHelp = `Show semantic-versioning version-release+build`
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
	if err := ctx.Manifest.ValidateManifest(); err != nil {
		return err
	}
	version := fmt.Sprintf(
		"%d.%d.%d",
		ctx.Manifest.Viper.GetInt("version.major"),
		ctx.Manifest.Viper.GetInt("version.minor"),
		ctx.Manifest.Viper.GetInt("version.patch"),
	)
	release := ""
	releaseLabel := ctx.Manifest.Viper.GetString("release.label")
	if len(releaseLabel) > 0 {
		releasePatch := ctx.Manifest.Viper.GetString("release.patch")
		release = fmt.Sprintf("-%s.%s", releaseLabel, releasePatch)
	}

	build := ""
	buildLabel := ctx.Manifest.Viper.GetString("build.label")
	if len(buildLabel) > 0 {
		buildPatch := ctx.Manifest.Viper.GetString("build.patch")
		build = fmt.Sprintf("+%s.%s", buildLabel, buildPatch)
	}

	ctx.Out.Printf("%s%s%s", version, release, build)
	return nil

}
