package usecase

import (
	"errors"
	"socialmedia-app/domain"
)

type newsUseCase struct {
	data domain.NewsData
}

func New(model domain.NewsData) domain.NewsUseCase {
	return &newsUseCase{
		data: model,
	}
}

func (nu *newsUseCase) AddNews(IDUser int, newNews domain.News) (domain.News, error) {
	if IDUser == -1 {
		return domain.News{}, errors.New("invalid user")
	}

	newNews.Pemilik = IDUser

	res := nu.data.Insert(newNews)
	if res.ID == 0 {
		return domain.News{}, errors.New("error insert news")
	}

	return res, nil
}

func (nu *newsUseCase) GetAllN() ([]domain.News, error) {
	res := nu.data.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (nu *newsUseCase) GetMyN(IDUser int) ([]domain.News, error) {

	if IDUser == -1 {
		return nil, errors.New("invalid user")
	}

	res := nu.data.GetMy(IDUser)

	return res, nil

}

func (nu *newsUseCase) DelNews(IDNews int) (bool, error) {
	res := nu.data.Delete(IDNews)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}
