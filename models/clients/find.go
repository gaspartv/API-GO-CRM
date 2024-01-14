package modelsClients

import (
	"github.com/gaspartv/API-GO-CRM/database"
)

func Find(id int64) (client ClientResponse, err error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT * FROM clients WHERE id = $1;`

	row := conn.QueryRow(sql, id)

	err = row.Scan(
		&client.ID,
		&client.CreatedAt,
		&client.UpdatedAt,
		&client.DeletedAt,
		&client.FirstName,
		&client.LastName,
		&client.Birthday,
		&client.Gender,
		&client.RG,
		&client.CPF,
		&client.Phone,
		&client.Email,
	)

	return
}
