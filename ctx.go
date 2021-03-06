package pm

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

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

const (
	//VersionName ..
	VersionName = "version"
	//VersionType ..
	VersionType = "yml"
)

// PMDir ... workdir
const PMDir = ".pm"

//Ctx ...
type Ctx struct {
	WorkingDir, PMDir            string
	Out, Err                     *log.Logger
	Manifest, Changelog, Version *viper.Viper
}

//NewCtx ..
func NewCtx() *Ctx {
	return &Ctx{
		Out:   log.New(os.Stdout, "", 0),
		Err:   log.New(os.Stderr, "", 0),
		PMDir: PMDir,
	}

}

//PreLoad ...
func (ctx *Ctx) PreLoad() error {
	if err := ctx.LoadManifest(); err != nil {
		return err
	}
	if err := ctx.LoadVersion(); err != nil {
		return err
	}
	if err := ctx.LoadChangelog(); err != nil {
		return err
	}
	return nil
}

/// MANIFEST

// LoadManifest ...
func (ctx *Ctx) LoadManifest() error {
	ctx.Manifest = viper.New()
	ctx.Manifest.SetConfigName(ManifestName)
	ctx.Manifest.SetConfigType(ManifestType)
	ctx.Manifest.AddConfigPath(path.Join(ctx.WorkingDir, ctx.PMDir))

	if err := ctx.Manifest.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return ctx.Manifest.SafeWriteConfig()
		}
	}
	return nil
}

// GetManifestPath ..
func (ctx Ctx) GetManifestPath() string {
	return path.Join(ctx.WorkingDir, ctx.PMDir, fmt.Sprintf("%s.%s", ManifestName, ManifestType))
}

// VERSION

//LoadVersion ..
func (ctx *Ctx) LoadVersion() error {

	v := viper.New()
	v.SetConfigName(VersionName)
	v.SetConfigType(VersionType)
	v.AddConfigPath(path.Join(ctx.WorkingDir, ctx.PMDir))
	ctx.Version = v
	if err := ctx.Version.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			ctx.Version.Set("version.major", 0)
			ctx.Version.Set("version.minor", 1)
			ctx.Version.Set("version.patch", 0)
			return ctx.Version.SafeWriteConfig()
		}
	}
	return nil
}

// GetVersionPath ..
func (ctx Ctx) GetVersionPath() string {
	return path.Join(ctx.WorkingDir, ctx.PMDir, fmt.Sprintf("%s.%s", VersionName, VersionType))
}

// Changelog

//LoadChangelog ..
func (ctx *Ctx) LoadChangelog() error {

	ctx.Changelog = viper.NewWithOptions(viper.KeyDelimiter("::"))
	ctx.Changelog.SetConfigName(ChangelogName)
	ctx.Changelog.SetConfigType(ChangelogType)
	ctx.Changelog.AddConfigPath(path.Join(ctx.WorkingDir, ctx.PMDir))

	if err := ctx.Changelog.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			v := [1]string{"Initialized pm (Project Metadata)"}
			ctx.Changelog.Set(time.Now().Format("2006-01-02"), v)
			return ctx.Changelog.SafeWriteConfig()
		}
	}
	return nil
}

// GetChangelogPath ..
func (ctx Ctx) GetChangelogPath() string {
	return path.Join(ctx.WorkingDir, ctx.PMDir, fmt.Sprintf("%s.%s", ChangelogName, ChangelogType))
}
