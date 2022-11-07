package app

import (
	"context"
	"github.com/spf13/cobra"
	"io.github/devopssphere/cmd/ds-apiserver/app/options"
	signals "io.github/devopssphere/pkg/utils"
	"net/http"
)

func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()

	cmd := &cobra.Command{
		Use: "ds-apiserver",
		RunE: func(cmd *cobra.Command, args []string) error {

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
	apiserver, err := s.NewAPIServer(ctx.Done())
	if err != nil {
		return err
	}

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
