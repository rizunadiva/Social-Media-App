package domain

import (
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	IDNews         int
	Content        string
	Images         string
	FileAttachment string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type newsHandler interface {
	InsertNews() echo.HandlerFunc
	DeleteNews() echo.HandlerFunc
	UpdateNews() echo.HandlerFunc
	GetAllNews() echo.HandlerFunc
	GetMyNews() echo.HandlerFunc
}

type NewsUseCase interface {
	GetAllM() ([]News, error)
	GetMyM(IDUser int) ([]News, error)
	AddBook(IDUser int, newNews News) (News, error)
	DelBook(IDNews int) (bool, error)
	UpBook(IDNews int, updateData News) (News, error)
}

type NewsData interface {
	GetAll() []News
	GetMy(IDUser int) []News
	Insert(newNews News) News
	Delete(IDNews int) bool
	Update(IDNews int, updatedNews News) News
}
