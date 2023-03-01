package userentity

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// User -> initialization user entity
type User struct {
	ID        string    `gorm:"size:36;not null;unique index;primaryKey"`
	Firstname string    `gorm:"size:255;"`
	Lastname  string    `gorm:"size:255;"`
	Phone     string    `gorm:"size:50;"`
	Avatar    string    `gorm:"size:255;"`
	Email     string    `gorm:"size:100;unique"`
	Username  string    `gorm:"size:100;unique"`
	Password  string    `gorm:"size:100;"`
	Status    string    `gorm:"size:50;"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// Users represent body for list users
type Users struct {
	ID        string `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Phone     string `json:"phone_number"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Users represent body for get data from user
type UserData struct {
	ID        string `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Phone     string `json:"phone_number"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserRequest struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Phone     string `json:"phone_number"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Status    string `json:"status"`
}

// UserRequestValidate is to validate input request
func UserRequestValidate(ur *UserRequest) error {
	err := validation.Errors{
		"first_name": validation.Validate(&ur.Firstname, validation.Required, validation.Length(2, 40)),
		"last_name":  validation.Validate(&ur.Lastname, validation.Required),
		"email":      validation.Validate(&ur.Email, validation.Required),
		"username":   validation.Validate(&ur.Username, validation.Required),
		"password":   validation.Validate(&ur.Password, validation.Required),
	}

	return err.Filter()
}

// FindUser is struct to handle respone while find user by ID
type FindUser struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Phone     string `json:"phone_number"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func SetParameterUpdateUser(user *User, payload *UserData) *User {
	var userdata User

	if payload.Firstname == "" {
		userdata.Firstname = user.Firstname
	} else {
		userdata.Firstname = payload.Firstname
	}

	if payload.Lastname == "" {
		userdata.Lastname = user.Lastname
	} else {
		userdata.Lastname = payload.Lastname
	}

	if payload.Phone == "" {
		userdata.Phone = user.Phone
	} else {
		userdata.Phone = payload.Phone
	}

	if payload.Avatar == "" {
		userdata.Avatar = user.Avatar
	} else {
		userdata.Avatar = payload.Avatar
	}

	if payload.Email == "" {
		userdata.Email = user.Email
	} else {
		userdata.Email = payload.Email
	}

	if payload.Username == "" {
		userdata.Username = user.Username
	} else {
		userdata.Username = payload.Username
	}

	if payload.Status == "" {
		userdata.Status = user.Status
	} else {
		userdata.Status = payload.Status
	}
	
	userdata.Password = user.Password
	userdata.UpdatedAt = time.Now()
	
	return &userdata
}