package data

import (
	"fmt"
	"log"
	"socialmedia-app/domain"

	"golang.org/x/crypto/bcrypt"
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

func (ud *userData) Login(userLogin domain.User) (row int, data domain.User, err error) {
	var dataUser = FromModel(userLogin)
	password := dataUser.Password

	result := ud.db.Where("email = ?", dataUser.Email).First(&dataUser)

	if result.Error != nil {
		return 0, domain.User{}, result.Error
	}

	if result.RowsAffected != 1 {
		return -1, domain.User{}, fmt.Errorf("failed to login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dataUser.Password), []byte(password))

	if err != nil {
		return -2, domain.User{}, fmt.Errorf("failed to login")
	}

	return int(result.RowsAffected), dataUser.ToModel(), nil
}

// func (ud *userData) GetSpecific(userID int) (domain.User, error)
// func (ud *userData) Update(userID int, updatedData domain.User) domain.User
// func (ud *userData) Delete(userID int) (row int, err error)
// func (ud *userData) GetAll() ([]domain.User, error)
