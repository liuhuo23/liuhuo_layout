//go:build wireinject
// +build wireinject

package main

import (
	"liuhuo23/{{cookiecutter.project_name}}/internal/app"
	"liuhuo23/{{cookiecutter.project_name}}/internal/data"
	"liuhuo23/{{cookiecutter.project_name}}/internal/server"

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
