package main

import "os"

type PetService struct {
	imagePath map[string]string
}

func (p *PetService) NewPet() *PetService {
	p.imagePath = make(map[string]string)
	files, err := os.ReadDir("./frontend/public/PetImage")
	println("Reading files")

	if err != nil {
		println("Error reading files")
		return nil
	}
	for _, d := range files {
		println(d.Name())
		p.imagePath[d.Name()] = d.Name()
	}
	return p
}

func (p *PetService) GetPetNames() []string {
	names := make([]string, 0, len(p.imagePath))
	for k := range p.imagePath {
		names = append(names, k)
	}
	println(names)
	return names
}

func (p *PetService) GetImagePathByName(name string) string {
	return p.imagePath[name]
}
