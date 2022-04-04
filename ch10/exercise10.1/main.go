package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

// 练习9.6: 测试一下计算密集型的并发程序(练习8.5那样的)会被GOMAXPROCS怎样影响到。
// 在你的电脑上 最佳的值是多少？你的电脑CPU有多少个核心？

var (
	format string
)

func main() {
	flag.StringVar(&format, "format", "jpeg", "")
	flag.Parse()

	f, err := os.Open("test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, kind, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stderr, "Input format = ", kind)
	convert(format, img)
}

func convert(format string, img image.Image) {
	var (
		f   *os.File
		err error
	)
	fmt.Println(format)
	switch format {
	case "jpeg":
		f, err = os.Create("test.jpeg")
		if err != nil {
			return
		}
		defer f.Close()

		err := jpeg.Encode(f, img, &jpeg.Options{Quality: 95})
		if err != nil {
			log.Fatal(err)
		}
	case "png":
		f, err = os.Create("test.png")
		if err != nil {
			return
		}
		defer f.Close()

		err := png.Encode(f, img)
		if err != nil {
			log.Fatal(err)
		}
	case "gif":
		f, err = os.Create("test.gif")
		if err != nil {
			return
		}
		defer f.Close()

		err := gif.Encode(f, img, &gif.Options{})
		if err != nil {
			log.Fatal(err)
		}
	}
}
