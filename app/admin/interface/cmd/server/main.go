package main

import (
	"os"

	"github.com/beiduoke/go-scaffold/pkg/bootstrap"
	"github.com/beiduoke/go-scaffold/pkg/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
)

var (
	id, _   = os.Hostname()
	Service = bootstrap.NewServiceInfo(
		service.AdminService,
		"1.0.0",
		id,
	)
)

func newApp(logger log.Logger, rr registry.Registrar, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(Service.GetInstanceId()),
		kratos.Name(Service.Name),
		kratos.Version(Service.Version),
		kratos.Metadata(Service.Metadata),
		kratos.Logger(logger),
		kratos.Registrar(rr),
		kratos.Server(
			hs,
		),
	)
}

func main() {
	cfg, ll, reg := bootstrap.Bootstrap(Service)

	app, cleanup, err := wireApp(ll, reg, cfg)
	if err != nil {
		panic(err)
	}

	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
