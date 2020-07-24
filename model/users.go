package model

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"io"
)

type User struct {
	Firstname   string             `json:"firstname" validate:"required"`
	Middlename  string             `json:"middlename"`
	Lastname    string             `json:"lastname"`
	Email       string             `json:"email"`
	CV 			string 				`json:"cv"`
	WorkHistory []*EmployerHistory `json:"workHistory"`
}

func (u *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(u)
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

type EmployerHistory struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Link    string `json:"link"`
}

type ApplicationResponse struct {
	Success bool
	Message string
	User    *User `json:"user"`
}



