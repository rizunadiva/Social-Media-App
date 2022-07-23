package delivery

import "socialmedia-app/domain"

type InsertFormat struct {
	UserName string `json:"username" form:"username"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	FullName string `json:"fullname" form:"fullname"`
	Photo    string `json:"photo" form:"photo"`
}

func (i *InsertFormat) ToModel() domain.User {
	return domain.User{
		UserName: i.UserName,
		Email:    i.Email,
		Password: i.Password,
		FullName: i.FullName,
		Photo:    i.Photo,
	}
}
