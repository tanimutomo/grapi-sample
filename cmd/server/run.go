package main

import (
	"github.com/srvc/appctx"

	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/tanimutomo/grapi-sample/app/server"
)

func run() error {
	// Application context
	ctx := appctx.Global()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewBookServiceServer(),
		),
	)
	return s.Serve(ctx)
}
