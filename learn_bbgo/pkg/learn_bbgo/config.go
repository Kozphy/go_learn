package learn_bbgo

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"

	"github.com/k0kubun/pp/v3"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/zixsa/learn_bbgo/pkg/datatype"
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

type ExchangeStrategyMount struct {
	Mounts   []string               `json:"mounts"`
	Strategy SingleExchangeStrategy `json:"strategy"`
}

type SyncConfig struct {
	// Sessions to sync, if ignored, all defined sessions will sync
	Sessions []string `json:"sessions,omitempty" yaml:"sessions,omitempty"`
}

type Config struct {
	*BuildConfig `json:"buildConfig,omitempty" yaml:"buildConfig,omitempty"`

	// Persistence *PersistenceConfig          `json:"persistence,omitempty" yaml:"persistence,omitempty"`
	// Sessions    map[string]*ExchangeSession `json:"sessions,omitempty" yaml:"sessions,omitempty"`
	ExchangeStrategies      []ExchangeStrategyMount `json:"-" yaml:"-"`
	CrossExchangeStrategies []CrossExchangeStrategy `json:"-" yaml:"-"`
}

func reUnmarshal(conf interface{}, tpe interface{}) (interface{}, error) {
	rt := reflect.TypeOf(tpe)
	val := reflect.New(rt)
	valRef := val.Interface()
	plain, err := json.Marshal(conf)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(plain, valRef); err != nil {
		return nil, errors.Wrapf(err, "json parsing error, given payload: %s", plain)
	}
	return val.Elem().Interface(), nil
}

func NewStrategyFromMap(id string, conf interface{}) (SingleExchangeStrategy, error) {
	if st, ok := LoadedExchangeStrategies[id]; ok {
		val, err := reUnmarshal(conf, st)
		if err != nil {
			return nil, err
		}
		return val.(SingleExchangeStrategy), nil
	}
	return nil, fmt.Errorf("strategy %s not found", id)
}

func loadExchangesStrategies(config *Config, stash Stash) (err error) {
	log.Debug("stash: ")
	pp.Print(stash)
	exchangeStrategiesConfig, ok := stash["exchangeStrategies"]
	if !ok {
		exchangeStrategiesConfig, ok = stash["strategies"]
		if !ok {
			return nil
		}
	}
	log.Debug("LoadedExchangeStrategies: ", LoadedExchangeStrategies)
	// The LoadedExchangeStrategies value setting process is:
	// 1. create init() with learn_bbgo.RegisterStrategy() func in learn_bbgo/pkg/strategy
	// 2. create builtin.go in learn_bbgo/pkg/cmd/strategy and use import with _ to run init() funciton
	// 3. create import.go in learn_bbgo/pkg/cmd and use import with _ to import package pkg/cmd/strategy
	// 4. conclude: when execute cmd package, go will run import.go -> builtin.go -> RegisterStrategy(ID, &Strategy{}) of init() func.
	if len(LoadedExchangeStrategies) == 0 {
		return errors.New("no exchange strategy is registered")
	}
	// filter out exchangeStrategies content
	configList, ok := exchangeStrategiesConfig.([]interface{})
	if !ok {
		return errors.New("expecting list([]interface{}) in exchangeStrategies")
	}

	log.Debug("configList: ")
	pp.Print(configList)
	// TODO: fix strategy not found in config
	for _, entry := range configList {
		configStash, ok := entry.(Stash)
		if !ok {
			return fmt.Errorf("strategy config should be a map[string]interface{}, given: %T %+v", entry, entry)
		}
		log.Debug("configStash")
		pp.Print(configStash)

		var mounts []string
		if val, ok := configStash["on"]; ok {
			switch tv := val.(type) {
			case []string:
				mounts = append(mounts, tv...)
			case string:
				mounts = append(mounts, tv)
			case []interface{}:
				for _, f := range tv {
					s, ok := f.(string)
					if !ok {
						return fmt.Errorf("%v (%T) is not a string", f, f)
					}
					mounts = append(mounts, s)
				}
			default:
				return fmt.Errorf("unexpected mount type: %T value %+v", val, val)
			}
		}

		for id, conf := range configStash {
			if _, ok := LoadedExchangeStrategies[id]; ok {
				st, err := NewStrategyFromMap(id, conf)
				if err != nil {
					return err
				}

				config.ExchangeStrategies = append(config.ExchangeStrategies, ExchangeStrategyMount{
					Mounts:   mounts,
					Strategy: st,
				})
			} else if id != "on" && id != "off" {
				return fmt.Errorf("strategy %s in config not found", id)
			}
		}
	}

	os.Exit(1)
	return nil
}

func loadCrossExchangeStrategies(config *Config, stash Stash) (err error) {
	// exchangeStrategiesConf, ok := stash["crossExchangeStrategies"]
	return nil
}

type Stash map[string]interface{}

func loadStash(config []byte) (Stash, error) {
	stash := make(Stash)
	if err := yaml.Unmarshal(config, stash); err != nil {
		return nil, err
	}
	return stash, nil
}

// loadConfig file with strategies
func Load(configFile string, loadStrategies bool) (*Config, error) {
	log.Infof("start Load config file with load Strategies %v", loadStrategies)
	var config Config

	// origin ioutil
	content, err := os.ReadFile(configFile)

	if err != nil {
		return nil, err
	}
	// decode content to config
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

	// decode content to stash
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

func GetNativeBuildTargetConfig() BuildTargetConfig {
	return BuildTargetConfig{
		Name: "bbgow",
		Arch: runtime.GOARCH,
		OS:   runtime.GOOS,
	}
}
