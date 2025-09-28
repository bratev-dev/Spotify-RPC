package repositorio

import modelos "servidor.local/grpc-servidor/capaModelos"

// Repositorio en memoria
var generos []modelos.Genero

func InicializarDatos() {
	// Género: Salsa
	salsa := modelos.Genero{
		ID:     1,
		Nombre: "Salsa",
	}

	// Canciones de Salsa
	salsa.Canciones = []modelos.Cancion{
		{
			ID:       1,
			Titulo:   "El cantante",
			Artista:  "Héctor Lavoe",
			Album:    "Comedia",
			Anio:     1978,
			Duracion: "4:36",
			GeneroID: salsa.ID,
			Url:      "audio/salsa/el_cantante.mp3",
		},
		{
			ID:       2,
			Titulo:   "Periódico de ayer",
			Artista:  "Héctor Lavoe",
			Album:    "Comedia",
			Anio:     1976,
			Duracion: "5:00",
			GeneroID: salsa.ID,
			Url:      "audio/salsa/periodico.mp3",
		},
	}

	// Género: Rock
	rock := modelos.Genero{
		ID:     2,
		Nombre: "Rock",
	}

	rock.Canciones = []modelos.Cancion{
		{
			ID:       3,
			Titulo:   "Nothing Else Matters",
			Artista:  "Metallica",
			Album:    "Comedia",
			Anio:     1991,
			Duracion: "6:28",
			GeneroID: rock.ID,
			Url:      "audio/rock/nothing_else_matters.mp3",
		},
	}

	// Agregarlos al repositorio
	generos = []modelos.Genero{salsa, rock}
}

func ObtenerGeneros() []modelos.Genero {
	return generos
}

// Función que busca canciones por el ID del género
func ObtenerCancionesPorGenero(id int32) []modelos.Cancion {
	for _, g := range generos {
		if g.ID == id {
			return g.Canciones
		}
	}
	return []modelos.Cancion{}
}
func ObtenerCancionPorID(id int32) *modelos.Cancion {
	for _, genero := range generos {
		for _, cancion := range genero.Canciones {
			if cancion.ID == id {
				return &cancion
			}
		}
	}
	return nil
}
