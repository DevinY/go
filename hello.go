package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "嗨, 執行Go囉: %s\n", r.URL.Path)
	})

        http.ListenAndServe(":80", nil)
}
