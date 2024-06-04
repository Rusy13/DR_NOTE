package dto

import (
	"awesomeProject/internal/user/model"
	"fmt"
	"github.com/asaskevich/govalidator"
	"time"
)

type AddUserDTO struct {
	Name     string `json:"name" valid:"required"`
	Email    string `json:"email" valid:"required,email"`
	Birthday string `json:"birthday" valid:"required"`
}

func (a *AddUserDTO) Validate() error {
	_, err := govalidator.ValidateStruct(a)
	if err != nil {
		fmt.Printf("Validation errors: %v\n", err)
	}
	return err
}

func ConvertToUser(b AddUserDTO) (model.User, error) {
	birthday, err := time.Parse("2006-01-02", b.Birthday)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Name:     b.Name,
		Email:    b.Email,
		Birthday: birthday,
	}, nil
}
