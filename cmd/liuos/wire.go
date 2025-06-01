//go:build wireinject
// +build wireinject

package main

import (
	"liuhuo23/liuos/internal/app"
	"liuhuo23/liuos/internal/data"
	"liuhuo23/liuos/internal/server"

	"github.com/google/wire"
)

func InitializeApp() *app.App {
	panic(wire.Build(
		app.NewAppConfig,
		app.NewDatabaseConfig,
		server.NewEcho,
		data.NewDb,
		app.NewApp,
	))
}
