package delivery

import "socialmedia-app/domain"

type CommentsResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	NewsID  int    `json:"news_id"`
}

func FromDomain(data domain.Comments) CommentsResponse {
	var res CommentsResponse
	res.ID = data.ID
	res.Content = data.Content
	res.NewsID = data.NewsID
	return res
}
