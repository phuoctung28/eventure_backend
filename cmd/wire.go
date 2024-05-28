//go:build wireinject
// +build wireinject

package main

import (
	"eventure_backend/app"
	"eventure_backend/internal"
	"github.com/google/wire"
)

func InitializeApplication() *app.Server {
	wire.Build(internal.ContainerSet,
		app.NewServer,
		app.NewGinEngine)
	return &app.Server{}
}
