package data

import (
	"log"
	"socialmedia-app/domain"

	"gorm.io/gorm"
)

type commentsData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.CommentsData {
	return &commentsData{
		db: db,
	}
}

func (cd *commentsData) Insert(newComment domain.Comments) domain.Comments {
	cnv := ToLocal(newComment)
	err := cd.db.Create(&cnv)
	if err.Error != nil {
		log.Println("Cannot insert data")
		return domain.Comments{}
	}
	return cnv.ToDomain()
}

func (cd *commentsData) GetAll() []domain.Comments {
	var data []Comments
	err := cd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArrDomain(data)
}

func (cd *commentsData) Delete(CommentID int) bool {
	err := cd.db.Where("ID = ?", CommentID).Delete(&Comments{})
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

// func (cd *commentsData) GetMy(userID int) []domain.Comments {
// 	var data []Comments
// 	err := cd.db.Where("Pemilik = ?", userID).Find(&data)

// 	if err.Error != nil {
// 		log.Println("There is problem with data", err.Error.Error())
// 		return nil
// 	}
// 	return ParseToArrDomain(data)
// }
