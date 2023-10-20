package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sergiovirahonda/echo-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag"
)

const goTemplateAPI = "go-template"

// Swagger Documentation
func ServeDocs(routeGroups *echo.Group) {
	docs.SwaggerInfo.InfoInstanceName = goTemplateAPI
	swag.Register(docs.SwaggerInfo.InstanceName(), docs.SwaggerInfo)
	routeGroups.GET("/docs/*", echoSwagger.EchoWrapHandler(
		func(c *echoSwagger.Config) { c.InstanceName = goTemplateAPI }),
	)
}
