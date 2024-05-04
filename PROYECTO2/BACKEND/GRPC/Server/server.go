package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	pb "Server/Proto"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port = ":5002"
)

type Data struct {
	Name  string `json:"name"`
	Album string `json:"album"`
	Year  string `json:"year"`
	Rank  string `json:"rank"`
}

func publicarAKafka(data Data) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		return err
	}

	defer p.Close()

	// Codifique los datos como bytes
	mensaje, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Cree un mensaje de Kafka
	topic := "topic-sopes1"
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          mensaje,
	}

	// Publique el mensaje
	err = p.Produce(msg, nil)
	if err != nil {
		return err
	}

	// Espere a que el mensaje sea entregado
	p.Flush(15 * 1000)

	return nil
}

type Message struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func enviarJSON(data Data) error {
	env_ruta := os.Getenv("ENV_RUTA_SERVER_GO") // Obtener el valor de la variable de entorno
	if env_ruta == "" {
		return fmt.Errorf("variable de entorno ENV_RUTA_SERVER_GO no configurada")
	}

	message := []Data{data}
	jsonData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	messageToSend := Message{
		Topic:   "topic-sopes1",
		Message: string(jsonData),
	}

	jsonValue, _ := json.Marshal(messageToSend)

	resp, err := http.Post(env_ruta, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Status de la respuesta:", resp.Status)

	// Enviar respuesta correcta
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil

}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println("Recibí de cliente: ", in.GetName())
	data := Data{
		Name:  in.GetName(),
		Album: in.GetAlbum(),
		Year:  in.GetYear(),
		Rank:  in.GetRank(),
	}
	fmt.Println(data)

	// Publique los datos en Kafka
	err := publicarAKafka(data)
	if err != nil {
		log.Println("Error publicando a Kafka:", err)
	}

	// Envía los datos a la URL especificada
	err = enviarJSON(data)
	if err != nil {
		log.Println("Error enviando JSON:", err)
	}

	return &pb.ReplyInfo{Info: "Hola cliente, recibí el comentario"}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	fmt.Println("Servidor corriendo en puerto: ", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
