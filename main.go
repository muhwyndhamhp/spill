package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/labstack/echo/v4"
	"github.com/muhwyndhamhp/spill/auth"
	"github.com/muhwyndhamhp/spill/config"
	"github.com/muhwyndhamhp/spill/db"
	"github.com/muhwyndhamhp/spill/modules/company"
	"github.com/muhwyndhamhp/spill/modules/spilluser"
	v1 "github.com/muhwyndhamhp/spill/modules/spilluser/v1"
	"github.com/muhwyndhamhp/spill/public"
	"github.com/muhwyndhamhp/spill/template"
	"github.com/muhwyndhamhp/spill/utils/routing"
)

func main() {
	e := echo.New()
	routing.SetupRouter(e)

	e.Static("/dist", "dist")

	template.NewTemplateRenderer(e)

	client, err := clerk.NewClient(config.Get(config.CLERK_SK_KEY))
	if err != nil {
		panic(err)
	}
	sessionMid := echo.WrapMiddleware(clerk.WithSessionV2(client, clerk.WithLeeway(10*time.Second)))
	e.Use(sessionMid)

	g := e.Group("/v1")

	ctx := context.Background()
	spillUserRepo := spilluser.NewSpillUserRepository(db.GetDB(ctx))
	companyRepo := company.NewCompanyRepository(db.GetDB(ctx))
	v1.NewSpillUserFrontend(g, &client, spillUserRepo, companyRepo)

	e.GET("/", func(c echo.Context) error {
		name := ""
		session, ok := clerk.SessionFromContext(c.Request().Context())
		if !ok {
			return c.Redirect(http.StatusFound, "/v1/login")
		}
		ss := &auth.Session{Claim: session}
		u, err := ss.GetUser(client)
		if err != nil {
			return c.Redirect(http.StatusFound, "/v1/login")
		}

		cExist, _ := companyRepo.CompanyWithServiceIDExists(c.Request().Context(), u.ID)

		if !cExist {
			return c.Redirect(http.StatusFound, "/v1/users/companies/register")
		}

		name = fmt.Sprintf("%s %s", *u.FirstName, *u.LastName)
		component := public.Index(name)
		return template.AssertRender(c, http.StatusOK, component)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Get(config.APP_PORT))))
}
