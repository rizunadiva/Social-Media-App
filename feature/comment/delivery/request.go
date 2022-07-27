package delivery

import "socialmedia-app/domain"

type commentsInsertRequest struct {
	Content string `json:"content" form:"content"`
}

func (ci *commentsInsertRequest) ToDomain() domain.Comments {
	return domain.Comments{
		Content: ci.Content,
	}
}
