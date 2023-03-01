package usermodel

import (
	"context"

	"github.com/fajarcandraaa/go-mux-crud/config/app"
	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
	"github.com/pkg/errors"
)

func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	var user userentity.User

	err := r.db.Model(&user).Where("id = ?", userID).Take(&user).Delete(user).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return app.ErrNoRowsFound
		case app.ErrUniqueViolation:
			return app.ErrUniqueViolation
		default:
			return errors.Wrap(parsed, "build statement query to delete user data from database")
		}
	}

	return nil
}
