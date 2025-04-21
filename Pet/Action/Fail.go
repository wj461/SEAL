package Action

import (
	"time"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type Fail struct {
	moveX         int
	moveY         int
	frameInterval int // fps = 10
	prevTime      int
	y             float64
	v             float64
	a             float64 // 加速度(px/s^2)
	dt            float64
	isFirst       bool
}

func NewFail() IPetAction {
	return &Fail{0, 0, 1000 / 10, 0, 0, 2000, 1000, 0.016, true}
}

func (f *Fail) Update(windowName string) error {
	if f.isFirst {
		f.isFirst = false
		_, y := application.Get().GetWindowByName(windowName).Position()
		f.y = float64(y)
		return nil
	}

	if int(time.Now().UnixMilli())-f.prevTime < f.frameInterval {
		return nil
	}
	f.prevTime = int(time.Now().UnixMilli())

	f.v += f.a * f.dt
	f.y += f.v * f.dt

	time.Sleep(time.Duration(f.dt * float64(time.Second)))

	x, _ := application.Get().GetWindowByName(windowName).Position()
	application.Get().GetWindowByName(windowName).SetPosition(x+2, int(f.y))
	return nil
}
