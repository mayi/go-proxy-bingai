package api

import (
	"mayi/go-proxy-bingai/common"
	"mayi/go-proxy-bingai/common/helper"
	"net/http"
	"strings"
)

func Th(w http.ResponseWriter, r *http.Request) {
	if !helper.CheckAuth(r) {
		helper.UnauthorizedResult(w)
		return
	}
	r.URL.Path = strings.ReplaceAll(r.URL.Path, "/th/th", "/th")
	common.NewSingleHostReverseProxy(common.BING_SOURCE_URL).ServeHTTP(w, r)
}
