package usecase

import (
	"errors"
	"socialmedia-app/domain"
)

type commentsUseCase struct {
	data domain.CommentsData
}

func New(model domain.CommentsData) domain.CommentsUseCase {
	return &commentsUseCase{
		data: model,
	}
}

func (cu *commentsUseCase) AddComments(IDUser int, newComments domain.Comments) (domain.Comments, error) {
	if IDUser == -1 {
		return domain.Comments{}, errors.New("invalid user")
	}

	newComments.Pemilik = IDUser

	res := cu.data.Insert(newComments)
	if res.ID == 0 {
		return domain.Comments{}, errors.New("error insert news")
	}

	return res, nil
}

func (cu *commentsUseCase) GetAllC() ([]domain.Comments, error) {
	res := cu.data.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (cu *commentsUseCase) GetMyC(IDUser int) ([]domain.Comments, error) {

	if IDUser == -1 {
		return nil, errors.New("invalid user")
	}

	res := cu.data.GetMy(IDUser)

	return res, nil

}

func (cu *commentsUseCase) DelComments(IDComments int) (bool, error) {
	res := cu.data.Delete(IDComments)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}
