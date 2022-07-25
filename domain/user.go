package domain

import "time"

type User struct {
	ID        int
	UserName  string
	Email     string
	Password  string
	FullName  string
	Photo     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserUseCase interface {
	AddUser(newUser User) (User, error)
	LoginUser(userLogin User) (row int, data User, err error)
	GetProfile(id int) (User, error)
	DeleteUser(id int) (row int, err error)
	UpdateUser(id int, updateProfile User) (User, error)
	// GetAll() ([]User, error)
}

type UserData interface {
	Insert(newUser User) (User, error)
	Login(userLogin User) (row int, data User, err error)
	GetSpecific(userID int) (User, error)
	Delete(userID int) (row int, err error)
	Update(userID int, updatedData User) User
	// GetAll() ([]User, error)
}
