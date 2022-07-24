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
	// GetAll() ([]User, error)
	// GetProfile(id int) (User, error)
	// UpdateUser(id int, updateProfile User) User
	// DeleteUser(id int) (row int, err error)
}

type UserData interface {
	Insert(newUser User) (User, error)
	Login(userLogin User) (row int, data User, err error)
	// GetAll() ([]User, error)
	// GetSpecific(userID int) (User, error)
	// Update(userID int, updatedData User) User
	// Delete(userID int) (row int, err error)
}
