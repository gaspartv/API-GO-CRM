package modelsClients

import (
	"time"
)

type ClientResponse struct {
	ID        int64      `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Birthday  string     `json:"birthday"`
	Gender    string     `json:"gender"`
	RG        string     `json:"rg"`
	CPF       string     `json:"cpf"`
	Phone     string     `json:"phone"`
	Email     string     `json:"email"`
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateClientResponse struct {
	Message string         `json:"message"`
	Data    ClientResponse `json:"data"`
}

type DeleteClientResponse struct {
	Message string         `json:"message"`
	Data    ClientResponse `json:"data"`
}

type ShowClientResponse struct {
	Message string         `json:"message"`
	Data    ClientResponse `json:"data"`
}

type ListClientsResponse struct {
	Message string           `json:"message"`
	Data    []ClientResponse `json:"data"`
}

type UpdateClientResponse struct {
	Message string         `json:"message"`
	Data    ClientResponse `json:"data"`
}
