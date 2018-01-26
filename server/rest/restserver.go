package rest

import (
	"uploadS3/config"
	"fmt"
	"uploadS3/api"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func StartServer() {
	params := config.GetParameters()
	fmt.Printf(fmt.Sprintf("Starting Rest Server on port %s\n", params.RestServer.Port))
	fmt.Printf("Waiting for requests\n")

	r := mux.NewRouter()
	initRoutes(r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", params.RestServer.Address, params.RestServer.Port), r))
}

func initRoutes(r *mux.Router) {
	r.HandleFunc("/api/add", api.AddSoftware).Methods("PUT")
}