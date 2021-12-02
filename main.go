package main

import (
	"basic"
	"fmt"
	"net/http"
)

func main() {
	app := basic.New()
	app.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
	})
	app.Run(":9000")
}
