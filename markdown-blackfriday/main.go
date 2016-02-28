package main

import (
	"net/http"
	"os"
	//"io"
	"fmt"

	"github.com/russross/blackfriday"
)

func main() {
	http.HandleFunc("/markdown", generateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}

func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	for k, v := range r.Header {
		for _, l := range v {
			fmt.Fprintf(rw, "%s: %s\n", k, l)
			//io.WriteString(rw, k+" : "+l+"\n")
			//rw.Write([]byte(k+" : "+l+"\n"))
		}

	}

	rw.Write(markdown)
}
