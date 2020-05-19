package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// tar包实现了tar格式压缩文件的存取。本包目标是覆盖大多数tar的变种，包括GNU和BSD生成的tar文件
func main() {
	// 压缩
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence."},
	}
	for _, file := range files {
		hdr := &tar.Header{
			Name:     file.Name,
			Linkname: "",
			Size:     int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	// 读压缩文件
	r := bytes.NewReader(buf.Bytes())
	tr := tar.NewReader(r)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Contents is %s\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
	}
}
