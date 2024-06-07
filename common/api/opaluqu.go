package api

import (
	"mayi/go-proxy-bingai/common"
	"mayi/go-proxy-bingai/common/helper"
	"net/http"
)

func Opaluqu(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	common.NewSingleHostReverseProxy(common.BING_SR_URT).ServeHTTP(w, r)
}
