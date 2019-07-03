package main

import (
	"bytes"
	"archive/zip"
	"io"
	"compress/flate"
	"log"
	"io/ioutil"
	"os"
	"fmt"
)
const FilePath  = "/Users/artist/go/src/gostudy/standard_library/archive/zip/tartest/tarexample.zip"

// 实现了zip档案文件的读写
func main() {
	// 写zip文件数据流
	buf := write()
	// 自动生成并写入文件
	if err := ioutil.WriteFile(FilePath, buf.Bytes(), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// 读取zip文件数据流
	read(FilePath)
}
func write() bytes.Buffer  {
	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	// 设置压缩级别，不指定则默认
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return  flate.NewWriter(out, flate.BestCompression)
	})
	var files = []struct{
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	return buf
}

func read(path string)  {
	r, err := zip.OpenReader(path)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	for _, f := range r.File{
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		content, err := ioutil.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", content)
		// 关闭文件
		rc.Close()
	}
}