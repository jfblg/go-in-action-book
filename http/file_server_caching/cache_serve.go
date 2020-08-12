package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

type cacheFile struct {
	content io.ReadSeeker
	modTime time.Time
}

var cache map[string]*cacheFile
var mutex = new(sync.RWMutex)

func main() {
	cache = make(map[string]*cacheFile)
	http.HandleFunc("/", serveFiles)
	http.ListenAndServe(":8080", nil)
}

func serveFiles(res http.ResponseWriter, req *http.Request) {
	mutex.RLock()
	v, found := cache[req.URL.Path]
	mutex.RUnlock()

	if found {
		log.Printf("From cache: %q\n", req.URL.Path)
	}

	if !found {
		mutex.Lock()
		defer mutex.Unlock()

		filename := "./files" + req.URL.Path
		f, err := os.Open(filename)
		defer f.Close()
		if err != nil {
			log.Println(err)
			http.NotFound(res, req)
			return
		}
		var b bytes.Buffer
		_, err = io.Copy(&b, f)
		if err != nil {
			log.Println(err)
			http.NotFound(res, req)
			return
		}
		r := bytes.NewReader(b.Bytes())

		info, _ := f.Stat()
		v := &cacheFile{
			content: r,
			modTime: info.ModTime(),
		}
		cache[req.URL.Path] = v
	}

	http.ServeContent(res, req, req.URL.Path, v.modTime, v.content)

}
