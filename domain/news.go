package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type News struct {
	ID             int
	Content        string
	Images         string
	FileAttachment string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	UserID         int
	User           User
}

type NewsHandler interface {
	InsertNews() echo.HandlerFunc
	UpdateNews() echo.HandlerFunc
	DeleteNews() echo.HandlerFunc
	GetAllNews() echo.HandlerFunc
}

type NewsUseCase interface {
	AddNews(IDUser int, useNews News) (News, error)
	UpNews(IDNews int, updateData News) (News, error)
	DelNews(IDNews int) (bool, error)
	GetAllN() ([]News, error)
}

type NewsData interface {
	Insert(insertNews News) News
	Update(IDNews int, updatedNews News) News
	Delete(IDNews int) bool
	GetAll() []News
}
