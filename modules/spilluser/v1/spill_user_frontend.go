package v1

import (
	"encoding/json"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/labstack/echo/v4"
	"github.com/muhwyndhamhp/spill/auth"
	"github.com/muhwyndhamhp/spill/modules/spilluser"
	"github.com/muhwyndhamhp/spill/modules/spilluser/dto"
	"github.com/muhwyndhamhp/spill/public"
	"github.com/muhwyndhamhp/spill/template"
	"github.com/muhwyndhamhp/spill/utils/errs"
	"github.com/muhwyndhamhp/spill/utils/resp"
)

type SpillUserFrontend struct {
	clerk *clerk.Client
	g     *echo.Group
	ctrl  *spilluser.SpillUserController
}

func NewSpillUserFrontend(g *echo.Group, clerk *clerk.Client, ctrl *spilluser.SpillUserController) {
	handler := &SpillUserFrontend{
		g:     g,
		clerk: clerk,
		ctrl:  ctrl,
	}

	g.GET("/login", handler.Login)

	g.GET("/redirect/login", handler.RedirectLogin)

	g.GET("/redirect/register", handler.RedirectRegister)

	g.POST("/register", handler.Register)
}

func (h SpillUserFrontend) Login(c echo.Context) error {
	component := public.Login()
	return template.AssertRender(c, http.StatusOK, component)
}

func (h SpillUserFrontend) RedirectLogin(c echo.Context) error {
	session := "none"
	sessionClaims, ok := clerk.SessionFromContext(c.Request().Context())
	if ok {
		ss := &auth.Session{Claim: sessionClaims}
		u, _ := ss.GetUser(*h.clerk)

		str, _ := json.Marshal(u)
		session = string(str)
	}
	component := public.RedirectLogin(session)
	return template.AssertRender(c, http.StatusOK, component)
}

func (h SpillUserFrontend) RedirectRegister(c echo.Context) error {
	sessionClaims, ok := clerk.SessionFromContext(c.Request().Context())
	if !ok {
		c.Response().Header().Add("Hx-Redirect", "/login")
		return c.JSON(http.StatusUnauthorized, nil)
	}

	ss := &auth.Session{Claim: sessionClaims}
	u, _ := ss.GetUser(*h.clerk)
	if u == nil {
		c.Response().Header().Add("Hx-Redirect", "/login")
		return c.JSON(http.StatusUnauthorized, nil)
	}

	spillUser, _ := h.ctrl.GetUserByServiceID(c.Request().Context(), u.ID)
	if spillUser != nil {
		c.Response().Header().Add("Hx-Redirect", "/")
		return c.JSON(http.StatusOK, nil)
	}

	component := public.RedirectRegister()
	return template.AssertRender(c, http.StatusOK, component)
}

func (h SpillUserFrontend) Register(c echo.Context) error {
	sessionClaims, ok := clerk.SessionFromContext(c.Request().Context())
	if !ok {
		c.Response().Header().Add("Hx-Redirect", "/login")
		return c.JSON(http.StatusUnauthorized, nil)
	}

	req := dto.PostRegisterReq{}

	if err := c.Bind(&req); err != nil {
		return resp.HTTPBadRequest(c, "", err.Error())
	}

	ss := &auth.Session{Claim: sessionClaims}
	u, _ := ss.GetUser(*h.clerk)
	if u == nil {
		c.Response().Header().Add("Hx-Redirect", "/login")
		return c.JSON(http.StatusUnauthorized, nil)
	}

	if _, err := h.ctrl.CreateSpillUser(c.Request().Context(), req.Alias, req.Bio, u.ID); err != nil {
		return errs.Wrap(err)
	}

	c.Response().Header().Add("Hx-Redirect", "/")
	return c.JSON(http.StatusOK, nil)
}
