package Action

type Idle struct {
}

func NewIdle() *Idle {
	return &Idle{}
}

func (i *Idle) Update() error {
	return nil
}
