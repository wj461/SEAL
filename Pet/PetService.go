package Pet

import (
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wj461/SEAL/Pet/Action"
)

type Bound struct {
	X      int
	Y      int
	Width  int
	Height int
}

type PetService struct {
	currentPetAction IPetAction
	actionMap        map[string]IPetAction
	bounds           Bound
	state            string
}

func NewPetService() *PetService {
	p := &PetService{}
	p.actionMap = make(map[string]IPetAction)
	p.state = "idle"
	// p.bounds = p.GetScreenBounds()

	p.actionMap["walk"] = Action.NewWalk()
	p.actionMap["idle"] = Action.NewIdle()
	return p
}

// func (p *PetService) Init() {
// 	p.actionMap = make(map[string]IPetAction)
// 	p.actionMap["walk"] = Action.NewWalk()
// 	p.actionMap["idle"] = Action.NewIdle()
// }

func (p *PetService) SetAction(action string) {
	p.state = action
	p.currentPetAction = p.actionMap[action]
}

func (p *PetService) GetState() string {
	return p.state
}

func (p *PetService) Update() error {
	if p.currentPetAction != nil {
		p.currentPetAction.Update()
	}
	return nil
}

func (p *PetService) GetScreenBounds() Bound {
	screenBounds, _ := application.Get().CurrentWindow().GetScreen()
	s := screenBounds.WorkArea

	return Bound{s.X, s.Y, s.Width, s.Height}
}
