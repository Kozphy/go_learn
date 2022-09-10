package cmd

import (
	"fmt"
	"os"

	// "runtime"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	"github.com/zixas/learn_bbgo/pkg/learn_bbgo"
)

var userConfig *learn_bbgo.Config

func init() {
	log.SetFormatter(&prefixed.TextFormatter{})

	RootCmd.PersistentFlags().Bool("debug", false, "debug mode")

	RootCmd.PersistentFlags().Bool("no-dotenv", false, "disable built-in dotenv")
	RootCmd.PersistentFlags().String("dotenv", ".env.local", "the dotenv file you want to load")

	RootCmd.PersistentFlags().String("config", "config.yaml", "config file")

}

var RootCmd = &cobra.Command{
	Use:   "learn_bbgo",
	Short: "bbgo learn",

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := cobraLoadDotenv(cmd, args); err != nil {
			return err
		}

		if viper.GetBool("debug") {
			log.Infof("debug mode is enabled")
			log.SetLevel(log.DebugLevel)
		}

		return cobraLoadConfig(cmd, args)
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func cobraLoadDotenv(cmd *cobra.Command, args []string) error {
	disableDotEnv, err := cmd.Flags().GetBool("no-dotenv")
	if err != nil {
		return err
	}
	if !disableDotEnv {
		dotenvFile, err := cmd.Flags().GetString("dotenv")
		if err != nil {
			return err
		}

		if _, err := os.Stat(dotenvFile); err == nil {
			if err := godotenv.Load(dotenvFile); err != nil {
				return errors.Wrap(err, "error loading dotenv file")
			}
		}
	}
	return nil
}

func cobraLoadConfig(cmd *cobra.Command, args []string) error {
	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		return errors.Wrapf(err, "failed to get the config flag")
	}

	if len(configFile) > 0 {
		if _, err := os.Stat(configFile); err == nil {
			userConfig, err = learn_bbgo.Load(configFile, false)
			if err != nil {
				// if err is nil, return nil, otherwise return an error with stack trace.
				return errors.Wrapf(err, "can not load config file: %s", configFile)
			}
			// IsNotExist returns a boolean indicating whether the error is known to report that a file or directory does not exist.
		} else if os.IsNotExist(err) {
			userConfig = &learn_bbgo.Config{}
		} else {
			return errors.Wrapf(err, "config file load error: %s", configFile)
		}
	}

	return nil
}

func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
