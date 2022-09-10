package learn_bbgo

import (
	"io/ioutil"
	"os"

	"github.com/zixas/learn_bbgo/pkg/datatype"
)

type BuildTargetConfig struct {
	Name    string               `json:"name" yaml:"name"`
	Arch    string               `json:"arch" yaml:"arch"`
	OS      string               `json:"os" yaml:"os"`
	LDFlags datatype.StringSlice `json:"ldflags,omitempty" yaml:"ldflags,omitempty"`
	GCFlags datatype.StringSlice `json:"gcflags,omitempty" yaml:"gcflags,omitempty"`
	Imports []string             `json:"imports,omitempty" yaml:"imports,omitempty"`
}

type BuildConfig struct {
	BuildDir string              `json:"buildDir,omitempty" yaml:"buildDir,omitempty"`
	Imports  []string            `json:"imports,omitempty" yaml:"imports,omitempty"`
	Targets  []BuildTargetConfig `json:"targets,omitempty" yaml:"targets,omitempty"`
}

type Config struct {
	Build *BuildConfig `json:"build,omitempty yaml:"builid,omitempty"`
}

func Load(configFile string, loadStrategies bool) (*Config, error) {
	var config Config

	content, err := ioutil.ReadFile(configFile)

	if err != nil {
		return nil, err
	}
	os.Exit(1)
	return , nil
}
