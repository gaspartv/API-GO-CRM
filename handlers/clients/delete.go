package handlerClients

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	modelsClients "github.com/gaspartv/API-GO-CRM/models/clients"
	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
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

	rows, err := modelsClients.Delete(number)
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		sendError(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "registro removido com sucesso!",
	}

	sendSuccess(ctx, "delete opened successfully", resp)
}
