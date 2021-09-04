package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var palette = []color.Color{color.RGBA{0, 0, 0, 0xff}, color.RGBA{0, 0xff, 0, 0xff}}

// 1.5 改变利萨茹程序的画板颜色为绿底黑字来增加真实性。
// 使用 color.RGBA{0xRR,0xGG,0xBB,0xFF} 创建一种
// Web 颜色 #RRGGBB，每一对十六进制数字表示组成一个像素
// 红、绿、蓝分量的亮度
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 完整的 x 振荡器变化的个数
		res     = 0.001 // 角度的分辨率
		size    = 100   // 画像画布包含 [-size..+size]
		nframes = 64    // 动画中的帧数
		delay   = 8     // 以 10ms 为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
