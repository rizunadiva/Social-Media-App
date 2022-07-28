package delivery

import (
	"socialmedia-app/config"
	"socialmedia-app/domain"
	"socialmedia-app/feature/common"

	// "socialmedia-app/feature/news/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteBook(e *echo.Echo, bc domain.NewsHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/news", bc.InsertNews(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/news/:id", bc.UpdateNews(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/news/:id", bc.DeleteNews(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/news", bc.GetAllNews())
	e.GET("/news/:id", bc.GetNewsID())
}
