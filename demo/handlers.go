package main

import (
	"net/http"
	"github.com/gorilla/handlers"
	"os"
	"io"
)

func main() {
	http.Handle("/", useLoggingHandler(handler()))
	http.ListenAndServe(":9999", nil)
}

func handler() http.Handler {
	return  http.HandlerFunc(myHandler)
}

func myHandler(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello World")
}

func useLoggingHandler(next http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, negitxt)
}