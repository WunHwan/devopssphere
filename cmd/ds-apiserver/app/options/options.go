package options

import (
	"fmt"
	"go.uber.org/zap"
	"io.github/devopssphere/pkg/apiserver"
	apiServerConfig "io.github/devopssphere/pkg/apiserver/config"
	genericoptions "io.github/devopssphere/pkg/server/options"
	"net/http"
)

type ServerRunOptions struct {
	ConfigFile              string
	GenericServerRunOptions *genericoptions.ServerRunOptions
	*apiServerConfig.Config

	//
	DebugMode bool
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		GenericServerRunOptions: genericoptions.NewServerRunOptions(),
		Config:                  apiServerConfig.New(),
	}
}

func (s *ServerRunOptions) NewAPIServer(stopch <-chan struct{}) (*apiserver.APIServer, error) {
	apiServer := &apiserver.APIServer{
		Config: s.Config,
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	apiServer.Log = logger

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", s.GenericServerRunOptions.InsecurePort),
	}

	apiServer.Server = server

	return apiServer, nil
}
