package delivery

import (
	"socialmedia-app/config"
	"socialmedia-app/domain"
	"socialmedia-app/feature/news/middlewares"

	"github.com/labstack/echo/"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteComments(e *echo.Echo, ch domain.CommentsHandler) {
	e.GET("/allcomments", ch.GetAllComments())
	e.GET("/comments", ch.GetMyComments(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.POST("/comments", ch.InsertComments(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/comments/:id", ch.DeleteComments(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))

}
