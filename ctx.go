package pm

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// ManifestName ... file name
const ManifestName = ".pm"

// ManifestType ... file name format
const ManifestType = "yml"

//Ctx ...
type Ctx struct {
	WorkingDir string
	Out, Err   *log.Logger
	Manifest   *viper.Viper
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
	}

}

// InitManifest ...
func (c Ctx) InitManifest() error {
	c.Manifest.Set("version.major", 0)
	c.Manifest.Set("version.minor", 1)
	c.Manifest.Set("version.patch", 0)
	c.Manifest.Set("version.label", "")
	return c.Manifest.SafeWriteConfig()
}
