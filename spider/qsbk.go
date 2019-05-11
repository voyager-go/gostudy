package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"strconv"
	"strings"
)
// goquery 爬糗事百科
func main() {
	GetDocument()
}
func GetDocument()  {
	text := "qiushibaike.txt"
	f, err := os.OpenFile(text, os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	partUrl := "https://www.qiushibaike.com/text/page/"
	num := 1
	for page := 1; page <= 10; page++ {
		url     := fmt.Sprintf("%s%d", partUrl, page)
		document, err := goquery.NewDocument(url)
		if err != nil {
			log.Fatal(err)
		}
		// words content
		document.Find("#content-left .article").Each(func(i int, selection *goquery.Selection) {
			author  := strings.Replace(selection.Find(".author h2").Text(), "\n", "", -1)
			content := strings.Replace(selection.Find(".content span").Text(), "\n", "", -1)
			n, _ := f.Seek(0, os.SEEK_END)
			_, err = f.WriteAt([]byte("编号:" + strconv.Itoa(num) + "\n作者：" + author + "\n内容:" + content + "\n"), n)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(num)
			fmt.Println(author)
			fmt.Println(content)
			fmt.Println("================")
			num++
		})
	}
	defer f.Close()
}
