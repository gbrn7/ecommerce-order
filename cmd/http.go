package cmd

import (
	"ecommerce-order/external"
	"ecommerce-order/helpers"
	"ecommerce-order/internal/api"
	"ecommerce-order/internal/interfaces"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	d := dependencyInject()

	e := echo.New()
	e.GET("/healthcheck", d.HealthcheckAPI.Healthcheck)

	e.Start(":" + helpers.GetEnv("PORT", "9000"))
}

type Dependency struct {
	External       interfaces.IExternal
	HealthcheckAPI *api.HealthcheckAPI
}

func dependencyInject() Dependency {

	external := &external.External{}

	return Dependency{
		External:       external,
		HealthcheckAPI: &api.HealthcheckAPI{},
	}
}
