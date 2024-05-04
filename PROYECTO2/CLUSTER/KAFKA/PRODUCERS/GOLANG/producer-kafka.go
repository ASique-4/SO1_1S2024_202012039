package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

// Cargar configuracion de kafka
type ConfigKafka struct {
	Port         string
	KafkaBrokers []string
}

var configKafka ConfigKafka

// Cargar configuracion de kafka
func cargarConf() ConfigKafka {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return ConfigKafka{
		Port:         os.Getenv("PORT"),
		KafkaBrokers: strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
	}
}

func main() {
	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Configuración
	configKafka = cargarConf() // Asigna el valor retornado por loadConfig a la variable config

	// Rutas
	http.HandleFunc("/", inicioHandler)
	http.HandleFunc("/sendMessage", enviarMensajeHandler)

	// Iniciar servidor
	fmt.Printf("Listening on port %s\n", configKafka.Port)
	log.Fatal(http.ListenAndServe(":"+configKafka.Port, nil))
}

func inicioHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a KAFKA-API - 202012039")
}

func enviarMensajeHandler(w http.ResponseWriter, r *http.Request) {
	// Leer cuerpo de la solicitud
	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Conexión a Kafka
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: configKafka.KafkaBrokers,
		Topic:   requestData["topic"],
	})

	// Enviar mensaje
	err := writer.WriteMessages(r.Context(),
		kafka.Message{
			Key:   []byte("key"),
			Value: []byte(requestData["message"]),
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mensaje de éxito
	log.Println("Datos enviados correctamente a Kafka.")

	// Respuesta
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Mensaje enviado correctamente a Kafka. :D"}`)
}
