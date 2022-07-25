package data

import (
	"socialmedia-app/domain"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model
	Content        string `json:"content" form:"content"`
	Images         string `json:"images" form:"images"`
	FileAttachment string `json:"file" form:"file"`
	UserID         uint
	User           User
}
type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
	News     []News
}

func (b *News) ToDomain() domain.News {
	return domain.News{
		ID:             int(b.ID),
		Content:        b.Content,
		Images:         b.Images,
		FileAttachment: b.FileAttachment,
		CreatedAt:      b.CreatedAt,
		UpdatedAt:      b.UpdatedAt,
		Pemilik: domain.User{
			ID:       int(b.User.ID),
			UserName: b.User.Username,
			Email:    b.User.Email,
		},
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
	res.UserID = uint(data.Pemilik.ID)
	res.Content = data.Content
	res.Images = data.Images
	res.FileAttachment = data.FileAttachment
	return res
}
