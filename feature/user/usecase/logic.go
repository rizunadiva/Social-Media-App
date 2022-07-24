package usecase

import (
	"errors"
	"log"
	"socialmedia-app/domain"
	"socialmedia-app/feature/user/data"

	validator "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userData domain.UserData
	validate *validator.Validate
}

func New(ud domain.UserData, v *validator.Validate) domain.UserUseCase {
	return &userUseCase{
		userData: ud,
		validate: v,
	}
}

func (ud *userUseCase) AddUser(newUser domain.User) (domain.User, error) {
	var cnv = data.FromModel(newUser)
	err := ud.validate.Struct(cnv)
	if err != nil {
		log.Println("Validation error : ", err.Error())
		return domain.User{}, err
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error encrypt password", err)
		return domain.User{}, err
	}
	newUser.Password = string(hashed)
	inserted, err := ud.userData.Insert(newUser)

	if err != nil {
		log.Println("error from usecase", err.Error())
		return domain.User{}, err
	}
	if inserted.ID == 0 {
		return domain.User{}, errors.New("cannot insert data")
	}
	return inserted, nil
}

func (ud *userUseCase) LoginUser(userLogin domain.User) (response int, data domain.User, err error) {
	response, data, err = ud.userData.Login(userLogin)

	return response, data, err
}

// func (ud *userUseCase) GetAll() ([]domain.User, error)
// func (ud *userUseCase) GetProfile(id int) (domain.User, error)
// func (ud *userUseCase) UpdateUser(id int, updateProfile domain.User) domain.User
// func (ud *userUseCase) DeleteUser(id int) (row int, err error)
