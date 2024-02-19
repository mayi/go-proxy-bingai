package api

import (
	"mayi/go-proxy-bingai/common"
	"net/http"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if common.SB {
			w.WriteHeader(http.StatusUnavailableForLegalReasons)
			return
		}
		next(w, r)
	})
}
