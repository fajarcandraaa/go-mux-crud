package usermodel

import (
	"github.com/fajarcandraaa/go-mux-crud/config/app"
	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
	"github.com/pkg/errors"
)

func (r *UserRepository) ListUser(sortBy, orderBy string, perPage, page int) (*[]userentity.Users, int64, error) {
	var users []userentity.User
	var usersGet []userentity.Users
	var count int64

	offset := (page - 1) * perPage
	order := orderBy + " " + sortBy

	err := r.db.Model(&users).
		Order(order).
		Limit(perPage).
		Offset(offset).
		Find(&usersGet).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return nil, 0, app.ErrNoRowsFound
		case app.ErrUniqueViolation:
			return nil, 0, app.ErrUniqueViolation
		default:
			return nil, 0, errors.Wrap(parsed, "build statement query to find user from database")
		}
	}

	count = int64(len(users))

	return &usersGet, count, nil

}
