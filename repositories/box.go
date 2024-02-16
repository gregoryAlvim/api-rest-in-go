package repositories

import (
	"fmt"

	"github.com/gregoryAlvim/api-rest-in-go/data"
	"github.com/gregoryAlvim/api-rest-in-go/models"
)

type BoxRepo struct{}

func (b *BoxRepo) Create(newData *models.Box) {
	data.Boxes[newData.Id] = *newData
}

func (b *BoxRepo) Find(id string) (*models.Box, error) {
	if len(data.Boxes) == 0 {
		return nil, fmt.Errorf("não há caixas cadastradas")
	}

	box, boxFound := data.Boxes[id]

	if boxFound {
		return &box, nil
	} else {
		return nil, fmt.Errorf("nenhuma caixa foi encontrada")
	}
}

func (b *BoxRepo) FindAll() ([]models.Box, error) {

	if len(data.Boxes) == 0 {
		return nil, fmt.Errorf("não há caixas cadastradas")
	}
	var boxes []models.Box

	for _, box := range data.Boxes {
		boxes = append(boxes, box)
	}

	return boxes, nil
}

func (b *BoxRepo) Update(id string, newData *models.Box) error {

	_, idExists := data.Boxes[id]

	newData.Id = id

	if idExists {
		data.Boxes[id] = *newData
		return nil
	} else {
		return fmt.Errorf("não foi possível atualizar essa caixa")
	}
}

func (b *BoxRepo) Remove(id string) error {
	print("Chegou")
	_, idExists := data.Boxes[id]

	if idExists {
		delete(data.Boxes, id)
		return nil
	} else {
		return fmt.Errorf("não foi possível excluir essa caixa")
	}
}
