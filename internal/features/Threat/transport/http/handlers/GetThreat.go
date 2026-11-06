package transport_threat_handlers

import (
	common_errors "Threat-intelligence-service/internal/core/errors/common"
	core_validMd5 "Threat-intelligence-service/internal/core/validMd5"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ThreatResponce struct {
	Md5  string `json:"md5"`
	Name string `json:"name"`
}

func (h *ThreatHandler) GetThreat(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	md5 := vars["md5"]

	if !core_validMd5.IsValidMD5(md5) {
		log.Println("incorrect md5")

		http.Error(w,"inccorrect md5",http.StatusBadRequest)

		return 
	}

	threat,err :=  h.threatService.GetThreat(ctx,md5)

	if err != nil {

		if errors.Is(err,common_errors.ErrNotFound) {
			log.Printf("threat not found %v",err)

			http.Error(w,"threat not found",http.StatusNotFound)

			return
		} else {
			log.Printf("internal server error %v",err)
			
			http.Error(w,"error",http.StatusInternalServerError)

			return 
		}
	}

	responce := &ThreatResponce{
		Md5: threat.Md5,
		Name: threat.Name,
	}

	log.Println("get threat successfully")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responce)

}
