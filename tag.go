package clip

import (
	"time"
)

type Tag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TagService interface {
	Create(name string) error
	Delete(id int) error
}
