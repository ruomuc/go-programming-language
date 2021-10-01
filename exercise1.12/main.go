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
	"strconv"
	"time"
)

var nColors = 100

// 1.12 修改Lissajour服务，从URL读取变量，比如你可以访问
// http://localhost:8000/?cycles=20 这个URL，这样
// 访问可以将程序里的cycles默认的5修改为20。
// 字符串转换为数字可以调用strconv.Atoi函数。
// 你可以在godoc里查看strconv.Atoi的详细说明。
func main() {
	rand.Seed(time.Now().Unix())

	var palette []color.Color
	for i := 0; i < nColors; i++ {
		r := uint8(rand.Uint32() % 256)
		g := uint8(rand.Uint32() % 256)
		b := uint8(rand.Uint32() % 256)
		palette = append(palette, color.RGBA{r, g, b, 0xff})
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		// 获取query参数 cycles
		cycles, _ := strconv.Atoi(r.URL.Query().Get("cycles"))
		lissajous(w, palette, cycles)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// cycles is 完整的 x 振荡器变化的个数
func lissajous(out io.Writer, palette []color.Color, cycles int) {
	const (
		res     = 0.001 // 角度的分辨率
		size    = 100   // 画像画布包含 [-size..+size]
		nframes = 64    // 动画中的帧数
		delay   = 20    // 以 10ms 为单位的帧间延迟
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		nColor := uint8(i % len(palette))
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), nColor)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
