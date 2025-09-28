package capacontroladores

import (
	"context"

	pb "servidor.local/grpc-servidor/ServidorDeCanciones/serviciosCancion"
	capafachadaservices "servidor.local/grpc-servidor/capaFachadaServices"
)

type ServidorCanciones struct {
	pb.UnimplementedAudioServiceServer
}

type ControladorServidor struct {
	pb.UnimplementedAudioServiceServer
}

// Implementación del procedimiento remoto
func (s *ControladorServidor) EnviarCancionMedianteStream(req *pb.PeticionDTO, stream pb.AudioService_EnviarCancionMedianteStreamServer) error {

	//Usamos la fachada directamente
	return capafachadaservices.StreamAudioFile(
		req.Titulo,
		func(data []byte) error {
			return stream.Send(&pb.FragmentoCancion{Data: data})
		})
}

func (s *ServidorCanciones) ObtenerCancionesDeGenero(ctx context.Context, req *pb.PeticionGenero) (*pb.ListaCanciones, error) {
	generoID := req.GetIdGenero()

	canciones := capaFachadaServices.ObtenerCancionesPorGenero(generoID)

	var lista []*pb.CancionDTO
	for _, c := range canciones {
		lista = append(lista, &pb.CancionDTO{
			Id:       c.ID,
			Titulo:   c.Titulo,
			Artista:  c.Artista,
			Album:    c.Album,
			Anio:     c.Anio,
			Duracion: c.Duracion,
			GeneroId: c.GeneroID,
		})
	}

	return &pb.ListaCanciones{Canciones: lista}, nil
}
