package usermodel

import (
	"context"

	"github.com/fajarcandraaa/go-mux-crud/config/app"
	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
	"github.com/pkg/errors"
)

func (r *UserRepository) UpdateUserData(ctx context.Context, userID string, payload *userentity.User) (*userentity.UserData, error) {
	var users userentity.User

	err := r.db.Debug().Model(&users).Where("id = ?", userID).Take(&users).UpdateColumns(payload).Error

	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return nil, app.ErrNoRowsFound
		case app.ErrUniqueViolation:
			return nil, app.ErrUniqueViolation
		default:
			return nil, errors.Wrap(parsed, "build statement query to update user data from database")
		}
	}

	userData := &userentity.UserData{
		ID:        userID,
		Firstname: users.Firstname,
		Lastname:  users.Lastname,
		Phone:     users.Phone,
		Avatar:    users.Avatar,
		Email:     users.Email,
		Username:  users.Username,
		Status:    users.Status,
		CreatedAt: users.CreatedAt.String(),
		UpdatedAt: users.UpdatedAt.String(),
	}

	return userData, nil
}
