package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"log"
	"net/http"
)
var str string
func main() {
	if err := termbox.Init(); err != nil{
		panic(err)
	}
	fmt.Println("Process begin...")
	KeepUp()
}
func KeepUp()  {
	fmt.Println("keep up")
Loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			str = string(rune(ev.Ch))
			fmt.Printf("Key is: %s", str)
			switch ev.Key {
			case termbox.KeyEsc:
				termbox.Close()
				break Loop
			}
		}
	}
}
func toWrite(w http.ResponseWriter, r *http.Request)  {
	html := "<body>"+str+"<body><script>setInterval(function(){window.location.reload()}, 1000);</script>"
	w.Write([]byte(html))
}
func startHttp()  {
	http.HandleFunc("/", toWrite)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
