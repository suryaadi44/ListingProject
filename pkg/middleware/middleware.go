package middleware

import (
	"net/http"

	global "github.com/suryaadi44/ListingProject/pkg/dto"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				switch response := r.(type) {
				case *global.BaseResponse:
					global.NewBaseResponse(response.Code, true, response.Data).SendResponse(&w)
				case error:
					global.NewBaseResponse(http.StatusInternalServerError, true, response.Error()).SendResponse(&w)
				default:
					global.NewBaseResponse(http.StatusInternalServerError, true, "runtime error").SendResponse(&w)
				}
			}
		}()
		next.ServeHTTP(w, r)
	})
}
