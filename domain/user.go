package domain

import "time"

type User struct {
	ID        int
	UserName  string
	Password  string
	FullName  string
	Photo     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUseCase interface {
	AddUser(newUser User) (User, error)
	GetAll() ([]User, error)
	GetProfile(id int) (User, error)
	UpdateUser(id int, updateProfile User) User
	DeleteUser(id int) (row int, err error)
	LoginUser(email string, password string) (token string, nama string, id int, errLog error)
}

type UserData interface {
	Insert(newUser User) (User, error)
	GetAll() ([]User, error)
	GetSpecific(userID int) (User, error)
	Update(userID int, updatedData User) User
	Delete(userID int) (row int, err error)
	Login(email string, password string) (token string, nama string, id int, err error)
}
