package main

import (
    "log"
    "net/http"

    "github.com/codegangsta/negroni"
)

func main() {

    mux:=http.NewServeMux()

    mux.Handle("/", http.FileServer(http.Dir(".")))
    // Middleware stack
    n := negroni.New(
        negroni.NewRecovery(),
        negroni.HandlerFunc(myMiddleware),
        negroni.NewLogger(),
        //negroni.NewStatic(http.Dir(".")),
        
    )
    n.UseHandler(mux)

    n.Run(":8080")
}

func myMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    log.Println("Logging on the way there...")

    if r.URL.Query().Get("password") == "secret123" {
        next(rw, r)
    } else {
        http.Error(rw, "Not Authorized", 401)
    }

    log.Println("Logging on the way back...")
}