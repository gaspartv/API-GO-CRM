package modelsClients

import (
	"time"

	"github.com/gaspartv/API-GO-CRM/database"
)

func Update(id int64, client Client) (int64, error) {
	conn, err := database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `UPDATE clients SET updatedAt=$1, firstName=$2, lastName=$3, birthday=$4, gender=$5, rg=$6, cpf=$7, phone=$8, email=$9 WHERE id=$10;`

	res, err := conn.Exec(
		sql,
		time.Now(), client.FirstName, client.LastName, client.Birthday, client.Gender, client.RG, client.CPF, client.Phone, client.Email, id,
	)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
