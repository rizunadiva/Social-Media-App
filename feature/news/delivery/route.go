package delivery

import (
	"socialmedia-app/config"
	"socialmedia-app/domain"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RouteNews(e *echo.Echo, nh domain.NewsHandler) {
	e.GET("/allnews", nh.GetAllNews())
	e.GET("/news", nh.GetMyNews(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.POST("/news", nh.InsertNews(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/news/:id", nh.DeleteNews(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/news/:id", nh.UpdateNews(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
