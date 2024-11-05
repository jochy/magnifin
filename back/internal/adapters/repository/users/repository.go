package users

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"magnifin/internal/adapters/repository"
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
)

type Repository struct {
	db        database.Service
	CypherKey string
}

func NewRepository(db database.Service, cypherKey string) *Repository {
	return &Repository{
		db:        db,
		CypherKey: cypherKey,
	}
}

func (r *Repository) GetUserByUsernameAndPassword(ctx context.Context, username string, password string) (*model.User, error) {
	cryptedPassword, err := repository.EncryptString(&password, r.CypherKey)
	if err != nil {
		return nil, err
	}

	hasher := sha256.New()
	hasher.Write([]byte(*cryptedPassword))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	u, err := r.db.GetUserByUsernameAndHashedPassword(ctx, database.GetUserByUsernameAndHashedPasswordParams{
		Username:       username,
		HashedPassword: hashedPassword,
	})

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return toDomain(&u), nil
}

func (r *Repository) CreateUser(ctx context.Context, username string, password string) (*model.User, error) {
	cryptedPassword, err := repository.EncryptString(&password, r.CypherKey)
	if err != nil {
		return nil, err
	}

	hasher := sha256.New()
	hasher.Write([]byte(*cryptedPassword))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	u, err := r.db.CreateUser(ctx, database.CreateUserParams{
		Username:       username,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return toDomain(&u), nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int32) (*model.User, error) {
	u, err := r.db.GetUserByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return toDomain(&u), nil
}
