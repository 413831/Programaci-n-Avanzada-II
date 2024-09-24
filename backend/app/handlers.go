package app

import (
	"backend/services"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type CustomerHandlers struct {
	service services.AudioService
}

func (ch *CustomerHandlers) getAllAudios(w http.ResponseWriter, r *http.Request) {
	audios, _ := ch.service.GetAllAudios()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(audios)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(audios)
	}
}
