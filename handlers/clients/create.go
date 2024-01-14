package handlerClients

import (
	"fmt"
	"net/http"

	modelsClients "github.com/gaspartv/API-GO-CRM/models/clients"
	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context) {
	var client modelsClients.Client

	ctx.BindJSON(&client)

	if err := client.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := modelsClients.Insert(client)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("fixed an error when trying to insert: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("client created successful ID: %d", id),
		}
	}

	sendSuccess(ctx, "create-opening", resp)
}
