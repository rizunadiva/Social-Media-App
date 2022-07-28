package domain

import (
	"time"

	"github.com/labstack/echo"
)

type Comments struct {
	ID        int
	Content   string
	CreatedAt time.Time
	DeletedAt time.Time
	NewsID    int
	News      News
}

type CommentsHandler interface {
	InsertComments() echo.HandlerFunc
	DeleteComments() echo.HandlerFunc
	GetAllComments() echo.HandlerFunc
}

type CommentsUseCase interface {
	GetAllC() ([]Comments, error)
	AddComments(IDUser int, newComments Comments) (Comments, error)
	DelComments(IDComments int) (bool, error)
}

type CommentsData interface {
	GetAll() []Comments
	Insert(newComments Comments) Comments
	Delete(IDComments int) bool
}
