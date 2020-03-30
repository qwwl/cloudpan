package main

import "github.com/labstack/echo/v4"

type Inter interface {
	Get(c echo.Context) error
	Post(c echo.Context) error
	Put(c echo.Context) error
	Delete(c echo.Context) error
	Gets(c echo.Context) error
}
