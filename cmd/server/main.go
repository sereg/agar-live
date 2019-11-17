package main

import (
	"flag"
	"fmt"
	"github.com/NYTimes/gziphandler"
	_  "github.com/bradrydzewski/go-mimetype"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
)

type staticHandler struct {
}

func (m *staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prefix := "./assets/public"
	path := r.URL.Path
	if path == "/" {
		path = prefix + "/index.html"
	} else {
		path = prefix + path
	}
	data, err := ioutil.ReadFile(path)
	if err == nil {
		mimeType := getMimeType(path)
		w.Header().Set("Content-Type", mimeType)
		_, _ = w.Write(data)
	} else {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
	}
}

func getMimeType(path string) string{
	ext := filepath.Ext(path)
	return mime.TypeByExtension(ext)
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	http.Handle("/", gziphandler.GzipHandler(new(staticHandler)))
	log.Fatal(http.ListenAndServe(*listen, nil))
}
