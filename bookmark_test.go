package clip_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/allyraza/clip"
)

func TestBookmarkCreateHandler(t *testing.T) {
	tests := []struct {
		method string
		url    string
		status int
	}{
		{
			method: http.MethodPost,
			url:    "/bookmark",
			status: http.StatusOK,
		},
		{
			method: http.MethodGet,
			url:    "/bookmark",
			status: http.StatusMethodNotAllowed,
		},
		{
			method: http.MethodDelete,
			url:    "/bookmark",
			status: http.StatusMethodNotAllowed,
		},
		{
			method: http.MethodPut,
			url:    "/bookmark",
			status: http.StatusMethodNotAllowed,
		},
		{
			method: http.MethodPost,
			url:    "/invalid",
			status: http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		name := fmt.Sprintf("%s %s %d", tc.method, tc.url, tc.status)

		t.Run(name, func(t *testing.T) {
			w := httptest.NewRecorder()

			r, err := http.NewRequest(tc.method, tc.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			s, err := clip.NewServer()
			if err != nil {
				t.Error(err)
			}

			s.ServeHTTP(w, r)

			// bookmarkCreateHandler(w, r)

			if w.Code != tc.status {
				t.Errorf("Want %d, got %d", tc.status, w.Code)
			}
		})
	}
}
