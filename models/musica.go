package models

type Musica struct {
	ID      int     `json:"id"`
	Titulo  string  `json:"titulo"`
	Artista string  `json:"artista"`
	Album   string  `json:"album"`
	Ano     int     `json:"ano"`
	Genero  string  `json:"genero"`
}
