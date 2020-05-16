package v1

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/ngotzmann/gommon"
	"github.com/ngotzmann/gorror"
)

func EchoHandler(configPath string) *echo.Echo {
	c := gommon.NewConfig(configPath)
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.RequestID())
	e.HTTPErrorHandler = gorror.CustomEchoHTTPErrorHandler
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(c.Webservice.SessionSecret))))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	//Declaration of routes
	e.POST("/unrestricted", Unrestricted)
	e.GET("/unrestricted", Unrestricted)

	return e
}
