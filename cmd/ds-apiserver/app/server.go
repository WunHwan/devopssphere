package app

import (
	"context"
	"github.com/spf13/cobra"
	"io.github/devopssphere/cmd/ds-apiserver/app/options"
	apiServerConfig "io.github/devopssphere/pkg/apiserver/config"
	"io.github/devopssphere/pkg/simple/client/db"
	signals "io.github/devopssphere/pkg/utils"
	"log"
	"net/http"
)

func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()

	// Load configuration from file
	if conf, err := apiServerConfig.TryLoadFromDisk(); err != nil {
		log.Fatal("Failed to load configuration from disk", err)
	} else {
		s = &options.ServerRunOptions{
			GenericServerRunOptions: s.GenericServerRunOptions,
			Config:                  conf,
		}
	}

	cmd := &cobra.Command{
		Use:   "ds-apiserver",
		Short: "",
		Long:  "",
		Args: func(cmd *cobra.Command, args []string) error {
			cmd.PersistentFlags().Bool("debug-mode", false, "open debug mode")

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if debugMode, _ := cmd.Flags().GetBool("debug-mode"); debugMode {
				s.DebugMode = true
			}

			return Run(s, signals.SetupSignalHandler())
		},
		SilenceUsage: true,
	}

	return cmd
}

func Run(s *options.ServerRunOptions, ctx context.Context) error {
	sctx, cancelFunc := context.WithCancel(context.TODO())
	errCh := make(chan error)
	defer close(errCh)
	go func() {
		if err := run(s, sctx); err != nil {
			errCh <- err
		}
	}()

	for {
		select {
		case <-ctx.Done():
			cancelFunc()
			return nil
		case err := <-errCh:
			cancelFunc()
			return err
		}
	}
}

func run(s *options.ServerRunOptions, ctx context.Context) error {
	apiserver, err := s.NewAPIServer()
	if err != nil {
		return err
	}

	database, err := db.NewDatabase(s.Config.DatabaseOptions)
	if err != nil {
		return err
	}
	apiserver.Database = database

	err = apiserver.PrepareRun(ctx.Done())
	if err != nil {
		return err
	}

	err = apiserver.Run(ctx)
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}
