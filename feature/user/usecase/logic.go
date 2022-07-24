package usecase

import (
	"errors"
	"log"
	"socialmedia-app/domain"
	"socialmedia-app/feature/user/data"

	validator "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
func (ud *userUseCase) UpdateUser(id int, updateProfile domain.User) (row int, err error) {
	if updateProfile.Email == "" || updateProfile.Password == "" || updateProfile.UserName == "" || updateProfile.FullName == "" || updateProfile.Photo == "" {
		return -1, errors.New("all input data must be filled")
	}
	row, err = ud.userData.Update(id, updateProfile)
	return row, err
}

func (ud *userUseCase) LoginUser(userLogin domain.User) (response int, data domain.User, err error) {
	response, data, err = ud.userData.Login(userLogin)

	return response, data, err
}

func (ud *userUseCase) GetProfile(id int) (domain.User, error) {
	data, err := ud.userData.GetSpecific(id)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data not found")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}

	return data, nil
}

func (ud *userUseCase) DeleteUser(id int) (row int, err error) {
	row, err = ud.userData.Delete(id)
	if err != nil {
		log.Println("delete usecase error", err.Error())
		if err == gorm.ErrRecordNotFound {
			return row, errors.New("data not found")
		} else {
			return row, errors.New("server error")
		}
	}
	return row, nil
}
