package main

import (
	d00 "dojo-game-of-life/00-start"
	d01 "dojo-game-of-life/01-chaos"
	d02 "dojo-game-of-life/02-order"
	d03 "dojo-game-of-life/03-optimization"
	d04 "dojo-game-of-life/04-end"
	"dojo-game-of-life/game"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	def    = "00"
	width  = 500
	height = 500
)

type Launcher struct {
	l    func(int, int) game.Game
	w, h int
}

func main() {
	dojos := map[string]*Launcher{
		"00": {l: d00.NewGame},
		"01": {l: d01.NewGame},
		"02": {l: d02.NewGame},
		"03": {l: d03.NewGame, w: 1800, h: 900},
		"04": {l: d04.NewGame, w: 1800, h: 900},
	}

	dojo := dojos[def]
	if len(os.Args) == 2 {
		if _, ok := dojos[os.Args[1]]; ok {
			dojo = dojos[os.Args[1]]
		}
	}

	if dojo.w == 0 || dojo.h == 0 {
		dojo.w = width
		dojo.h = height
	}

	g := game.NewDefaultGame(dojo.l(dojo.w, dojo.h), dojo.w, dojo.h)
	w, h := g.Layout(dojo.w, dojo.h)

	ebiten.SetWindowTitle("Game of life")
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
