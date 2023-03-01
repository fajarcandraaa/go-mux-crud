package usermodel

import (
	"github.com/fajarcandraaa/go-mux-crud/helpers/errorcodehandling"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db        *gorm.DB
	codeError *errorcodehandling.CodeError
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
