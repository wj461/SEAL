package Action

import (
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Walk struct {
	moveX         int
	moveY         int
	frameInterval int // fps = 10
	prevTime      int
}

func NewWalk() IPetAction {
	return &Walk{10, 0, 1000 / 10, 0}
}

func (w *Walk) Update(windowName string) error {
	if int(time.Now().UnixMilli())-w.prevTime < w.frameInterval {
		return nil
	}

	w.prevTime = int(time.Now().UnixMilli())
	x, y := application.Get().GetWindowByName(windowName).Position()
	application.Get().GetWindowByName(windowName).SetPosition(x-w.moveX, y+w.moveY)
	return nil
}
