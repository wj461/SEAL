package Action

import (
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Move struct {
	moveX         int
	moveY         int
	frameInterval int // fps = 10
	prevTime      int
}

func NewLeftMove() IPetAction {
	return &Move{-10, 0, 1000 / 8, 0}
}

func NewRightMove() IPetAction {
	return &Move{10, 0, 1000 / 8, 0}
}

func NewUpMove() IPetAction {
	return &Move{0, -10, 1000 / 8, 0}
}

func NewDownWove() IPetAction {
	return &Move{0, 10, 1000 / 8, 0}
}

func NewFailMove() IPetAction {
	return &Move{0, 20, 1000 / 10, 0}
}

// func (w *Walk) SetMoveX(x int) {
// 	w.moveX = x
// }

// func (w *Walk) SetMoveY(Y int) {
// 	w.moveY = Y
// }

// func (w *Walk) SetMoveXY(x, y int) {
// 	w.moveX = x
// 	w.moveY = y
// }

func (w *Move) Update(windowName string) error {
	if int(time.Now().UnixMilli())-w.prevTime < w.frameInterval {
		return nil
	}
	// get is dragging
	if application.Get().GetWindowByName(windowName).IsFocused() {
		return nil
	}

	w.prevTime = int(time.Now().UnixMilli())
	x, y := application.Get().GetWindowByName(windowName).Position()
	application.Get().GetWindowByName(windowName).SetPosition(x+w.moveX, y+w.moveY)
	return nil
}
