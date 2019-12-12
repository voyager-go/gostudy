package main

import (
	"net/rpc"
	"net/http"
	"io"
	"net/rpc/jsonrpc"
)

func main()  {
	rpc.RegisterName("HelloService", new(HelloService))
	http.HandleFunc("jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}