package semver

// ManifestName is the manifest file name used by semver.
const ManifestName = "semver.toml"

// Manifest holds manifest file data
type Manifest struct {
	Name string
}

// NewManifest ..
func NewManifest() Manifest {
	return Manifest{Name: ManifestName}
}

type rawVersion struct {
	major int `toml:"major"`
	minor int `toml:"minor"`
	patch int `toml:"patch"`
	label int `toml:"label,omitempty"`
}

type rawReleaseVersion struct {
}
type rawBuildVersion struct {
}
