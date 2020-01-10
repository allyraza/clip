package clip

import "net/http"

type server struct {
	mux *http.ServeMux
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func NewServer() (*server, error) {
	s := &server{}

	return s, nil
}
