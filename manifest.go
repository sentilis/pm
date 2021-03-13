package pm

import (
	"fmt"
	"path"

	"github.com/spf13/viper"
)

// Manifest type

type Manifest struct {
	FileName string
	FileType string
	File     *viper.Viper
}

func (m *Manifest) Load() error {
	m.File = viper.New()
	m.File.SetConfigName(m.FileName)
	m.File.SetConfigType(m.FileType)
	m.File.AddConfigPath(PMDir)

	if err := m.File.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return m.File.SafeWriteConfig()
		}
	}
	return nil
}

func (m Manifest) GetPath() string {
	return path.Join(PMDir, fmt.Sprintf("%s.%s", m.FileName, m.FileType))
}
