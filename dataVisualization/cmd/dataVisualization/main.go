// Package main is the grpc server of the application.
package main

import (
	"dataVisualization/cmd/dataVisualization/initial"

	"github.com/zhufuyi/sponge/pkg/app"
)

func main() {
	initial.InitApp()
	services := initial.CreateServices()
	closes := initial.Close(services)

	a := app.New(services, closes)
	a.Run()
}
