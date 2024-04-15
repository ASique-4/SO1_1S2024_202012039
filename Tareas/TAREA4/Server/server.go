package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	pb "server/Proto"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var ctx = context.Background()
var db *sql.DB

type server struct {
	pb.UnimplementedGetInfoServer
}

const (
	port = ":3001"
)

type Data struct {
	Name  string `json:"name"`
	Album string `json:"album"`
	Year  string `json:"year"`
	Rank  string `json:"rank"`
}

func mysqlConnect() {

	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}

	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	host := os.Getenv("MYSQL_HOST")

	// dsn := "root:secret@tcp(localhost:3306)/clase9"
	dsn := "root:" + pass + "@tcp(" + host + ":3306)/tarea4"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Conexión a MySQL exitosa")
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
	insertMySQL(data)
	return &pb.ReplyInfo{Info: "Hola cliente, recibí el comentario"}, nil
}

func insertMySQL(voto Data) {
	// Prepara la consulta SQL para la inserción en MySQL
	query := "INSERT INTO music (name, album, year, ranking) VALUES (?, ?, ?, ?)"
	_, err := db.ExecContext(ctx, query, voto.Name, voto.Album, voto.Year, voto.Rank)
	if err != nil {
		log.Println("Error al insertar en MySQL:", err)
	}
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	mysqlConnect()

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
