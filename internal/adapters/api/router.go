package api

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var prefix = "/v0"

type HttpServer struct {
	router *echo.Echo
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		router: echo.New(),
	}
}

func (hs *HttpServer) BuildRoutes() *echo.Group {
	routes := hs.router.Group(prefix)
	routes.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n"}))
	routes.Use(
		middleware.RequestID(),
		middleware.Recover(),
	)
	routes.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 60 * time.Second}))
	routes.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		AllowMethods: []string{
			echo.OPTIONS,
			echo.GET,
			echo.HEAD,
			echo.PUT,
			echo.PATCH,
			echo.POST,
			echo.DELETE,
		},
	}))
	return routes
}

func (r *HttpServer) Start(port string) {
	listeningAddr := ":" + port
	r.router.Logger.Fatal(r.router.Start(listeningAddr))
}
