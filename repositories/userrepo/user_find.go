package userrepo

import (
	"context"

	"github.com/fajarcandraaa/go-mux-crud/config/app"
	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
	"github.com/pkg/errors"
)

func (r *UserRepository) FindUserByID(ctx context.Context, userID string) (*userentity.User, error) {
	var user userentity.User
	err := r.db.First(&user, "id = ?", userID).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return nil, app.ErrNoRowsFound
		case app.ErrUniqueViolation:
			return nil, app.ErrUniqueViolation
		default:
			return nil, errors.Wrap(parsed, "build statement query to find user from database")
		}
	}
	return &user, nil
}
