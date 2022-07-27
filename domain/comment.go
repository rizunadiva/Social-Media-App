package domain

import "time"

type Comments struct {
	ID             int
	Content        string
	Images         string
	FileAttachment string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	NewsID         int
	News           News
}
