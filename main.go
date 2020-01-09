package main

import (
	"net/http"
)

func handleDefault(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/" {
		http.ServeFile(w, req, "www/index.html")
		return
	}
	req.URL.RawQuery = "" // skip "?params"
	wwwServer.ServeHTTP(w, req)
}

var (
	wwwServer = http.FileServer(http.Dir("www"))
)

func main() {
	http.HandleFunc("/", handleDefault)
	http.ListenAndServe(":80", nil)
}
