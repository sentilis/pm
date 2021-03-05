package pm

import (
	"log"
	"os"
	"path"

	"github.com/spf13/viper"
)

const (
	//ManifestName ..
	ManifestName = "pm"
	//ManifestType ..
	ManifestType = "yml"
)

const (
	// ChangelogName ..
	ChangelogName = "changelog"
	//ChangelogType ..
	ChangelogType = "yml"
)

// PMDir ... workdir
const PMDir = ".pm"

//Ctx ...
type Ctx struct {
	WorkingDir, PMDir   string
	Out, Err            *log.Logger
	Manifest, Changelog *viper.Viper
}

//NewCtx ..
func NewCtx() *Ctx {
	v := viper.New()
	v.SetConfigName(ManifestName)
	v.SetConfigType(ManifestType)

	return &Ctx{
		Manifest: v,
		Out:      log.New(os.Stdout, "", 0),
		Err:      log.New(os.Stderr, "", 0),
		PMDir:    PMDir,
	}

}

// InitManifest ...
func (ctx Ctx) InitManifest() error {
	ctx.Manifest.Set("version.major", 0)
	ctx.Manifest.Set("version.minor", 1)
	ctx.Manifest.Set("version.patch", 0)
	return ctx.Manifest.SafeWriteConfig()
}

//LoadChangelog ..
func (ctx *Ctx) LoadChangelog() error {

	v := viper.NewWithOptions(viper.KeyDelimiter("::"))
	v.SetConfigName(ChangelogName)
	v.SetConfigType(ChangelogType)
	v.AddConfigPath(path.Join(ctx.WorkingDir, ctx.PMDir))
	ctx.Changelog = v
	if err := ctx.Changelog.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return ctx.Changelog.SafeWriteConfig()
		}
	}
	return nil
}
