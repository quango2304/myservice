package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := getenv("PORT", "8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host, _ := os.Hostname()
		fmt.Fprintf(w, "hello from myservice\nhost=%s\npath=%s\n", host, r.URL.Path)
	})
	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
