package data

import (
	"log"
	"socialmedia-app/domain"

	"gorm.io/gorm"
)

type newsData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.NewsData {
	return &newsData{
		db: db,
	}
}

func (nd *newsData) GetAll() []domain.News {
	var data []News
	err := nd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArrDomain(data)
}

func (nd *newsData) GetMy(userID int) []domain.News {
	var data []News
	err := nd.db.Where("Pemilik = ?", userID).Find(&data)

	if err.Error != nil {
		log.Println("There is problem with data", err.Error.Error())
		return nil
	}
	return ParseToArrDomain(data)
}

func (nd *newsData) Insert(newNews domain.News) domain.News {
	cnv := ToLocal(newNews)
	err := nd.db.Create(&cnv)
	if err.Error != nil {
		log.Println("Cannot insert data", err.Error.Error())
		return domain.News{}
	}

	return cnv.ToDomain()
}

func (nd *newsData) Delete(newsID int) bool {
	err := nd.db.Where("ID = ?", newsID).Delete(&News{})
	if err.Error != nil {
		log.Println("Cannot delete data", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error.Error())
		return false
	}
	return true
}

func (bd *newsData) Update(newsID int, updatedNews domain.News) domain.News {
	cnv := ToLocal(updatedNews)
	err := bd.db.Model(&cnv).Where("ID = ?", newsID).Updates(updatedNews)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.News{}
	}
	cnv.ID = uint(newsID)
	return cnv.ToDomain()
}
