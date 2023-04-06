package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello/", hello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		handler404(w, r)
		return
	}
	w.Write([]byte("корень"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	pathRegexp := regexp.MustCompile(`^/hello/\w+$`)
	if !pathRegexp.Match([]byte(r.URL.Path)) {
		handler404(w, r)
		return
	}

	name := strings.Split(r.URL.Path, "/")[2]
	w.Write([]byte(fmt.Sprintf("хелп %s", name)))
}

func handler404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 not page found"))
}
