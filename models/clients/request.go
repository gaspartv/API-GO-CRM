package modelsClients

import (
	"fmt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type Client struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	RG        string `json:"rg"`
	CPF       string `json:"cpf"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (r *Client) Validate() error {
	if r.FirstName == "" &&
		r.LastName == "" &&
		r.Birthday == "" &&
		r.Gender == "" &&
		r.RG == "" &&
		r.CPF == "" &&
		r.Phone == "" &&
		r.Email == "" {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.FirstName == "" {
		return errParamIsRequired("FirstName", "string")
	}
	if r.LastName == "" {
		return errParamIsRequired("LastName", "string")
	}
	if r.Birthday == "" {
		return errParamIsRequired("Birthday", "string")
	}
	if r.Gender == "" {
		return errParamIsRequired("Gender", "string")
	}
	if r.RG == "" {
		return errParamIsRequired("RG", "string")
	}
	if r.CPF == "" {
		return errParamIsRequired("CPF", "string")
	}
	if r.Phone == "" {
		return errParamIsRequired("Phone", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("Email", "string")
	}
	return nil
}

type UpdateClientRequest struct {
	FirstName string
	LastName  string
	Birthday  string
	Gender    string
	RG        string
	CPF       string
	Phone     string
	Email     string
}

func (r *UpdateClientRequest) Validate() error {
	if r.FirstName != "" ||
		r.LastName != "" ||
		r.Birthday != "" ||
		r.Gender != "" ||
		r.RG != "" ||
		r.CPF != "" ||
		r.Phone != "" ||
		r.Email != "" {
		return fmt.Errorf("request body is empty or malformed")
	}

	return fmt.Errorf("at least one field must be provided")
}
