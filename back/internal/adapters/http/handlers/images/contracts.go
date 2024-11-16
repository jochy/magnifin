package images

import (
	"context"
	"magnifin/internal/app/model"
)

type ImageRepository interface {
	GetByID(ctx context.Context, id string) (*model.Image, error)
}

type Handlers struct {
	ImageRepository ImageRepository
}

func NewHandlers(imageRepository ImageRepository) *Handlers {
	return &Handlers{
		ImageRepository: imageRepository,
	}
}
