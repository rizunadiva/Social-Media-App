package factory

import (
	nd "socialmedia-app/feature/news/data"
	newsDelivery "socialmedia-app/feature/news/delivery"
	nu "socialmedia-app/feature/news/usecase"
	ud "socialmedia-app/feature/user/data"
	userDelivery "socialmedia-app/feature/user/delivery"
	us "socialmedia-app/feature/user/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userDelivery.New(e, useCase)

	newsData := nd.New(db)
	newsCase := nu.New(newsData)
	newsHandler := newsDelivery.New(newsCase)
	newsDelivery.RouteBook(e, newsHandler)
}
