package delivery

import "socialmedia-app/domain"

type NewsResponse struct {
	ID             int    `json:"id"`
	Content        string `json:"content"`
	Images         string `json:"images"`
	FileAttachment string `json:"file"`
	UserID         int    `json:"user_id"`
}

func FromDomain(data domain.News) NewsResponse {
	var res NewsResponse
	res.ID = int(data.ID)
	res.UserID = int(data.UserID)
	res.Content = data.Content
	res.Images = data.Images
	res.FileAttachment = data.FileAttachment
	return res
}
