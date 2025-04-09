package Action

type Idle struct {
}

func NewIdle() IPetAction {
	return &Idle{}
}

func (i *Idle) Update(windowName string) error {
	return nil
}
