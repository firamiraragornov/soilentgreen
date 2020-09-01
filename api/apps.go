package handler

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Handler(res http.ResponseWriter, req *http.Request) {
    url, _ := url.Parse("http://pmiappserverprod-env.eba-i3hnze8t.us-east-2.elasticbeanstalk.com/")
    httputil.NewSingleHostReverseProxy(url).ServeHTTP(res, req)
}