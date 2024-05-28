// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"eventure_backend/app"
	"eventure_backend/internal/api"
)

// Injectors from wire.go:

func InitializeApplication() *app.Server {
	engine := app.NewGinEngine()
	router := api.NewRouter()
	server := app.NewServer(engine, router)
	return server
}
