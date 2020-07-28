package service

import (
	"errors"
	"stark/domain"
	"stark/model"
)

type UserService interface {
	Apply(input model.User) (*model.User, error)
	GetApplications() ([]*model.User, error)
	Validate(user *model.User) error
	ValidateTwo(work *model.EmployerHistory) error
}

type service struct {
	userDomain domain.UserDomain
}

func NewUserService(userDomain domain.UserDomain) UserService { return &service{userDomain} }

func (s *service) Apply(input model.User) (*model.User, error) {
	return s.userDomain.CreateUser(input)
}

func (s *service) GetApplications() ([]*model.User, error) {
	return s.userDomain.GetUsers()
}

func (s *service) Validate(user *model.User) error {
	if user == nil {
		err := errors.New("user can not be empty")
		return err
	}

	if user.Firstname == "" {
		err := errors.New("firstname is required")
		return err
	}

	if user.Lastname == "" {
		err := errors.New("lastname is required")
		return err
	}

	if user.Middlename == "" {
		err := errors.New("middlename is required")
		return err
	}

	if user.Email == "" {
		err := errors.New("email is required")
		return err
	}

	//if user.WorkHistory == nil {
	//	err := errors.New("work history can not be empty")
	//	return err
	//}
	return nil
}

func (s *service) ValidateTwo(work *model.EmployerHistory) error {
	if work.Name == "" {
		err := errors.New("company name is required")
		return err
	}

	if work.Email == "" {
		err := errors.New("company email is required")
		return err
	}

	if work.Address == "" {
		err := errors.New("company address is required")
		return err
	}

	if work.Link == "" {
		err := errors.New("company website link is required")
		return err
	}

	if work.Phone == "" {
		err := errors.New("company phone number is required")
		return err
	}
	return nil
}

