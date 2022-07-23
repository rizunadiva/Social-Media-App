package data

import (
	"log"
	"socialmedia-app/domain"

	"gorm.io/gorm"
)

type userData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.UserData {
	return &userData{
		db: db,
	}
}

func (ud *userData) Insert(newUser domain.User) (domain.User, error) {
	var cnv = FromModel(newUser)
	err := ud.db.Create(&cnv).Error
	if err != nil {
		log.Println("Cannot create object", err.Error())
		return domain.User{}, err
	}

	return cnv.ToModel(), nil
}
func (ud *userData) Login(email string, password string) (token string, nama string, id int, err error)
func (ud *userData) GetSpecific(userID int) (domain.User, error)
func (ud *userData) Update(userID int, updatedData domain.User) domain.User
func (ud *userData) Delete(userID int) (row int, err error)
func (ud *userData) GetAll() ([]domain.User, error)
