package domain

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	ID             int
	Content        string
	Images         string
	FileAttachment string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type NewsUseCase interface {
	AddNews(newNews News) (News, error)
}

type NewsData interface {
	Insert(newNews News) (News, error)
}
