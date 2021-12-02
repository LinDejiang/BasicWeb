package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	app := gee.New()
	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
	})
	app.Run(":9000")
}
