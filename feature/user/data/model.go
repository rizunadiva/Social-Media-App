package data

import (
	"socialmedia-app/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required,username"`
	Password string `json:"password" form:"email" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Photo    string `json:"image_url"`
	// News     []data.News `gorm:"foreignKey:By"`
}

func (u *User) ToModel() domain.User {
	return domain.User{
		ID:        int(u.ID),
		UserName:  u.Username,
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
	res.Password = data.Password
	res.FullName = data.FullName
	res.Photo = data.Photo
	// res.ID = data.ID
	return res
}
