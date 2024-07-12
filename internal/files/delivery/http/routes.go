package http

import (
	"github.com/labstack/echo/v4"

	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/files"
	"github.com/dhuruvah-apps/BuildSMEs-Api/internal/middleware"
)

// Map comments routes
func MapFilesRoutes(commGroup *echo.Group, h files.Handlers, mw *middleware.MiddlewareManager) {
	commGroup.POST("", h.Create(), mw.AuthSessionMiddleware, mw.CSRF)
	commGroup.DELETE("/:file_id", h.Delete(), mw.AuthSessionMiddleware, mw.CSRF)
	commGroup.PUT("/:file_id", h.Update(), mw.AuthSessionMiddleware, mw.CSRF)
	commGroup.GET("/:file_id", h.GetByID())
	commGroup.GET("", h.GetAll())
}
