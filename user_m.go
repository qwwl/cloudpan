package main

import (
	echo "github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID       string
	Username string
	Password string
	Name     string
	Sex      bool
	Address  string
}

func init() {
	ops = append(ops, AddUserRouter)
}

func AddUserRouter(e *echo.Echo) {
	g := e.Group("/users")
	g.GET("/:id", User{}.Get)
	g.POST("", User{}.Post)
	g.PUT("/:id", User{}.Put)
	g.DELETE("/:id", User{}.Delete)
	g.GET("", User{}.Gets)
}

func (u User) Get(c echo.Context) error {
	cu := new(User)
	cu.ID = c.Param("id")
	if err := Database().First(cu).Error; err != nil {
		logger.Error().Msgf("get user by id: %s error: %s\n", cu.ID, err.Error())
		return c.JSON(200, NewResultError(500, err))
	}
	return c.JSON(200, NewResultSuccess(200, "get user success", cu))
}

func (u User) Post(c echo.Context) error {
	cu := new(User)
	if err := c.Bind(cu); err != nil {
		logger.Error().Msgf("post user by data: %+v error: %s\n", cu, err.Error())
		return c.JSON(200, NewResultError(400, err))
	}
	cu.ID = uuid.NewV4().String()
	if err := Database().Create(cu).Error; err != nil {
		logger.Error().Msg(err.Error())
		return c.JSON(200, NewResultError(500, err))
	}
	return c.JSON(200, NewResultSuccess(200, "post user success", cu))
}

func (u User) Put(c echo.Context) error {
	cu := new(User)
	cu.ID = c.Param("id")
	if err := c.Bind(cu); err != nil {
		logger.Error().Msg(err.Error())
		return c.JSON(200, NewResultError(400, err))
	}
	if err := Database().Model(u).Updates(cu).Error; err != nil {
		logger.Error().Msgf("put user by data: %+v error: %s\n", cu, err.Error())
		return c.JSON(200, NewResultError(500, err))
	}
	return c.JSON(200, NewResultSuccess(200, "put user success", cu))
}

func (u User) Delete(c echo.Context) error {
	cu := new(User)
	cu.ID = c.Param("id")
	if err := c.Bind(cu); err != nil {
		logger.Error().Msg(err.Error())
		return c.JSON(200, NewResultError(400, err))
	}
	if err := Database().Delete(cu).Error; err != nil {
		logger.Error().Msgf("delete user by id: %s error: %s\n", cu.ID, err.Error())
		return c.JSON(200, NewResultError(500, err))
	}
	return c.JSON(200, NewResultSuccess(200, "put user success", cu))
}

func (u User) Gets(c echo.Context) error {
	cu := new([]User)
	if err := c.Bind(cu); err != nil {
		logger.Error().Msg(err.Error())
		return c.JSON(200, NewResultError(400, err))
	}
	p, l := MakePageAndLimit(c)
	cus := make([]User, 0, l)
	if err := Database().Where(cu).Find(cus).Limit(l).Offset((p - 1) * l).Error; err != nil {
		logger.Error().Msgf("gets user by params: %v error: %s\n", cu, err.Error())
		return c.JSON(200, NewResultError(500, err))
	}
	return c.JSON(200, NewResultSuccess(200, "put user success", cus))
}
