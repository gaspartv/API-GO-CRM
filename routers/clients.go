package routers

import (
	handlerClients "github.com/gaspartv/API-GO-CRM/handlers/clients"
	"github.com/gin-gonic/gin"
)

func initializeClientsRoutes(router *gin.Engine) {
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.POST("/clients", handlerClients.Create)
		v1.DELETE("/clients/:id", handlerClients.Delete)
		v1.PUT("/clients/:id", handlerClients.Update)
		v1.GET("/clients/:id", handlerClients.Find)
		// v1.GET("/clients", handlerClients.FindAll)
	}
}
