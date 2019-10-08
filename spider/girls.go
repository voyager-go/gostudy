package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strconv"
)
// 功能描述：使用Golang编写爬虫，抓取“妹子图”网站部分图片
// 作者声明：本篇代码只用作学习与交流，不可作商业用途
func main() {
	GrabIndexPage()
}
// 抓取首页信息
func GrabIndexPage()  {
	url := "http://www.mmmjpg.com"
	document, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}
	document.Find(".pic li").Each(func(i int, selection *goquery.Selection) {
		// 缩略图
		//smallImg, _ := selection.Find("img").Attr("src")
		// 详情页的链接后半部分
		pageHref, _ :=  selection.Find("a").Attr("href")
		detailPageLink := url + pageHref
		GrabDetailPage(detailPageLink)
	})
}
// 获取详情页
func GrabDetailPage(link string)  {
	document,err := goquery.NewDocument(link)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}
	imgSrc,_ := document.Find(".content img").Attr("src")
	series,_ := document.Find(".content img").Attr("alt")
	// 结果如：http://img.hi0590.com/1453/1.jpg
	// 根据页数查共有多少张图片
	pageNumStr := document.Find(".page a").Eq(-2).Text()
	pageNum, _ := strconv.Atoi(pageNumStr)
	// 截取http://img.hi0590.com/1453/
	for i := 1; i <= pageNum; i++ {
		filename := strconv.Itoa(i) + ".jpg"
		tmpSrc := imgSrc[ : len(imgSrc) - 5] + filename
		DownloadImage(tmpSrc, series, filename)
	}
}
// 下载图片
func DownloadImage(src, series, filename string)  {
	dir := "spider/images/www.mmmjpg.com/" + series + "/"
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}
	req, err := http.NewRequest("Get", src, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}
	resp, _ := http.DefaultClient.Do(req)
	file, err := os.OpenFile(dir + filename, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}
	fmt.Println("正在抓取图片链接为:" + src)
	if _, err  = io.Copy(file, resp.Body);err != nil {
		panic("文件保存失败: " + fmt.Sprintf("%v", err))
	}

	fmt.Println("抓取" + src + "结束")
}