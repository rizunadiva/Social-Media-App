package usecase

import (
	"errors"
	"fmt"
	"socialmedia-app/domain"
)

type newsUseCase struct {
	newsData domain.NewsData
}

func New(model domain.NewsData) domain.NewsUseCase {
	return &newsUseCase{
		newsData: model,
	}
}

func (nu *newsUseCase) AddNews(IDUser int, newNews domain.News) (domain.News, error) {
	if IDUser == -1 {
		return domain.News{}, errors.New("invalid user")
	}

	newNews.UserID = IDUser
	fmt.Println("news", newNews)
	res := nu.newsData.Insert(newNews)

	if res.ID == 0 {
		return domain.News{}, errors.New("error insert news")
	}
	return res, nil
}

func (nu *newsUseCase) UpNews(IDNews int, updateData domain.News) (domain.News, error) {
	if IDNews == -1 {
		return domain.News{}, errors.New("invalid news")
	}

	// updateData.UserID = IDNews
	res := nu.newsData.Update(IDNews, updateData)
	if res.ID == 0 {
		return domain.News{}, errors.New("error update news")
	}

	return res, nil
}

func (nu *newsUseCase) DelNews(IDNews int) (bool, error) {
	res := nu.newsData.Delete(IDNews)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

// func (nu *newsUseCase) GetAllN() ([]domain.News, error) {
// 	res := nu.newsData.GetAll()

// 	if len(res) == 0 {
// 		return nil, errors.New("no data found")
// 	}

// 	return res, nil
// }

// func (nu *newsUseCase) GetMyN(IDUser int) ([]domain.News, error) {

// 	if IDUser == -1 {
// 		return nil, errors.New("invalid user")
// 	}

// 	res := nu.data.GetMy(IDUser)

// 	return res, nil

// }
