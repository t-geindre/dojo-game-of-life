package controls

import "github.com/hajimehoshi/ebiten/v2"

var cursorStack = make([]ebiten.CursorShapeType, 0)

func pushCursor(c ebiten.CursorShapeType) {
	cursorStack = append(cursorStack, c)
	refreshCursor()
}

func popCursor(cr ebiten.CursorShapeType) {
	for i, c := range cursorStack {
		if c == cr {
			cursorStack = append(cursorStack[:i], cursorStack[i+1:]...)
			break
		}
	}
	refreshCursor()
}

func refreshCursor() {
	if len(cursorStack) > 0 {
		ebiten.SetCursorShape(cursorStack[len(cursorStack)-1])
	} else {
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	}
}
