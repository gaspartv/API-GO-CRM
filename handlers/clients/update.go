package handlerClients

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	modelsClients "github.com/gaspartv/API-GO-CRM/models/clients"
	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	var client modelsClients.Client

	ctx.BindJSON(&client)

	id := ctx.Query("id")
	if id == "" {
		fmt.Errorf("param: %s (type: %s) is required", "id", "queryParameter")
		return
	}

	number, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Erro ao converter a string para int64:", err)
		return
	}

	rows, err := modelsClients.Update(number, client)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		sendError(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "dados atualizados com sucesso!",
	}

	sendSuccess(ctx, "delete opened successfully", resp)
}
