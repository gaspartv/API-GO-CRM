package modelsClients

import "github.com/gaspartv/API-GO-CRM/database"

func FindAll() (clients []ClientResponse, err error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM clients`)
	if err != nil {
		return
	}

	for rows.Next() {
		var client ClientResponse

		err = rows.Scan(
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
		if err != nil {
			continue
		}

		clients = append(clients, client)
	}

	return
}
