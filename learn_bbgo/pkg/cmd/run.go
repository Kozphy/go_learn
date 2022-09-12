package cmd

import (
	"context"
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/zixas/learn_bbgo/pkg/learn_bbgo"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "run strategies from config file",

	RunE: run,
}

func init() {
	RunCmd.Flags().Bool("setup", false, "use setup mode")
	RunCmd.Flags().Bool("no-compile", false, "do not compile wrapper binary")
	RunCmd.Flags().Bool("no-sync", false, "do not sync on startup")

	RunCmd.Flags().Bool("enable-webserver", false, "enable webserver")
	RunCmd.Flags().String("webserver-bind", ":8080", "webserver binding")

	RunCmd.Flags().Bool("enable-grpc", false, "enable grpc server")
	RunCmd.Flags().String("grpc-bind", ":50051", "grpc server binding")

	RunCmd.Flags().Bool("lightweight", false, "lightweight mode")
	RootCmd.AddCommand(RunCmd)
}

// func runSetup(baseCtx context.Context, userConfig *learn_bbgo.Config, enableApiServer bool) error {
// 	ctx, cancelTrading := context.WithCancel(baseCtx)
// 	defer cancelTrading()

// 	environ := learn_bbgo.NewEnvironment()

// 	trader := learn_bbgo.NewTrader(environ)

// 	if enableApiServer {
// 		go func() {
// 			s := &server.Server{
// 				Config:        userConfig,
// 				Environ:       environ,
// 				Trader:        trader,
// 				OpenInBrowser: true,
// 				Setup: &server.Setup{
// 					Context: ctx,
// 					Cancel:  cancelTrading,
// 					Token:   "",
// 				},
// 			}

// 			if err := s.Run(ctx); err != nil {
// 				log.WithError(err).Errorf("server error")
// 			}
// 		}()
// 	}

// }

func runConfig(basectx context.Context, cmd *cobra.Command, userConfig *learn_bbgo.Config) error {
	// noSync, err := cmd.Flags().GetBool("no-sync")
	// if err != nil {
	// 	return err
	// }

	// enableWebServer, err := cmd.Flags().GetBool("enable-webserver")
	// if err != nil {
	// 	return err
	// }

	// webServerBind, err := cmd.Flags().GetString("webserver-bind")
	// if err != nil {
	// 	return err
	// }
	// enableGrpc, err := cmd.Flags().GetBool("enable-grpc")
	// if err != nil {
	// 	return err
	// }

	// grpcBind, err := cmd.Flags().GetString("grpc-bind")
	// if err != nil {
	// 	return err
	// }

	// _ = grpcBind
	// _ = enableGrpc

	// ctx, cancelTrading := context.WithCancel(basectx)
	// defer cancelTrading()

	// environ := bbgo.NewEnvironment()

	// lightweight, err := cmd.Flags().GetBool("lightweight")
	return nil
}

func run(cmd *cobra.Command, args []string) error {
	log.Info("start run")
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
	// detect --config cmd and configfile  whether exist
	if !setup {
		if len(configFile) == 0 {
			return errors.New("--config option is required")
		}

		if _, err := os.Stat(configFile); err != nil {
			return err
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	log.Debug("ctx: ", ctx)
	defer cancel()

	// log.Debug("BuildConfig:", userConfig.BuildConfig)
	// log.Debug("BuildConfig import: ", userConfig.BuildConfig.Imports)
	if learn_bbgo.IsWrapperBinary || (userConfig.BuildConfig != nil && len(userConfig.BuildConfig.Imports) == 0) || noCompile {
		if learn_bbgo.IsWrapperBinary {
			log.Info("running wrappper binary...")
		}

		// if setup {
		// 	return runSetup(ctx, userConfig, true)
		// }

		// LoadStrategies is true
		userConfig, err = learn_bbgo.Load(configFile, true)
		if err != nil {
			return err
		}
		return runConfig(ctx, cmd, userConfig)
	}
	return runWrapperBinary(ctx, cmd, userConfig, args)
}

func runWrapperBinary(ctx context.Context, cmd *cobra.Command, userConfig *learn_bbgo.Config, args []string) error {
	var runArgs = []string{"run"}
	cmd.Flags().Visit(func(flag *flag.Flag) {
		runArgs = append(runArgs, "--"+flag.Name, flag.Value.String())
	})
	runArgs = append(runArgs, args...)
	log.Debug("runArgs: ", runArgs)
	return nil
}
