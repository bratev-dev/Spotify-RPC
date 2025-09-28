package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "servidor.local/grpc-servidor/ServidorDeCanciones/serviciosCancion"
	capacontroladores "servidor.local/grpc-servidor/capaControladores"
	repositorio "servidor.local/grpc-servidor/capaRepositorios"
)

func main() {
	repositorio.InicializarDatos()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error escuchando en el puerto: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Se registra el controlador que ofrece el procedimiento remoto
	pb.RegisterAudioServiceServer(grpcServer, &capacontroladores.ControladorServidor{})

	fmt.Println("Servidor gRPC escuchando en :50051...")

	// Iniciar el servidor
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar servidor gRPC: %v", err)
	}
}
