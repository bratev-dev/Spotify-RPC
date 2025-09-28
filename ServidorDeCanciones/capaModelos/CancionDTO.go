package modelos

type CancionDTO struct {
	ID          int32
	Titulo      string
	Artista     string
	Album       string
	Anio        int32
	Duracion    string
	Genero      GeneroDTO
	RutaArchivo string
}
