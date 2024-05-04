package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/segmentio/kafka-go"
)

// Cargar configuracion de kafka
type ConfigKafka struct {
	Port         string
	Topic        string
	KafkaBrokers []string
}

// Estructura de un mensaje
type Message struct {
	Topic string
	Value string
}

// Cargar configuracion de kafka
func CargaConfig() ConfigKafka {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return ConfigKafka{
		Port:         os.Getenv("PORT"),
		KafkaBrokers: strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		Topic:        os.Getenv("KAFKA_TOPIC"),
	}
}

func main() {

	config := CargaConfig()
	fmt.Println(config)

	// Crear un nuevo consumidor
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   config.KafkaBrokers,
		Topic:     config.Topic,
		GroupID:   "test-group",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	// Escucha señales de interrupción para cerrar el consumidor adecuadamente
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Escuchar mensajes
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error al leer mensaje: ", err)
		}
		fmt.Printf("Mensaje recibido: %s\n", string(m.Value))
	}
}
