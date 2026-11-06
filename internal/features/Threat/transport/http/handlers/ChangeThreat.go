package transport_threat_handlers

import (
	domain_threatDomain "Threat-intelligence-service/internal/core/domain"
	common_errors "Threat-intelligence-service/internal/core/errors/common"
	core_validMd5 "Threat-intelligence-service/internal/core/validMd5"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ChangeThreatRequest struct {
	Name string `json:"name"`
}

type ChangeThreatResponce struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
}

func (h *ThreatHandler) ChangeThreat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	md5 := vars["md5"]

	if !core_validMd5.IsValidMD5(md5) {
		log.Println("incorrect md5")

		http.Error(w, "inccorrect md5", http.StatusBadRequest)

		return
	}

	var request ChangeThreatRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("decoding error json: %v", err)

		http.Error(w, "decoding error json", http.StatusBadRequest)

		return
	}

	changeThreat := &domain_threatDomain.Threat{
		Md5:  md5,
		Name: request.Name,
	}

	newChangeThreat, err := h.threatService.Change(ctx, changeThreat)

	if err != nil {
		if errors.Is(err, common_errors.ErrNotFound) {
			log.Printf("threat not found %v", err)

			http.Error(w, "threat not found", http.StatusNotFound)

			return
		} else {
			log.Printf("internal server error %v", err)

			http.Error(w, "error", http.StatusInternalServerError)

			return
		}
	}

	responce := &ChangeThreatResponce{
		Md5: newChangeThreat.Md5,
		Name:  newChangeThreat.Name,
	}

	log.Println("change threat successfully!")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responce)
}
