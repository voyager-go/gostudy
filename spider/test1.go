package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func main() {
	GetPage()
}
func GetPage()  {
	document, err := goquery.NewDocument("https://www.saikr.com/a/2669")
	if err != nil {
		log.Fatal(err)
	}
	//document.Text()
	title := document.Find("title").Text()
	fmt.Println("The website title is " + title)
	keywords, ok := document.Find("meta[name=keywords]").Attr("content")
	if ok {
		fmt.Println("The website keywords content is " + keywords)
	}

	description, ok := document.Find("meta[name=description]").Attr("content")
	if ok {
		fmt.Println("The website keywords description is " + description)
	}
	document.Find(".sk-circle4-1-para-box").Find(".para").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
