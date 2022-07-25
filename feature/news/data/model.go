package data

import (
	"socialmedia-app/domain"
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Content        string `json:"content" form:"content"`
	Images         string `json:"images" form:"images"`
	FileAttachment string `json:"file" form:"file"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (b *News) ToDomain() domain.News {
	return domain.News{
		ID:             int(b.ID),
		Content:        b.Content,
		Images:         b.Images,
		FileAttachment: b.FileAttachment,
	}
}

func ParseToArrDomain(arr []News) []domain.News {
	var res []domain.News

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}

	return res
}

func ToLocal(data domain.News) News {
	var res News
	res.ID = uint(data.ID)
	res.Content = data.Content
	res.Images = data.Images
	res.FileAttachment = data.FileAttachment
	return res
}
