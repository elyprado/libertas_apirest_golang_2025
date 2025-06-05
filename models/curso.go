package models

type Curso struct {
	Idcurso      int      `json:"idcurso"`
	Nome         string   `json:"nome"`
	Descricao    *string  `json:"descricao"`    
	CargaHoraria *string  `json:"cargaHoraria"` 
	Valor        *float64 `json:"valor"`        
}
