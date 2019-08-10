// var   数据类型 = 值  # 定义变量
// const 数据类型 = 值  # 定义常量
// 复合声明 , 实例化复合类型数据结构
// struct 类型, 一组值或者字段的集合, 访问内部变量使用 点操作符
package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
)

// 定义包变量 - 调色板
// 复合声明 []color.Color{color.White, color.Black}
var palette = []color.Color{color.White, color.Black}

// 定义包常量。。
const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	out := os.Stdout

	// 定义局部常量
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	//
	freq := rand.Float64() * 3.0
	// 复合声明 gif.GIF{LoopCount: nframes}
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		// 生成动画帧
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// 设置两个偏振值
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	//
	gif.EncodeAll(out, &anim)
}
