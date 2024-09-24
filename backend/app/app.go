package app

import (
	"backend/domain"
	"backend/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service: services.InitAudioService(domain.InitAudioRepoStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllAudios).Methods(http.MethodGet)

	log.Println("starting server at port 8000")
	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
