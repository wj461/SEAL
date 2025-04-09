package Pet

import "github.com/wj461/SEAL/Pet/Action"

type Pet struct {
	id            int
	windowName    string
	status        string
	currentAction IPetAction
}

func NewPet(id int, windowName string) *Pet {
	return &Pet{id: id, windowName: windowName, status: "idle", currentAction: Action.NewIdle()}
}

func (p *Pet) SetAction(status string, action IPetAction) {
	p.status = status
	p.currentAction = action
	println("SetAction", p.status)
}

func (p *Pet) GetId() int {
	return p.id
}
