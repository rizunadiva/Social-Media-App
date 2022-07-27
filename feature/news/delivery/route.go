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
	e.POST("/news", bc.InsertNews(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
}
