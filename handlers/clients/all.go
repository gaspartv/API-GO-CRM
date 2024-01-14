package handlerClients

import (
	"encoding/json"
	"log"
	"net/http"

	modelsClients "github.com/gaspartv/API-GO-CRM/models/clients"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	clients, err := modelsClients.FindAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(clients); err != nil {
		log.Printf("Erro ao atualizar registro 2: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
