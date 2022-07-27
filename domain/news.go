package domain

import (
	"time"

	"github.com/labstack/echo"
)

type News struct {
	ID             int
	Content        string
	Images         string
	FileAttachment string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Pemilik        User
}

type NewsHandler interface {
	InsertNews() echo.HandlerFunc
	DeleteNews() echo.HandlerFunc
	UpdateNews() echo.HandlerFunc
	GetAllNews() echo.HandlerFunc
	GetMyNews() echo.HandlerFunc
}

type NewsUseCase interface {
	GetAllN() ([]News, error)
	GetMyN(IDUser int) ([]News, error)
	AddNews(IDUser int, newNews News) (News, error)
	DelNews(IDNews int) (bool, error)
	UpNews(IDNews int, updateData News) (News, error)
}

type NewsData interface {
	GetAll() []News
	GetMy(IDUser int) []News
	Insert(newNews News) News
	Delete(IDNews int) bool
	Update(IDNews int, updatedNews News) News
}
