package modelos

type RespuestaDTO[T any] struct {
	ObjCancion Cancion
	Codigo     int32
	Mensaje    string
}
