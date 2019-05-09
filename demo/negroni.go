package main

import (
	"net/http"
	"io"
	"github.com/urfave/negroni"
)

func main()  {
	n := negroni.Classic()
	n.UseHandler(NewHandler())
	n.Run(":9999")
}
func NewHandler() http.Handler {
	return http.HandlerFunc(MyNewHandler)
}
func MyNewHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Hello World ...")
}