package Pet

type PetObject struct {
	id         int
	windowName string
	actionName string
	// currentAction Action.IPetAction
}

func NewPet(id int, windowName string) *PetObject {
	// return &PetObject{id: id, windowName: windowName, actionName: "idle", currentAction: Action.NewIdle()}
	return &PetObject{id: id, windowName: windowName, actionName: "failing"}
}

//	func (p *PetObject) SetAction(actionName string, action Action.IPetAction) {
//		p.actionName = actionName
//		p.currentAction = action
//	}
func (p *PetObject) SetAction(actionName string) {
	p.actionName = actionName
}

func (p *PetObject) GetId() int {
	return p.id
}

func (p *PetObject) GetWindowName() string {
	return p.windowName
}

func (p *PetObject) GetActionName() string {
	return p.actionName
}

// func (p *PetObject) GetCurrentAction() Action.IPetAction {
// 	return p.currentAction
// }
