package delivery

import "socialmedia-app/domain"

type NewsInsertRequest struct {
	Content        string `json:"content" form:"content"`
	Images         string `json:"images" form:"images"`
	FileAttachment string `json:"file" form:"file"`
	UserId         uint   `json:"user_id" form:"user_id"`
}

func (ni *NewsInsertRequest) ToDomain() domain.News {
	return domain.News{
		Content:        ni.Content,
		Images:         ni.Images,
		FileAttachment: ni.FileAttachment,
	}
}
