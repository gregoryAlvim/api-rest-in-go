package services

import (
	"github.com/gregoryAlvim/api-rest-in-go/models"
	"github.com/gregoryAlvim/api-rest-in-go/repositories"
)

type BoxService struct{}

func (b *BoxService) Create(newData *models.Box) {
	boxRepo := repositories.BoxRepo{}

	boxRepo.Create(newData)
}

func (b *BoxService) Find(id string) (*models.Box, error) {
	boxRepo := repositories.BoxRepo{}

	box, err := boxRepo.Find(id)

	if err != nil {
		return nil, err
	}

	return box, nil
}

func (b *BoxService) FindAll() ([]models.Box, error) {
	boxRepo := repositories.BoxRepo{}

	boxes, err := boxRepo.FindAll()

	if err != nil {
		return nil, err
	}

	return boxes, nil
}

func (b *BoxService) Update(id string, newData *models.Box) error {
	boxRepo := repositories.BoxRepo{}

	err := boxRepo.Update(id, newData)

	if err != nil {
		return err
	}

	return nil
}

func (b *BoxService) Remove(id string) error {
	boxRepo := repositories.BoxRepo{}

	err := boxRepo.Remove(id)

	if err != nil {
		return err
	}

	return nil
}
