package apiserver

import (
	"context"
	"fmt"
	"github.com/emicklei/go-restful"
	"go.uber.org/zap"
	apiServerConfig "io.github/devopssphere/pkg/apiserver/config"
	"net/http"
)

type APIServer struct {
	Server *http.Server

	Config *apiServerConfig.Config

	// webservice container, where all webservice defines
	container *restful.Container

	Log *zap.Logger
}

func (s *APIServer) PrepareRun(stopCh <-chan struct{}) error {
	s.container = restful.NewContainer()
	s.container.Router(restful.CurlyRouter{})

	for _, ws := range s.container.RegisteredWebServices() {
		fmt.Println(ws)
	}

	s.Server.Handler = s.container
	s.buildHandlerChain(stopCh)

	return nil
}

func (s *APIServer) Run(ctx context.Context) error {
	err := s.waitForResourceSync(ctx)
	if err != nil {
		return err
	}

	shutdownCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-ctx.Done()
		_ = s.Server.Shutdown(shutdownCtx)
	}()

	s.Log.Sugar().Infof("Start listening on %s", s.Server.Addr)

	err = s.Server.ListenAndServe()

	return err
}

func (s *APIServer) waitForResourceSync(ctx context.Context) error {

	return nil
}

func (s *APIServer) buildHandlerChain(stopCh <-chan struct{}) {

}