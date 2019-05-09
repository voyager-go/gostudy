package main

import (
	"net/http"
	"io"
	"github.com/urfave/negroni"
)
// Negroni配合路由
func main() {
	neg := negroni.Classic()
	mux := http.NewServeMux()
	mux.HandleFunc("/", TestHandler)
	mux.HandleFunc("/poetry", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "text/html;charset=utf-8")
		io.WriteString(writer, "<a href='http://www.blockchina.art'>中华有诗词</a>")
	})
	neg.UseHandler(mux)
	neg.Run(":9999")
}

func TestDoHandler() http.Handler {
	return http.HandlerFunc(TestHandler)
}

func TestHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/plain")
	io.WriteString(w, "Hello World ..... Hello People")
}