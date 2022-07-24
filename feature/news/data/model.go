package data

import (
	"socialmedia-app/domain"
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	ID             int
	Content        string `json:"content" form:"content" validate:"required"`
	Images         string `json:"image_url" form:"Images"`
	FileAttachment string `json:"File" form:"File"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (u *News) ToModel() domain.News {
	return domain.News{
		ID:             int(u.ID),
		Content:        u.Content,
		Images:         u.Images,
		FileAttachment: u.FileAttachment,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}

func ParseToArr(arr []News) []domain.News {
	var res []domain.News

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.News) News {
	var res News
	res.Content = data.Content
	res.Images = data.Images
	res.FileAttachment = data.FileAttachment
	return res
}
