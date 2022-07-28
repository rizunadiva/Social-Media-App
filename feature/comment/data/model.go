package data

import (
	"socialmedia-app/domain"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Content string `json:"content" form:"content"`
	NewsID  uint   `json:"news_id" form:"news_id"`
	// News    News
}

// type News struct {
// 	gorm.Model
// 	Content        string `json:"content" form:"content"`
// 	Images         string `json:"images" form:"images"`
// 	FileAttachment string `json:"file" form:"file"`
// 	UserID         uint   `json:"user_id" form:"user_id"`
// 	User           User
// }
// type User struct {
// 	gorm.Model
// 	Username string `json:"username" form:"username" validate:"required"`
// 	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
// 	Password string `json:"password" form:"password" validate:"required"`
// 	FullName string `json:"fullname" form:"fullname" validate:"required"`
// 	Photo    string `json:"image_url"`
// }

func (c *Comments) ToDomain() domain.Comments {
	return domain.Comments{
		ID:        int(c.ID),
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		NewsID:    int(c.NewsID),
	}
}

func ParseToArrDomain(arr []Comments) []domain.Comments {
	var res []domain.Comments

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}

	return res
}

func ToLocal(data domain.Comments) Comments {
	var res Comments
	res.ID = uint(data.ID)
	res.NewsID = uint(data.NewsID)
	res.Content = data.Content
	return res
}
