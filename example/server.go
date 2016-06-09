package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/der-antikeks/go-webdav"
)

var (
	path = os.Getenv("OPENSHIFT_DATA_DIR")
)

func main() {
	os.Mkdir(path, os.ModeDir)
	
	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))

	// http.StripPrefix is not working, webdav.Server has no knowledge
	// of stripped component, but needs for COPY/MOVE methods.
	// Destination path is supplied as header and needs to be stripped.
	http.Handle("/webdav/", &webdav.Server{
		Fs:         webdav.Dir(path),
		TrimPrefix: "/webdav/",
		Listings:   true,
	})

	http.HandleFunc("/", index)

	log.Println("Listening on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(bind, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q\n", r.URL.Path)
}
