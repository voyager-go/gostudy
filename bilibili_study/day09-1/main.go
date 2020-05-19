package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	toWrite()
	toRead()
}

func toWrite() {
	file, err := os.OpenFile("file.txt", os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal("%v", err)
	}
	defer file.Close()
	var str = "nice to meet you too!"
	count, err := file.WriteString(str)
	if err != nil {
		log.Fatal("%v", err)
	}
	fmt.Printf("%d", count)
}
func toRead() {
	file, err := os.OpenFile("file.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal("%v", err)
	}
	defer file.Close()
	var data = make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal("%v", err)
	}
	fmt.Printf("%d %q", count, data[:count])
}
