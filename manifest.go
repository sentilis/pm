package semver

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

// ManifestName is the manifest file name used by semver.
const ManifestName = "semver"

// ManifestType is the manisfest file name format
const ManifestType = "toml"

// Manifest holds manifest file data
type Manifest struct {
	Name  string
	Viper *viper.Viper
}

// NewManifest ..
func NewManifest() *Manifest {
	v := viper.New()
	v.SetConfigName(ManifestName)
	v.SetConfigType(ManifestType)

	return &Manifest{
		Name:  ManifestName,
		Viper: v,
	}
}

// ValidateManifest ...
func (m *Manifest) ValidateManifest() error {

	if err := m.Viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("fatal: has not initialized semver")
		}
		return err
	}
	rManifest := rawManifest{}
	m.Viper.UnmarshalExact(&rManifest)
	rVersion := rManifest.Version
	if rVersion.Major < 0 || rVersion.Minor < 0 || rVersion.Patch < 0 {
		return fmt.Errorf("fatal: Version number MUST are non-negative integers %d.%d.%d", rVersion.Major, rVersion.Minor, rVersion.Patch)
	}

	return nil

}

// DefaultVersion ...
func (m Manifest) DefaultVersion() error {

	m.Viper.Set("version.major", 0)
	m.Viper.Set("version.minor", 1)
	m.Viper.Set("version.patch", 0)
	m.Viper.Set("version.label", "")
	return m.Viper.SafeWriteConfig()
}
func (m Manifest) version() {
	//	m.Viper.Un
}

type rawManifest struct {
	Version rawVersion `mapstructure:"version"`
}
type rawVersion struct {
	Major int    `mapstructure:"major"`
	Minor int    `mapstructure:"minor"`
	Patch int    `mapstructure:"patch"`
	Label string `mapstructure:"label"`
}

type rawReleaseVersion struct {
}
type rawBuildVersion struct {
}
