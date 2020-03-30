package main

import (
	"strconv"

	echo "github.com/labstack/echo/v4"
)

func MakePageAndLimit(c echo.Context) (int, int) {
	return MustAtoI(c.Param("page")), MustAtoI(c.Param("limit"))
}

func MustAtoI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}
