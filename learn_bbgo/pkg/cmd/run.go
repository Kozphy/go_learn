package cmd

import (
	"context"
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "run strategies from config file",

	RunE: run,
}

func init() {
	RunCmd.Flags().Bool("setup", false, "use setup mode")
	RunCmd.Flags().Bool("no-compile", false, "do not compile wrapper binary")
	RootCmd.AddCommand(RunCmd)
}

func run(cmd *cobra.Command, args []string) error {
	setup, err := cmd.Flags().GetBool("setup")
	if err != nil {
		return err
	}
	noCompile, err := cmd.Flags().GetBool("no-compile")
	if err != nil {
		return err
	}

	// in pkg/cmd/root.go
	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		return err
	}

	if !setup {
		if len(configFile) == 0 {
			return errors.New("--config option is required")
		}

		if _, err := os.Stat(configFile); err != nil {
			return err
		}

	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
}
