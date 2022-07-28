package delivery

import "socialmedia-app/domain"

type commentsInsertRequest struct {
	Content string `json:"content" form:"content"`
	NewsId  uint   `json:"news_id" form:"news_id"`
}

func (ci *commentsInsertRequest) ToDomain() domain.Comments {
	return domain.Comments{
		Content: ci.Content,
	}
}
