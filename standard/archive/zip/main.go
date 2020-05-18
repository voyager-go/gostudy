package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = f.Write([]byte(file.Body))
		if _, err = f.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}
	if err := w.Close(); err != nil {
		log.Fatalln(err)
	}
	zipFile, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err = buf.WriteTo(zipFile); err != nil {
		log.Fatalln(err)
	}
	defer zipFile.Close()

	// 读压缩文件
	r, err := zip.OpenReader("file.zip")
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Close()
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatalln(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatalln(err)
		}
		defer r.Close()
		fmt.Println()
	}
}
