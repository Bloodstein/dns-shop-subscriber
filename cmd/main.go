package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Bloodstein/dns-shop-subscriber-backend/pkg/handler"
	"github.com/Bloodstein/dns-shop-subscriber-backend/pkg/repository"
	"github.com/Bloodstein/dns-shop-subscriber-backend/pkg/service"
)

const (
	subscribes_db_file_path = "database/subscribes.json"
)

func main() {

	_, err := os.ReadFile(subscribes_db_file_path)

	if os.IsNotExist(err) {
		log.Fatalf("fail to read subscribes.json. %s", err.Error())
	}

	repository := repository.NewRepository()
	services := service.NewService(repository)
	handler := handler.NewHandler(services)

	go handler.RunPricesListening()

	if err := http.ListenAndServe(":8000", handler.Routes()); err != nil {
		log.Fatalf("Error to run web server: %s", err.Error())
	}
}
