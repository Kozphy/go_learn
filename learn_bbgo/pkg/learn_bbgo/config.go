package learn_bbgo

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/zixas/learn_bbgo/pkg/datatype"
	"gopkg.in/yaml.v3"
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
	*BuildConfig `json:"buildConfig,omitempty" yaml:"buildConfig,omitempty"`

	// Persistence *PersistenceConfig          `json:"persistence,omitempty" yaml:"persistence,omitempty"`
	// Sessions    map[string]*ExchangeSession `json:"sessions,omitempty" yaml:"sessions,omitempty"`
	// ExchangeStrategies      []ExchangeStrategyMount `json:"-" yaml:"-"`
	// CrossExchangeStrategies []CrossExchangeStrategy `json:"-" yaml:"-"`
}

type Stash map[string]interface{}

func loadStash(config []byte) (Stash, error) {
	stash := make(Stash)
	if err := yaml.Unmarshal(config, stash); err != nil {
		return nil, err
	}
	return stash, nil
}

func loadExchangesStrategies(config *Config, stash Stash) (err error) {
	// exchangeStrategiesConfig, ok := stash["exchangeStrategies"]
	return nil
}

func loadCrossExchangeStrategies(config *Config, stash Stash) (err error) {
	// exchangeStrategiesConf, ok := stash["crossExchangeStrategies"]
	return nil
}

func Load(configFile string, loadStrategies bool) (*Config, error) {
	log.Info("start Load config file")
	var config Config

	// origin ioutil
	content, err := os.ReadFile(configFile)

	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}

	// for backward compatible
	if config.BuildConfig == nil {
		config.BuildConfig = &BuildConfig{
			BuildDir: "build",
			Targets: []BuildTargetConfig{
				{Name: "bbgow-amd64-darwin", Arch: "amd64", OS: "darwin"},
				{Name: "bbgow-amd64-linux", Arch: "amd64", OS: "linux"},
			},
		}
	}

	log.Debug("config:", config)

	stash, err := loadStash(content)

	if loadStrategies {
		if err := loadExchangesStrategies(&config, stash); err != nil {
			return nil, err
		}

		if err := loadCrossExchangeStrategies(&config, stash); err != nil {
			return nil, err
		}
	}

	return &config, nil
}
