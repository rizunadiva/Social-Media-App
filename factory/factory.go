package factory

import (
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

	// bookData := bd.New(db)
	// bookCase := bs.New(bookData)
	// bookHandler := bookDelivery.New(bookCase)
	// bookDelivery.RouteBook(e, bookHandler)
}
