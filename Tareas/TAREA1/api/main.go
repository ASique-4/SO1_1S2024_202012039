package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type User struct {
	Nombres   string `json:"nombres"`
	Apellidos string `json:"apellidos"`
	Edad      int    `json:"edad"`
	Carnet    string `json:"carnet"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var user User
	user.Nombres = "Angel Francisco"
	user.Apellidos = "Sique Santos"
	user.Edad = 21
	user.Carnet = "202012039"
	json.NewEncoder(w).Encode(user)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/data", userHandler)
	log.Println("Servidor en el puerto " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
