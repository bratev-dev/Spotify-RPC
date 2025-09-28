package capaFachadaServices

import (
	"servidor.local/grpc-servidor/ServidorDeCanciones/capaModelos"
	"servidor.local/grpc-servidor/ServidorDeCanciones/capaRepositorios"
)

func ObtenerCancionesPorGenero(id int32) []capaModelos.Cancion {
	// Aquí puedes hacer validaciones si lo deseas
	return capaRepositorios.ObtenerCancionesPorGenero(id)
}
