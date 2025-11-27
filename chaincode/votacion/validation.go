package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

)

// Constantes de validación

const (
	MaxNombreLength = 100
	MinUIDLength       = 10
)
// Validadores básicos
func ValidateUI(ui string) error {
	if utf8.RuneCountInString(ui) < MinUIDLength {
		return ValidationError{
			Field:   "UI",
			Message: fmt.Sprintf("debe tener minimo %d caracteres", MinUIDLength),
			Code:    "INVALID_UI_LENGTH",
		}
	}
	
	if match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", ui); !match {
		return ValidationError{
			Field:   "UI",
			Message: "debe ser alfanumérico (solo letras y números)",
			Code:    "INVALID_UI_FORMAT",
		}
	}
	return nil
}

func ValidateNombre(nombre string) error {
	nombre = strings.TrimSpace(nombre)
	if utf8.RuneCountInString(nombre) == 0 {
		return ValidationError{
			Field:   "Nombre",
			Message: "no puede estar vacío",
			Code:    "EMPTY_NAME",
		}
	}
	if utf8.RuneCountInString(nombre) > MaxNombreLength {
		return ValidationError{
			Field:   "Nombre",
			Message: fmt.Sprintf("no puede exceder %d caracteres", MaxNombreLength),
			Code:    "NAME_TOO_LONG",
		}
	}
	return nil
}

func ValidateFechas(inicio, fin time.Time) error {

	
    currentTime := time.Now().UTC()

	// Validar que las fechas sean en el futuro y coherentes
	if  currentTime.After(inicio){
		return ValidationError{
			Field:   "Inicio",
			Message: "debe ser una fecha futura",
			Code:    "PAST_START_DATE",
		}
	}
	
	if currentTime.After(fin){
		return ValidationError{
			Field:   "Inicio",
			Message: "debe ser una fecha futura",
			Code:    "PAST_END_DATE",
		}
	}

	if inicio.After(fin){
		return ValidationError{
			Field:   "Fin",
			Message: "debe ser posterior a la fecha de inicio",
			Code:    "INVALID_END_DATE",
		}
	}

	return nil
}

func ValidateVotationStarted(votacion *Votacion) error {
	
	currentTime := time.Now().UTC()
	if currentTime.Before(votacion.Inicio) {
		return ValidationError{
		Field:   "Votacion",
		Message: "la votación no ha iniciado",
		Code:    "VOTACION_NO_INICIADA",
	}
	}
	if currentTime.After(votacion.Fin) {
		return ValidationError{
		Field:   "Votacion",
		Message: "la votación ha finalizado",
		Code:    "VOTACION_TERMINADA",
	}
	} 
	return nil
}	

func ValidateBeforeStart(votacion *Votacion) error {
	currentTime := time.Now().UTC()	
	if currentTime.After(votacion.Inicio) {
		if currentTime.After(votacion.Fin) {
			return ValidationError{
				Field:   "Votacion",
				Message: "la votación ha finalizado",
				Code:    "VOTACION_TERMINADA",
			}
		}
		return ValidationError{
			Field:   "Votacion",
			Message: "la votación a iniciado",
			Code:    "VOTACION_INICIADA",
		}
	}
	return nil


}

func ValidateAfterEnd(votacion *Votacion) error {
	currentTime := time.Now().UTC()	
	
	if currentTime.Before(votacion.Inicio) {
		return ValidationError{
			Field:   "Votacion",
			Message: "la votación no a iniciado",
			Code:    "VOTACION_NO_INICIADA",
		}
	}
	if currentTime.After(votacion.Fin) {
		return ValidationError{
			Field:   "Votacion",
			Message: "la votación ha finalizado",
			Code:    "VOTACION_TERMINADA",
		}
	}
	
	return nil


}

// Validadores específicos de dominio
func ValidateCandidato(uiCandidato string, candidatosExistentes *CandidatosVotacion) error {
	if err := ValidateUI(uiCandidato); err != nil {
		return err
	}

	// Validar unicidad
	for _, candidato := range candidatosExistentes.Candidatos {
		if candidato.UI == uiCandidato {
			return ValidationError{
				Field:   "UI",
				Message: "el candidato ya existe en esta votación",
				Code:    "DUPLICATE_CANDIDATE",
			}
		}
	}
	return nil
}

// Validar si el votante ya ha votado y si está registrado
func ValidateVoterCanVote(uiVotante string, votacion *Votacion, votosExistentes *VotosVotacion) error {
	if err := ValidateUI(uiVotante); err != nil {
		return err
	}

	// Validar que el votante esté registrado en la votación
	isRegistered := false
	for _, votante := range votacion.Votantes.Votantes {
		if votante.UI == uiVotante {
			isRegistered = true
			break
		}
	}
	if !isRegistered {
		return ValidationError{
			Field:   "UI",
			Message: "el votante no está registrado en esta votación",
			Code:    "VOTER_NOT_REGISTERED",
		}
	}

	// Validar que el votante no haya votado ya
	for _, voto := range votosExistentes.Votos {
		if voto.UIVotante == uiVotante {
			return ValidationError{
				Field:   "UI",
				Message: "el votante ya ha emitido su voto en esta votación",
				Code:    "VOTER_ALREADY_VOTED",
			}
		}
	}

	return nil
}