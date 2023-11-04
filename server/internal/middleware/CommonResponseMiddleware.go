package middleware

import "net/http"

type CommonResponseMiddleware struct {
}

func NewCommonResponseMiddleware() *CommonResponseMiddleware {
	return &CommonResponseMiddleware{}
}

func (m *CommonResponseMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		next(w, r)
	}
}
