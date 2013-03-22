package main

import (
	"github.com/skelterjohn/go.wde"
	_ "github.com/skelterjohn/go.wde/init"
	"image/color"
	"log"
)

func main() {
	go block()
	wde.Run()
}

func draw(dw wde.Window) {
	width, height := dw.Size()
	scr := dw.Screen()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			s := x + y
			if (s/10)%2 == 0 {
				scr.Set(x, y, color.Black)
			} else {
				scr.Set(x, y, color.RGBA{100, 100, 100, 255})
			}
		}
	}
	dw.FlushImage()
}

func block() {
	dw, err := wde.NewWindow(100, 100)
	dw.Show()
	if err != nil {
		log.Fatal(err)
		return
	}
	draw(dw)

	for e := range dw.EventChan() {
		if _, ok := e.(wde.ResizeEvent); ok {
			draw(dw)
		}
	}
}
