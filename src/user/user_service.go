package user

import (
	"context"
	"time"

	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
	"github.com/fajarcandraaa/go-mux-crud/helpers"
	"github.com/fajarcandraaa/go-mux-crud/helpers/errorcodehandling"
	repositories "github.com/fajarcandraaa/go-mux-crud/model"
	"github.com/google/uuid"
)

type service struct {
	repo *repositories.Repository
	err  *errorcodehandling.CodeError
}

func NewService(repo *repositories.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) InsertNewUser(ctx context.Context, payload *userentity.UserRequest) (*userentity.User, error) {

	err := userentity.UserRequestValidate(payload)
	if err != nil {
		return nil, err
	}
	hashPassword, _ := helpers.HashPassword(payload.Password)

	user := &userentity.User{
		ID:        uuid.NewString(),
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Phone:     payload.Phone,
		Avatar:    payload.Avatar,
		Email:     payload.Email,
		Username:  payload.Username,
		Password:  hashPassword,
		Status:    payload.Status,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	err = s.repo.User.SaveNewUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) FindUser(ctx context.Context, userID string) (*userentity.FindUser, error) {

	user, err := s.repo.User.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	resultUser := userentity.FindUser{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Phone:     user.Phone,
		Avatar:    user.Avatar,
		Email:     user.Email,
		Username:  user.Username,
		Password:  user.Password,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	return &resultUser, nil
}

func (s *service) GetListUsers(ctx context.Context, sortBy, orderBy string, perPage, page int) (*[]userentity.Users, int64, error) {
	users, total, err := s.repo.User.ListUser(sortBy, orderBy, perPage, page)
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (s *service) UpdateUser(ctx context.Context, payload *userentity.UserData) (*userentity.UserData, error) {
	user, err := s.repo.User.FindUserByID(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	userdata := userentity.SetParameterUpdateUser(user, payload)

	updateUser, err := s.repo.User.UpdateUserData(ctx, payload.ID, userdata)
	if err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (s *service) DeleteDataUser(ctx context.Context, userID string) error {

	err := s.repo.User.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}
