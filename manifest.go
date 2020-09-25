package semver

import "github.com/spf13/viper"

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

// DefaultVersion ...
func (m Manifest) DefaultVersion() error {

	m.Viper.Set("version.major", 0)
	m.Viper.Set("version.minor", 1)
	m.Viper.Set("version.patch", 0)
	m.Viper.Set("version.label", "")
	return m.Viper.SafeWriteConfig()
}

type rawManifest struct {
	Version rawVersion `mapstructure:"version"`
}
type rawVersion struct {
	major int    `mapstructure:"major"`
	minor int    `mapstructure:"minor"`
	patch int    `mapstructure:"patch"`
	label string `mapstructure:"label"`
}

type rawReleaseVersion struct {
}
type rawBuildVersion struct {
}
