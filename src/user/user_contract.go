package user

import (
	"context"

	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
)

type Service interface {
	InsertNewUser(ctx context.Context, payload *userentity.UserRequest) (*userentity.User, error)
}
