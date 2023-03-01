package repositories

import (
	"context"

	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
	"github.com/fajarcandraaa/go-mux-crud/model/usermodel"
	"github.com/jinzhu/gorm"
)

type User interface {
	SaveNewUser(ctx context.Context, payload *userentity.User) error
	FindUserByID(ctx context.Context, userID string) (*userentity.User, error)
	ListUser(sortBy, orderBy string, perPage, page int) (*[]userentity.Users, int64, error)
	UpdateUserData(ctx context.Context, userID string, payload *userentity.User) (*userentity.UserData, error)
	DeleteUser(ctx context.Context, userID string) error
}

func NewUser(db *gorm.DB) User {
	return usermodel.NewUserRepository(db)
}
