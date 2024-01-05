package config

import (
	"prototype/app/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(inject Injection) *Router {
	route := gin.Default()

	route.Use(middleware.Logging(inject.Logging))

	route.GET("/version", HandleVersion)

	v1 := route.Group("v1")
	{
		v1.GET("/user", inject.UserController.Fetch)
		v1.GET("/user/:user_id", inject.UserController.GetByID)
		v1.POST("/user", inject.UserController.Create)
		v1.PUT("/user/:user_id", inject.UserController.Update)
		v1.DELETE("/user/:user_id", inject.UserController.Delete)
	}

	return &Router{route}
}
