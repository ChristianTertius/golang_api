package main

import (
	"building_an_api/internal/handlers"
	"fmt"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	r := chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service ...")
	fmt.Println(`=================== THIS IS GO API ==============`)

	err := http.ListenAndServe("localhost:8000", r)

	if err != nil {
		log.Error(err)
	}
}
