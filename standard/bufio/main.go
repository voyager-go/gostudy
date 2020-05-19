package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintf(w, "Hello, ")
	fmt.Fprintf(w, "World!")
	w.Flush()
}
