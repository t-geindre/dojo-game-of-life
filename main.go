package main

import (
	d00 "dojo-game-of-life/00-start"
	d01 "dojo-game-of-life/01-chaos"
	d02 "dojo-game-of-life/02-order"
	d03 "dojo-game-of-life/03-optimization"
	d04 "dojo-game-of-life/04-zoom"
	d05 "dojo-game-of-life/05-slow"
	d06 "dojo-game-of-life/06-test"
	d07 "dojo-game-of-life/07-loop"
	d08 "dojo-game-of-life/08-translate"
	d09 "dojo-game-of-life/09-draw"
	"dojo-game-of-life/game"
	"github.com/hajimehoshi/ebiten/v2"
	"os"
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
	mx, my := ebiten.Monitor().Size()
	ratio := float64(my) / float64(mx)

	dojos := map[string]*Launcher{
		"00": {l: d00.NewGame},
		"01": {l: d01.NewGame},
		"02": {l: d02.NewGame},
		"03": {l: d03.NewGame, w: 1800, h: 900},
		"04": {l: d04.NewGame, w: 1440, h: int(1440 * ratio)},
		"05": {l: d05.NewGame, w: 1440, h: int(1440 * ratio)},
		"06": {l: d06.NewGame, w: 1440, h: int(1440 * ratio)},
		"07": {l: d07.NewGame, w: 1440, h: int(1440 * ratio)},
		"08": {l: d08.NewGame, w: 1440, h: int(1440 * ratio)},
		"09": {l: d09.NewGame, w: 1440, h: int(1440 * ratio)},
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
