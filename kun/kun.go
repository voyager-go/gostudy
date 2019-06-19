package main

import (
	"fmt"
	"github.com/BurntSushi/graphics-go/graphics"
	"image"
	"image/gif"
	"os"
	"time"
)

var pic = "kun.gif"
func main() {
	f, err := os.Open(pic)
	if err!=nil {
		panic(err)
	}
	fg, err := gif.DecodeAll(f)
	if err!=nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()
	for i := 0; i < 5; i++ {
		for _, v := range fg.Image {
			//Resize image in case it is too big to display
			//重新调整图像大小以免过大显示不全
			ni := cutImage(v, 100)
			//Convert a image into a char image
			//转化图像为字幅图像并显示
			makeImageChar(ni)
			//Delay : delay between two frames. Modify this if needed
			//延时  : 两帧之间的延时，如果有何必要就修改之
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func cutImage(img image.Image, scale int) *image.RGBA {
	bou := img.Bounds()
	dx  := bou.Dx()
	dy  := bou.Dy()
	nim := image.NewRGBA(image.Rect(0, 0, scale, scale*dy/dx))
	_ = graphics.Scale(nim, img)
	return nim
}

func makeImageChar(img image.Image)  {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	arrs := []string{"J", "N", "T", "M", "$", "O", "C", "?", "7", "0", "–", "–", "–", "–", "."}
	var singleFrame []string
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			colorRgb := img.At(j, i)
			_, g, _, _ := colorRgb.RGBA()
			avg := uint8(g >> 8)
			num := avg / 18
			singleFrame = append(singleFrame, arrs[num])
			if j == dx-1 {
				//末尾加入回车换行
				singleFrame = append(singleFrame, "\n")
			}
		}
	}
	//输出一帧
	fmt.Printf("%s", singleFrame)
}
