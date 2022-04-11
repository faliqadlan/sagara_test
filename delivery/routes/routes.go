package routes

import (
	"be/delivery/controllers/auth"
	"be/delivery/controllers/user"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, ac *auth.Controller, uc *user.Controller) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}",
	}))
	// e.AcquireContext().Cookies()
	/* no jwt */

	e.GET("/test", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<strong>Hello, World!</strong>")
	})

	// login ====================================

	e.POST("/login", ac.Login())

	// user ====================================

	e.POST("/user", uc.Create())
}
