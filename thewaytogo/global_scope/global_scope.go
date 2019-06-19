package main

var b = "G"

func main() {
	x()
	y()
	x()
}

func x() {
	print(b)
}

func y() {
	b = "O"
	print(b)
}