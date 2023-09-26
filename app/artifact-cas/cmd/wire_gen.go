// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chainloop-dev/chainloop/app/artifact-cas/internal/conf"
	"github.com/chainloop-dev/chainloop/app/artifact-cas/internal/server"
	"github.com/chainloop-dev/chainloop/app/artifact-cas/internal/service"
	"github.com/chainloop-dev/chainloop/internal/credentials"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, auth *conf.Auth, reader credentials.Reader, logger log.Logger) (*app, func(), error) {
	providers := loadCASBackendProviders(reader)
	v := serviceOpts(logger)
	byteStreamService := service.NewByteStreamService(providers, v...)
	resourceService := service.NewResourceService(providers, v...)
	grpcServer, err := server.NewGRPCServer(confServer, auth, byteStreamService, resourceService, providers, logger)
	if err != nil {
		return nil, nil, err
	}
	downloadService := service.NewDownloadService(providers, v...)
	httpServer, err := server.NewHTTPServer(confServer, auth, downloadService, providers, logger)
	if err != nil {
		return nil, nil, err
	}
	httpMetricsServer, err := server.NewHTTPMetricsServer(confServer)
	if err != nil {
		return nil, nil, err
	}
	mainApp := newApp(logger, grpcServer, httpServer, httpMetricsServer, providers)
	return mainApp, func() {
	}, nil
}

// wire.go:

func serviceOpts(l log.Logger) []service.NewOpt {
	return []service.NewOpt{service.WithLogger(l)}
}
