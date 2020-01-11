package clip

import "net/http"

type server struct {
	mux *http.ServeMux
}

func (s *server) init() {
	s.mux.HandleFunc("/bookmark", bookmarkCreateHandler)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func NewServer() (*server, error) {
	s := &server{
		mux: http.NewServeMux(),
	}

	s.init()

	return s, nil
}
