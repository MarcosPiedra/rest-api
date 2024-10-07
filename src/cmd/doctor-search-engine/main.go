package main

import (
	"fmt"
	"net/http"
	"os"

	"doctor-search-engine/doctors"
	"doctor-search-engine/internal/config"
	"doctor-search-engine/internal/system"
	"doctor-search-engine/internal/web/static"
	"doctor-search-engine/migrations"
)

type monolith struct {
	*system.System
	doctorModule system.Module
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("backend exitted abnormally: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() (err error) {

	var cfg config.AppConfig
	cfg, err = config.Setup()
	if err != nil {
		return err
	}

	s, err := system.NewSystem(cfg)
	if err != nil {
		return err
	}

	defer func(s *system.System) {
		s.Shutdown()
	}(s)

	m := monolith{
		System:       s,
		doctorModule: doctors.Build(s),
	}

	err = m.MigrateDb(migrations.FS)
	if err != nil {
		return err
	}

	m.doctorModule.Start()

	const swaggerPath = "/swagger/"
	m.Mux().Mount(swaggerPath, http.StripPrefix(swaggerPath, http.FileServer(http.FS(static.SwaggerIndex))))
	const swaggerUiPath = "/swagger-ui/"
	m.Mux().Mount(swaggerUiPath, http.FileServer(http.FS(static.SwaggerUi)))

	fmt.Println("starting doctor-search-engine application")
	defer fmt.Println("stopped doctor-search-engine application")

	m.StartWebServer()

	return nil
}
