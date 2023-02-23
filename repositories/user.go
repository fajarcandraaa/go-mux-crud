package repositories

import (
	"context"

	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
	"github.com/fajarcandraaa/go-mux-crud/repositories/userrepo"
	"github.com/jinzhu/gorm"
)

type User interface {
	SaveNewUser(ctx context.Context, payload *userentity.User) error
}

func NewUser(db *gorm.DB) User {
	return userrepo.NewUserRepository(db)
}
