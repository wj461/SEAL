package Pet

import "github.com/wj461/SEAL/Pet/Action"

type Pet struct {
	id            int
	windowName    string
	actionName    string
	currentAction Action.IPetAction
}

func NewPet(id int, windowName string) *Pet {
	return &Pet{id: id, windowName: windowName, actionName: "idle", currentAction: Action.NewIdle()}
}

func (p *Pet) SetAction(actionName string, action Action.IPetAction) {
	p.actionName = actionName
	p.currentAction = action
}

func (p *Pet) GetId() int {
	return p.id
}
