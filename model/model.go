package model

import (
	"math"
	"math/rand"
	"time"
)

const (
	ROAD_L = 2000 // roadの長さ
	VIEW_L = 300  // 視界の長さ

	PART_L   = 100  // 各partの幅（z方向）
	CAMERA_D = 0.8  // roadとの距離がこの値に等しいときscaleが1になる
	ROAD_W   = 1000 // roadの幅（x方向）
	JIKI_Y   = 1000 // 自機の高さ
)

type Road struct {
	X  int32
	Y  int32
	Z  int32
	C  int32
	Sx int32
}

type Game struct {
	Rng    *rand.Rand
	IsOver bool
	Frame  int32
	Road   []Road
	Jiki_x int32
	Jiki_y int32
	Jiki_z int32
}

func NewGame() *Game {
	timestamp := time.Now().Unix()
	rng := rand.New(rand.NewSource(timestamp))

	g := &Game{}
	g.Rng = rng
	g.IsOver = false
	g.Frame = 0

	for i := 0; i < ROAD_L; i++ {
		part := Road{}
		part.X = 0
		part.Y = 0
		part.Z = int32(i) * PART_L
		part.C = 0
		part.Sx = 0

		if i > 100 && i <= 200 {
			part.C = 1
		}
		if i > 300 && i <= 400 {
			part.C = -4
		}

		if i > 400 {
			part.Y = int32(math.Sin((float64(i)-400)/30) * 1000)
		}
		g.Road = append(g.Road, part)
	}

	g.Jiki_x = 0
	g.Jiki_y = JIKI_Y
	g.Jiki_z = 0

	return g
}

func (g *Game) Update(command string) {
	if g.IsOver {
		return
	}

	switch command {
	case "left":
		g.Jiki_x -= 100
	case "right":
		g.Jiki_x += 100
	case "down":
		g.Jiki_z -= 100
	case "up":
		g.Jiki_z += 100
	}

	g.Frame += 1
}
