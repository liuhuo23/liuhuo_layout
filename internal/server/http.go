package server

import "github.com/labstack/echo/v4"

func NewEcho() *echo.Echo {
	e := echo.New()
	return e
}
