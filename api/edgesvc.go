package api

import (
	"mayi/go-proxy-bingai/api/helper"
	"mayi/go-proxy-bingai/common"
	"net/http"
)

func Edgesvc(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	common.NewSingleHostReverseProxy(common.EDGE_SVC_URL).ServeHTTP(w, r)
}
