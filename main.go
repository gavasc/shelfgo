package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/webdav"
)

func main() {
	user := os.Getenv("SHELFGO_USER")
	pass := os.Getenv("SHELFGO_PASS")
	serveDir := os.Getenv("SHELFGO_DIR")
	port := os.Getenv("SHELFGO_PORT")

	if err := ensureDir(serveDir); err != nil {
		log.Fatal("Failed to create directory: ", err)
	}

	handler := &webdav.Handler{
		Prefix:     "/",
		FileSystem: webdav.Dir(serveDir),
		LockSystem: webdav.NewMemLS(),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Request received from ", r.Host)
		if !isAuthorized(r, user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Print("Request from ", r.Host, " unauthorized")
			return
		}

		handler.ServeHTTP(w, r)
	})

	log.Printf("Starting ShelfGo server on port %s, serving files from %s", port, serveDir)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}

func ensureDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0755)
	}

	return nil
}
