package main

import (
	"time"

	m "github.com/aoyama-val/go-outrun/model"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	WINDOW_TITLE = "go-outrun"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(WINDOW_TITLE, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, m.SCREEN_W, m.SCREEN_H, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	err = renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		panic(err)
	}

	running := true
	game := m.NewGame()

	for running {
		command := ""
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.State == sdl.PRESSED {
					keyCode := t.Keysym.Sym
					switch keyCode {
					case sdl.K_ESCAPE:
						running = false
					case sdl.K_LEFT:
						command = "left"
					case sdl.K_RIGHT:
						command = "right"
					case sdl.K_DOWN:
						command = "down"
					case sdl.K_UP:
						command = "up"
					}
				}
			}
		}
		game.Update(command)
		render(renderer, window, game)
		time.Sleep((1000 / m.FPS) * time.Millisecond)
	}
}

func render(renderer *sdl.Renderer, window *sdl.Window, g *m.Game) {
	renderer.SetDrawColor(0x66, 0xaa, 0xff, 0)
	renderer.Clear()

	start := g.Jiki_z / m.PART_L

	mx := int32(0)
	my := int32(0)
	mw := int32(0)

	// 奥から手前に描画
	for i := start + m.VIEW_L - 1; i >= start; i-- {
		if i < 0 || i >= m.ROAD_L {
			continue
		}
		r := g.Road[i]
		dist := r.Z - g.Jiki_z
		if dist == 0 {
			continue
		}
		scale := float64(m.CAMERA_D) / float64(dist)

		px := int32((1 + float64(r.X-g.Jiki_x+r.Sx)*scale) * float64(m.SCREEN_W) / 2)
		py := int32((1 - float64(r.Y-g.Jiki_y)*scale) * float64(m.SCREEN_H) / 2)
		pw := int32(float64(m.ROAD_W) * scale * float64(m.SCREEN_W))

		// fmt.Printf("px = %d, py = %d\n", px, py)

		if mx != 0 {
			var col sdl.Color
			if i%3 != 0 {
				col = sdl.Color{R: 0xaa, G: 0xaa, B: 0xaa, A: 255}
			} else {
				col = sdl.Color{R: 0xbb, G: 0xbb, B: 0xbb, A: 255}
			}
			var edg sdl.Color
			if i%3 != 0 {
				edg = sdl.Color{R: 0xbb, G: 0xbb, B: 0xbb, A: 255}
			} else {
				edg = sdl.Color{R: 0xff, G: 0xff, B: 0xff, A: 255}
			}
			var grn sdl.Color
			if i%3 != 0 {
				grn = sdl.Color{R: 0x66, G: 0xff, B: 0x66, A: 255}
			} else {
				grn = sdl.Color{R: 0x88, G: 0xff, B: 0x88, A: 255}
			}
			drawRoad(renderer, grn, m.SCREEN_W/2, my, m.SCREEN_W, m.SCREEN_W/2, py, m.SCREEN_W)
			drawRoad(renderer, edg, mx, my, int32(float64(mw)*1.1), px, py, int32(float64(pw)*1.1))
			drawRoad(renderer, col, mx, my, mw, px, py, pw)
		}
		mx = px
		my = py
		mw = pw
	}
	if g.IsOver {
		renderer.SetDrawColor(0, 0, 0, 128)
		renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: m.SCREEN_W, H: m.SCREEN_H})
	}

	// panic("a")

	renderer.Present()
}

func drawRoad(renderer *sdl.Renderer, col sdl.Color, mx, my, mw, px, py, pw int32) {
	y1 := int16(my)
	y2 := int16(py)
	x1 := int16(mx - mw/2)
	x2 := int16(px - pw/2)
	// fmt.Printf("x1 = %d, y1 = %d, x2 = %d, y2 = %d\n", x1, y1, x2, y2)

	x1w := x1 + int16(mw)
	x2w := x2 + int16(pw)

	gfx.FilledPolygonColor(renderer,
		[]int16{x1, x1w, x2w, x2},
		[]int16{y1, y1, y2, y2},
		col)
}
