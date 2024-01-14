package handlerClients

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	modelsClients "github.com/gaspartv/API-GO-CRM/models/clients"
	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(
			ctx, http.StatusBadRequest,
			fmt.Sprintf("param: %s (type: %s) is required", id, "queryParameter"),
		)
		return
	}
	number, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("Erro ao converter a string para int64:", err)
		return
	}

	client, err := modelsClients.Find(number)
	if err != nil {
		log.Printf("Erro ao atualizar registro: %v", err)
		sendError(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	sendSuccess(ctx, "find opened successfully", client)
}
