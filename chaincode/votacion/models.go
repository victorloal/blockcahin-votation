package main

import (
	"time"
)

// Estructuras existentes (las que ya tienes)
type Candidato struct {
	UI       string `json:"UI"`
	Posicion int    `json:"Posicion"`
}

type CandidatosVotacion struct {
	UIVotacion string      `json:"UIVotacion"`
	Candidatos []Candidato `json:"Candidatos"`
}

type Voto struct {
	Fecha   string   `json:"Fecha"`
	Voto    []string `json:"Voto"`
	Votante string   `json:"Votante"`
}

type Resultado struct {
	Resultado  []string `json:"Resultado"`
	TotalVotos string   `json:"TotalVotos"`
}

type Votacion struct {
	UI     string `json:"Identidicador Unico"`
	Nombre string `json:"Nombre de la Votación"`
	Inicio time.Time `json:"Fecha de Inicio"`
	Fin    time.Time  `json:"Fecha de Finalización"`
}

type Voters struct {
	Votantes  []string `json:"Votantes"`
}

// Nuevas estructuras para validación
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (e ValidationError) Error() string {
	return e.Message
}

type ValidationContext struct {
	TxTimestamp time.Time
	Caller      string
}