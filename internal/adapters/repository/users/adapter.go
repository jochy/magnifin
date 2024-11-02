package users

import (
	"magnifin/internal/app/model"
	"magnifin/internal/infra/database"
)

func toDomain(user *database.User) *model.User {
	return &model.User{
		ID:       user.ID,
		Username: user.Username,
	}
}
