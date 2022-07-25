package data

import (
	"socialmedia-app/domain"
	"socialmedia-app/feature/news/data"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string      `json:"username" form:"username" validate:"required"`
	Email    string      `gorm:"unique" json:"email" form:"email" validate:"required"`
	Password string      `json:"password" form:"password" validate:"required"`
	FullName string      `json:"fullname" form:"fullname" validate:"required"`
	Photo    string      `json:"image_url"`
	News     []data.News `gorm:"foreignKey:By"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:        int(u.ID),
		UserName:  u.Username,
		Email:     u.Email,
		Password:  u.Password,
		FullName:  u.FullName,
		Photo:     u.Photo,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ParseToArr(arr []User) []domain.User {
	var res []domain.User

	for _, val := range arr {
		res = append(res, val.ToModel())
	}

	return res
}

func FromModel(data domain.User) User {
	var res User
	res.Username = data.UserName
	res.Email = data.Email
	res.Password = data.Password
	res.FullName = data.FullName
	res.Photo = data.Photo
	// res.ID = data.ID
	return res
}
