package data

import (
	"fmt"
	"log"

	// "log"
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

func (nd *newsData) Insert(dataNews domain.News) domain.News {
	// fmt.Println("data :", dataNews)
	cnv := ToLocal(dataNews)
	// fmt.Println("cnv", cnv)
	err := nd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.News{}
	}
	return cnv.ToDomain()
}

func (bd *newsData) Update(newsID int, updatedNews domain.News) domain.News {
	cnv := ToLocal(updatedNews)
	err := bd.db.Model(cnv).Where("ID = ?", newsID).Updates(updatedNews)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.News{}
	}
	cnv.ID = uint(newsID)
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

func (nd *newsData) GetAll() []domain.News {
	var data []News
	err := nd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (nd *newsData) GetNewsID(newsID int) []domain.News {
	var data []News
	err := nd.db.Where("ID = ?", newsID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
