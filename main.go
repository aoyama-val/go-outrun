package main

import (
	"time"

	m "github.com/aoyama-val/go-outrun/model"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEEN_W    = 600
	SCREEN_H     = 400
	FPS          = 60
	WINDOW_TITLE = "go-outrun"
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(WINDOW_TITLE, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, SCREEEN_W, SCREEN_H, sdl.WINDOW_SHOWN)
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
				if t.State == sdl.PRESSED && t.Repeat == 0 {
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
		time.Sleep((1000 / FPS) * time.Millisecond)
	}
}

func render(renderer *sdl.Renderer, window *sdl.Window, game *m.Game) {
	renderer.SetDrawColor(0x66, 0xaa, 0xff, 0)
	renderer.Clear()

	if game.IsOver {
		renderer.SetDrawColor(0, 0, 0, 128)
		renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: SCREEEN_W, H: SCREEN_H})
	}

	renderer.Present()
}
