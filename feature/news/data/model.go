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
	UserID         uint   `json:"user_id" form:"user_id"`
	User           User
}

// type User struct {
// 	gorm.Model
// 	Username string `json:"username" form:"username"`
// 	Email    string `gorm:"unique" json:"email" form:"email"`
// 	// News     []News
// }
type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Photo    string `json:"image_url"`
}

func (b *News) ToDomain() domain.News {
	return domain.News{
		ID:             int(b.ID),
		Content:        b.Content,
		Images:         b.Images,
		FileAttachment: b.FileAttachment,
		CreatedAt:      b.CreatedAt,
		UpdatedAt:      b.UpdatedAt,
		// DeletedAt:      b.DeletedAt,
		// PostedBy:       int(b.PostedBy),
		UserID: int(b.UserID),
		User: domain.User{
			ID:       int(b.User.ID),
			UserName: b.User.Username,
		},
	}
}

func ParseToArr(arr []News) []domain.News {
	var res []domain.News

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.News) News {
	var res News
	res.ID = uint(data.ID)
	res.UserID = uint(data.UserID)
	res.Content = data.Content
	res.Images = data.Images
	res.FileAttachment = data.FileAttachment
	return res
}
