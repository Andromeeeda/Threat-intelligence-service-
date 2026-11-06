package transport_threat_handlers

import (
	domain_threatDomain "Threat-intelligence-service/internal/core/domain"
	common_errors "Threat-intelligence-service/internal/core/errors/common"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type CreateThreatRequest struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
}

type CreateThreatResponce struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
}

func (h *ThreatHandler) CreateThreate(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var createThreatRequest CreateThreatRequest

	if err := json.NewDecoder(r.Body).Decode(&createThreatRequest); err != nil {
		log.Printf("decoding error json: %v", err)

		http.Error(w, "decoding error json", http.StatusBadRequest)

		return
	}

	threat := &domain_threatDomain.Threat{
		Md5:  createThreatRequest.Md5,
		Name: createThreatRequest.Name,
	}

	if err := h.threatService.AddThreat(ctx, threat); err != nil {
		log.Printf("handler CreateThreat error %v", err)

		if errors.Is(err, common_errors.ErrConflict) {
			http.Error(w, "Such a threat already exists", http.StatusConflict)

			return

		} else {
			http.Error(w, "error server", http.StatusInternalServerError)

			return
		}
	}

	responce := &CreateThreatResponce{
		Md5: threat.Md5,
		Name: threat.Name,
	}
	
	log.Println("Threat create")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responce)

}
