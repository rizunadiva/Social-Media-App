package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Comments struct {
	ID        int
	Content   string
	CreatedAt time.Time
	DeletedAt time.Time
	NewsID    int
	// News      News
}

type CommentsHandler interface {
	InsertComments() echo.HandlerFunc
	DeleteComments() echo.HandlerFunc
	GetAllComments() echo.HandlerFunc
}

type CommentsUseCase interface {
	AddComments(IDNews int, newComments Comments) (Comments, error)
	GetAllC() ([]Comments, error)
	DelComments(IDComments int) (bool, error)
}

type CommentsData interface {
	Insert(newComment Comments) Comments
	GetAll() []Comments
	Delete(IDComments int) bool
}
