package main

import (
	"flag"
	"github.com/NYTimes/gziphandler"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
)

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prefix := "../../assets"
	path := r.URL.Path
	if path == "/" {
		path = prefix + "/index.html"
	} else {
		path = prefix + path
	}
	data, err := ioutil.ReadFile(path)

	if err == nil {
		_, _ = w.Write(data)
	} else {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
	}
}

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	http.Handle("/", gziphandler.GzipHandler(new(MyHandler)))
	log.Fatal(http.ListenAndServe(*listen, nil))
}
