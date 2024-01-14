package modelsClients

import (
	"time"

	"github.com/gaspartv/API-GO-CRM/database"
)

func Insert(client Client) (id int64, err error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO clients (createdAt, updatedAt, firstName, lastName, birthday, gender, rg, cpf, phone, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	err = conn.QueryRow(sql, time.Now(), time.Now(), client.FirstName, client.LastName, client.Birthday, client.Gender, client.RG, client.CPF, client.Phone, client.Email).Scan(&id)

	return
}
