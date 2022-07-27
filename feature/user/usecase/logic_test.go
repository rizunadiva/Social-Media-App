package usecase

import (
	"socialmedia-app/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	// MockTrue
	mockData := domain.User{ID: 1, UserName: "antonio", Email: "antonio@gmail.com"}
	mockData2 := domain.User{}
	t.Run("Success case", func(t *testing.T) {
		useCase := New(&mockUserDataTrue{})
		res, err := useCase.AddUser(mockData)
		assert.Nil(t, err)
		assert.Greater(t, res.ID, 0)
		assert.Equal(t, "antonio", res.FullName)
	})

	t.Run("Cannot insert data", func(t *testing.T) {
		useCase := New(&mockUserDataFalse{})
		res, err := useCase.AddUser(mockData2)
		assert.NotNil(t, err)
		assert.Equal(t, 0, res.ID)
	})
}

type mockUserDataTrue struct{}

func (mud *mockUserDataTrue) AddUser(newUser domain.User) domain.User {
	newUser.ID = 1
	return newUser
}

func (mud *mockUserDataTrue) GetAll() []domain.User {
	return []domain.User{{ID: 1, UserName: "Antonio", Email: "antonio@gmail.com", Password: "XYZ1234"}}
}

func (mud *mockUserDataTrue) Delete(userID int) bool {
	return true
}

func (mud *mockUserDataTrue) UpdateUser(userID int, updatedData domain.User) domain.User {
	return domain.User{}
}

func (mud *mockUserDataTrue) GetSpecific(userID int) domain.User {
	return domain.User{}
}

type mockUserDataFalse struct{}

func (mudf *mockUserDataFalse) Insert(newUser domain.User) domain.User {
	return domain.User{}
}
func (mudf *mockUserDataFalse) Delete(userID int) bool {
	return false
}

func (mudf *mockUserDataFalse) GetAll() []domain.User {
	return nil
}

func (mudf *mockUserDataFalse) GetSpecific(userID int) domain.User {
	return domain.User{}
}

func (mudf *mockUserDataFalse) UpdateUser(userID int, updatedData domain.User) domain.User {
	return domain.User{}
}
