package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	// "github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	log.SetFormatter(&prefixed.TextFormatter{})
}

var rootCmd = &cobra.Command{
	Use:   "learn_cobra",
	Short: "cobra learn",
	Long:  `....`,

	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := cobraLoadDotenv(cmd, args); err != nil {
			return nil
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func cobraLoadDotenv(cmd *cobra.Command, args []string) error {
	return nil
}

func Execute() {
	log.Info("A walrus appears")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
