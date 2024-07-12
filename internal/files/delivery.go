package files

import "github.com/labstack/echo/v4"

// Comments HTTP Handlers interface
type Handlers interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetByID() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}
