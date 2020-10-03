package serverMiddleware

import "net/http"

type serverMiddleware struct {}

type ServerMiddleware interface {
	CORS(next http.Handler) http.Handler
}

func NewServerMiddleware() ServerMiddleware {
	return &serverMiddleware{}
}

func (s *serverMiddleware)CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "DELETE, POST, GET, PUT,OPTIONS")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
