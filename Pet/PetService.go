package Pet

import (
	"log"
	"log/slog"

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
	actionFactory map[string]func() Action.IPetAction
	bounds        Bound
	petList       []Pet
}

func NewPetService() *PetService {
	p := &PetService{}
	p.actionFactory = map[string]func() Action.IPetAction{}
	p.petList = make([]Pet, 0)

	p.actionFactory["walk"] = Action.NewWalk
	p.actionFactory["idle"] = Action.NewIdle
	return p
}

func (p *PetService) NewPetForFrontend() int {
	p.bounds = p.GetScreenBounds()
	windowName := application.Get().CurrentWindow().Name()
	id := p.GeneratePetId()
	p.AddPet(id, windowName)
	return id
}

func (p *PetService) GetPetById(id int) *Pet {
	for _, pet := range p.petList {
		application.DefaultLogger(slog.LevelDebug).Debug("🐟pet", &pet, "SetAction id:", pet.id, "action:", pet.status)
	}

	for i, pet := range p.petList {
		if pet.GetId() == id {
			return &p.petList[i]
		}
	}
	return nil
}

func (p *PetService) GeneratePetId() int {
	id := 0
	for _, pet := range p.petList {
		if pet.GetId() > id {
			id = pet.GetId()
		}
	}
	return id + 1
}

func (p *PetService) AddPet(id int, windowName string) {
	pet := NewPet(id, windowName)
	p.petList = append(p.petList, *pet)
}

func (p *PetService) SetAction(id int, action string) {
	pet := p.GetPetById(id)
	if pet == nil {
		return
	}
	pet.SetAction(action, p.actionFactory[action]())
}

func (p *PetService) GetState(id int) string {
	pet := p.GetPetById(id)
	if pet == nil {
		return "pet not found"
	}
	return pet.status
}

func (p *PetService) Update() error {
	for _, pet := range p.petList {
		if pet.currentAction != nil {
			log.Println("pet id:", pet.GetId(), "action:", pet.status)
			err := pet.currentAction.Update(pet.windowName)
			if err != nil {
				log.Println("Error updating action:", err)
				return err
			}
		}
	}

	return nil
}

func (p *PetService) GetScreenBounds() Bound {
	screenBounds, _ := application.Get().CurrentWindow().GetScreen()
	if screenBounds == nil {
		log.Println("GetScreenBounds: screenBounds is nil")
		return Bound{0, 0, 0, 0}
	}
	s := screenBounds.WorkArea
	return Bound{s.X, s.Y, s.Width, s.Height}
}
