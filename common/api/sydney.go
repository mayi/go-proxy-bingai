package api

import (
	"mayi/go-proxy-bingai/common"
	"mayi/go-proxy-bingai/common/helper"
	"net/http"
)

func Sydney(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	common.NewSingleHostReverseProxy(common.BING_SYDNEY_URL).ServeHTTP(w, r)
}
