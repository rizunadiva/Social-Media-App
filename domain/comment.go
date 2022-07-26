package domain

import (
	"time"

	"github.com/labstack/echo"
)

type Comments struct {
	ID        int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	Pemilik   User
}

type CommentsHandler interface {
	InsertComments() echo.HandlerFunc
	DeleteComments() echo.HandlerFunc
	UpdateComments() echo.HandlerFunc
	GetAllComments() echo.HandlerFunc
	GetMyComments() echo.HandlerFunc
}

type CommentsUseCase interface {
	GetAllC() ([]Comments, error)
	GetMyC(IDUser int) ([]Comments, error)
	AddComments(IDUser int, newComments Comments) (Comments, error)
	DelComments(IDComments int) (bool, error)
	UpComments(IDComments int, updateData Comments) (Comments, error)
}

type CommentsData interface {
	GetAll() []Comments
	GetMy(IDUser int) []Comments
	Insert(newComments Comments) Comments
	Delete(IDComments int) bool
	Update(IDComments int, updatedComments Comments) Comments
}
