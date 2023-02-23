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
