package clip

import (
	"net/http"
	"time"
)

type Bookmark struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func bookmarkCreateHandler(w http.ResponseWriter, r *http.Request) {
}
