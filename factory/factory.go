package factory

import (
	nd "socialmedia-app/feature/news/data"
	newsDelivery "socialmedia-app/feature/news/delivery"
	nu "socialmedia-app/feature/news/usecase"

	ud "socialmedia-app/feature/user/data"
	userDelivery "socialmedia-app/feature/user/delivery"
	us "socialmedia-app/feature/user/usecase"

	cd "socialmedia-app/feature/comment/data"
	commentDelivery "socialmedia-app/feature/comment/delivery"
	cu "socialmedia-app/feature/comment/usecase"

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

	commentsData := cd.New(db)
	commentsCase := cu.New(commentsData)
	commentsHandler := commentDelivery.New(commentsCase)
	commentDelivery.RouteComments(e, commentsHandler)
}
