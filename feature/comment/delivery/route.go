package delivery

import (
	"socialmedia-app/config"
	"socialmedia-app/domain"
	"socialmedia-app/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteComments(e *echo.Echo, ch domain.CommentsHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.GET("/comments", ch.GetAllComments())
	// e.GET("/comments", ch.GetMyComments(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.POST("/comments", ch.InsertComments(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/comments/:id", ch.DeleteComments(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))

}
