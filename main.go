package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	proxy := goproxy.NewProxyHttpServer()

	proxy.OnRequest(goproxy.DstHostIs("makki0205.com")).DoFunc(func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		return r, goproxy.NewResponse(r, goproxy.ContentTypeHtml, http.StatusOK, "<html lang='ja'>"+
			"<head>"+
			"<meta charset='utf-8'>"+
			"</head>"+
			"<body>"+
			"<h1>まっきー寝てやんの〜〜〜〜〜〜！！！</h1>"+
			"</body>"+
			"</html>")
	})

	fmt.Println("proxy server is started!")
	log.Fatal(http.ListenAndServe(":8000", proxy))
}
