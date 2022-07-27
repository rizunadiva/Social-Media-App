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
}

type NewsUseCase interface {
	AddNews(IDUser int, useNews News) (News, error)
}

type NewsData interface {
	Insert(insertNews News) News
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
// 	// 	DeleteNews() echo.HandlerFunc
// 	// 	UpdateNews() echo.HandlerFunc
// 	// 	GetAllNews() echo.HandlerFunc
// 	// 	GetMyNews() echo.HandlerFunc
// }

// type NewsUseCase interface {
// 	AddNews(IDUser int, newNews News) (News, error)
// 	GetAllN() ([]News, error)
// 	// GetMyN(IDUser int) ([]News, error)
// 	// DelNews(IDNews int) (bool, error)
// 	// UpNews(IDNews int, updateData News) (News, error)
// }

// type NewsData interface {
// 	Insert(newNews News) News
// 	GetAll() []News
// 	// GetMy(IDUser int) []News
// 	// Delete(IDNews int) bool
// 	// Update(IDNews int, updatedNews News) News
// }
