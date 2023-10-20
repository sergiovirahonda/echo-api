package main

import (
	"github.com/sergiovirahonda/echo-api/internal/adapters/api"
	"github.com/sergiovirahonda/echo-api/internal/adapters/db"
	"github.com/sergiovirahonda/echo-api/internal/cfg"
	"github.com/sergiovirahonda/echo-api/internal/repositories/echo"
	echos "github.com/sergiovirahonda/echo-api/internal/services/echo"
)

// @title Echo API
// @version 1.0
// @description This is a sample and very simple Echo API.

// @contact.name Sergio Virahonda
// @contact.url https://www.linkedin.com/in/sergiovirahonda/
// @contact.email svirahonda@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /v0
func main() {
	// App configurations
	conf := cfg.GetConfig()

	// Database connection
	database := db.NewConnection(*conf)
	// Migrations executed automatically bc of SQLite
	db.Migrate(database)

	// HTTP server
	httpServer := api.NewHttpServer()
	routes := httpServer.BuildRoutes()

	// Dependency injection for controllers
	echoRepo := echo.NewDefaultRepository(database)
	echoService := echos.NewDefaultService(echoRepo)
	api.NewController(routes, echoService)

	// Swagger
	api.ServeDocs(routes)

	// HTTP server initialization
	httpServer.Start(conf.Server.Port)
}
