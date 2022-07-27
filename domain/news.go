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
	UserID         int
	User           User
}

type NewsHandler interface {
	InsertNews() echo.HandlerFunc
	UpdateNews() echo.HandlerFunc
	DeleteNews() echo.HandlerFunc
}

type NewsUseCase interface {
	AddNews(IDUser int, useNews News) (News, error)
	UpNews(IDNews int, updateData News) (News, error)
	DelNews(IDNews int) (bool, error)
}

type NewsData interface {
	Insert(insertNews News) News
	Update(IDNews int, updatedNews News) News
	Delete(IDNews int) bool
}

// type News struct {
// 	ID             int
// 	Content        string `json:"content" form:"content"`
// 	Images         string `json:"images" form:"images"`
// 	FileAttachment string `json:"file" form:"file"`
// 	CreatedAt      time.Time
// 	UpdatedAt      time.Time
// 	PostedBy       int
// 	// PostedBy User
// }

// type NewsHandler interface {
// 	InsertNews() echo.HandlerFunc
// 	// 	GetAllNews() echo.HandlerFunc
// 	// 	GetMyNews() echo.HandlerFunc
// }

// type NewsUseCase interface {
// 	AddNews(IDUser int, newNews News) (News, error)
// 	GetAllN() ([]News, error)
// 	// GetMyN(IDUser int) ([]News, error)
// }

// type NewsData interface {
// 	Insert(newNews News) News
// 	GetAll() []News
// 	// GetMy(IDUser int) []News
// }
