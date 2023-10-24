package v1

import (
	"encoding/json"
	"net/http"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/muhwyndhamhp/spill/auth"
	"github.com/muhwyndhamhp/spill/db"
	"github.com/muhwyndhamhp/spill/modules/company"
	"github.com/muhwyndhamhp/spill/modules/spilluser"
	"github.com/muhwyndhamhp/spill/modules/spilluser/dto"
	"github.com/muhwyndhamhp/spill/public"
	"github.com/muhwyndhamhp/spill/template"
	"github.com/muhwyndhamhp/spill/utils/errs"
	"github.com/muhwyndhamhp/spill/utils/resp"
)

type SpillUserFrontend struct {
	clerk       *clerk.Client
	g           *echo.Group
	repo        *spilluser.SpillUserRepository
	companyRepo *company.CompanyRepository
}

func NewSpillUserFrontend(
	g *echo.Group,
	clerk *clerk.Client,
	repo *spilluser.SpillUserRepository,
	companyRepo *company.CompanyRepository,
) {
	handler := &SpillUserFrontend{
		g:           g,
		clerk:       clerk,
		repo:        repo,
		companyRepo: companyRepo,
	}

	g.GET("/redirect/register", handler.RedirectRegister)

	g.POST("/register", handler.Register)

	g.GET("/users/companies/register", handler.UsersCompaniesRegister)

	g.POST("/users/companies/upsert", handler.UsersCompaniesUpsert)
}

func (h SpillUserFrontend) UsersCompaniesUpsert(c echo.Context) error {
	sessionClaims, ok := clerk.SessionFromContext(c.Request().Context())
	if !ok {
		return c.Redirect(http.StatusFound, "/v1/login")
	}

	req := dto.PostUsersCompaniesUpsertReq{}

	if err := c.Bind(&req); err != nil {
		return resp.HTTPBadRequest(c, "", err.Error())
	}

	ss := &auth.Session{Claim: sessionClaims}
	u, _ := ss.GetUser(*h.clerk)
	if u == nil {
		return c.Redirect(http.StatusFound, "/v1/login")
	}
	user, _ := h.repo.GetUserByServiceID(c.Request().Context(), u.ID)
	if user == nil {
		return c.Redirect(http.StatusFound, "/v1/login")
	}

	if err := h.companyRepo.UpsertCompany(c.Request().Context(), req.CompanyName, user.ID); err != nil {
		return errs.Wrap(err)
	}

	c.Response().Header().Add("Hx-Redirect", "/")
	return c.JSON(http.StatusOK, nil)
}

func (h SpillUserFrontend) UsersCompaniesRegister(c echo.Context) error {
	_, ok := clerk.SessionFromContext(c.Request().Context())
	if !ok {
		return c.Redirect(http.StatusFound, "/v1/login")
	}

	component := public.UsersCompaniesRegister()
	return template.AssertRender(c, http.StatusOK, component)
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
		return c.Redirect(http.StatusFound, "/v1/login")
	}

	ss := &auth.Session{Claim: sessionClaims}
	u, _ := ss.GetUser(*h.clerk)
	if u == nil {
		return c.Redirect(http.StatusFound, "/v1/login")
	}

	spillUser, _ := h.repo.GetUserByServiceID(c.Request().Context(), u.ID)
	if spillUser != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	component := public.RedirectRegister()
	return template.AssertRender(c, http.StatusOK, component)
}

func (h SpillUserFrontend) Register(c echo.Context) error {
	sessionClaims, ok := clerk.SessionFromContext(c.Request().Context())
	if !ok {
		return c.Redirect(http.StatusFound, "/v1/login")
	}

	req := dto.PostRegisterReq{}

	if err := c.Bind(&req); err != nil {
		return resp.HTTPBadRequest(c, "", err.Error())
	}

	ss := &auth.Session{Claim: sessionClaims}
	u, _ := ss.GetUser(*h.clerk)
	if u == nil {
		return c.Redirect(http.StatusFound, "/v1/login")
	}

	if _, err := h.repo.CreateSpillUser(c.Request().Context(), &db.SpillUser{
		Alias: req.Alias,
		Bio: pgtype.Text{
			String: req.Bio,
			Valid:  true,
		},
		ServiceID: u.ID,
	}); err != nil {
		return errs.Wrap(err)
	}

	c.Response().Header().Add("Hx-Redirect", "/")
	return c.JSON(http.StatusOK, nil)
}
