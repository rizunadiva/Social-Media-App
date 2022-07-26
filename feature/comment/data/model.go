package data

import (
	"socialmedia-app/domain"

	"gorm.io/gorm"
)

type Comments struct {
	gorm.Model
	Content string `json:"content" form:"content"`
	UserID  uint
	User    User
}

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
	Comments []Comments
}

func (c *Comments) ToDomain() domain.Comments {
	return domain.Comments{
		ID:        int(c.ID),
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		Pemilik: domain.User{
			ID:       int(c.User.ID),
			UserName: c.User.Username,
			Email:    c.User.Email,
		},
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
	res.UserID = uint(data.Pemilik.ID)
	res.Content = data.Content
	return res
}
