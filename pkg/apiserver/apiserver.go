package apiserver

import (
	"context"
	"github.com/emicklei/go-restful"
	"go.uber.org/zap"
	"gorm.io/gorm"
	apiTenant "io.github/devopssphere/pkg/apis/tenant"
	apiServerConfig "io.github/devopssphere/pkg/apiserver/config"
	"io.github/devopssphere/pkg/models/tenant"
	"net/http"
)

type APIServer struct {
	Server *http.Server

	Config *apiServerConfig.Config

	// webservice container, where all webservice defines
	container *restful.Container

	// database container
	Database *gorm.DB

	Log *zap.Logger
}

func (s *APIServer) PrepareRun(stopCh <-chan struct{}) error {
	s.container = restful.NewContainer()
	s.container.Router(restful.CurlyRouter{})

	s.installDevopsAPIs()

	for _, ws := range s.container.RegisteredWebServices() {
		s.Log.Sugar().Infof("%s", ws.RootPath())
	}

	s.Server.Handler = s.container
	s.buildHandlerChain(stopCh)

	return nil
}

func (s *APIServer) installDevopsAPIs() {
	log := s.Log.Sugar()
	_ = apiTenant.AddToContainer(s.container, log, tenant.NewOperator(s.Database))

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
