package images

import (
	"context"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
	"strings"
)

const (
	UniqueViolationErr = "23505"
)

type Repository struct {
	db database.Service
}

func NewRepository(db database.Service) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetByID(ctx context.Context, id string) (*model.Image, error) {
	image, err := r.db.GetImageByID(ctx, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error getting image by id: %w", err)
	}

	s, err := fromBase64(image.Content)
	if err != nil {
		return nil, fmt.Errorf("error decoding image content: %w", err)
	}

	return &model.Image{
		ID:          image.ID,
		Content:     s,
		ContentType: image.ContentType,
	}, nil
}

func (r *Repository) Store(ctx context.Context, image *model.Image) (*model.Image, error) {
	existing, err := r.GetByID(ctx, image.ID)
	if err != nil {
		return nil, fmt.Errorf("error creating image: %w", err)
	} else if existing != nil {
		return existing, nil
	}

	img, err := r.db.StoreImage(ctx, database.StoreImageParams{
		ID:          image.ID,
		Content:     toBase64(image.Content),
		ContentType: image.ContentType,
	})
	if err != nil && isErrorCode(err, UniqueViolationErr) {
		return r.GetByID(ctx, image.ID)
	} else if err != nil {
		return nil, fmt.Errorf("error creating image: %w", err)
	}

	return &model.Image{
		ID:          img.ID,
		Content:     img.Content,
		ContentType: img.ContentType,
	}, nil
}

func toBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func fromBase64(input string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(input)
	return string(data), err
}

func isErrorCode(err error, errcode string) bool {
	return strings.Contains(err.Error(), errcode)
}
