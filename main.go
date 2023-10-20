package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/muhwyndhamhp/spill/config"
	"github.com/muhwyndhamhp/spill/public"
	"github.com/muhwyndhamhp/spill/template"
	"github.com/muhwyndhamhp/spill/utils/routing"
)

func main() {
	e := echo.New()
	routing.SetupRouter(e)

	e.Static("/dist", "dist")

	template.NewTemplateRenderer(e)

	e.GET("/", func(c echo.Context) error {
		component := public.Index("Wyndham")
		return template.AssertRender(c, http.StatusOK, component)
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Get(config.APP_PORT))))
}
