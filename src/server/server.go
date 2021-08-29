package server

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"shorturl/src/database"
	"shorturl/src/database/models"
	"shorturl/src/grpcgen"
	"shorturl/src/server/services"
)

type Server struct {
	DBInstance *database.DBInstance
}

func (gs *Server) Run() {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		host = ":8080"
	}
	listen, err := net.Listen("tcp", host)
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	sr := services.UrlService{
		UrlModel: &models.UrlModel{
			DBInstance: gs.DBInstance,
		},
	}
	grpcgen.RegisterShortUrlServer(server, &sr)

	log.Println("Starting gRPC server...")
	server.Serve(listen)
}
