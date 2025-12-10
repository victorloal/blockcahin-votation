package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type VotacionContract struct {
	contractapi.Contract
}

// ==================================================================
// =========================== Votacion =============================
// ==================================================================

func (s *VotacionContract) NuevaVotacion(ctx contractapi.TransactionContextInterface, ui string, nombre string, inicio time.Time, fin time.Time) error {
	// Validaciones centralizadas
	if err := ValidateUI(ui); err != nil {
		return err
	}
	
	if err := ValidateNombre(nombre); err != nil {
		return err
	}
	
	if err := ValidateFechas(inicio, fin); err != nil {
		return err
	}

	// Verificar que no exista
	existente, _ := s.ObtenerVotacion(ctx, ui)
	if existente != nil {
		return ValidationError{
			Field:   "UI",
			Message: "ya existe una votación con este identificador",
			Code:    "DUPLICATE_VOTACION",
		}
	}

	// Crear votación
	votacion := Votacion{
		UI:     ui,
		Nombre: nombre,
		Inicio: inicio,
		Fin:    fin,
	}
	
	votacionJSON, err := json.Marshal(votacion)
	if err != nil {
		return fmt.Errorf("error serializando la votación: %v", err)
	}

	return ctx.GetStub().PutState(ui, votacionJSON)
}

// ==================================================================
// =========================== Candidatos ===========================
// ==================================================================



func (s *VotacionContract) AgregarCandidato(ctx contractapi.TransactionContextInterface, uiVotacion string, uiCandidato string) error {
	//Validar UIvotacion 
	if err := ValidateUI(uiVotacion); err != nil {
		return err
	}
	
	//Validar UIcandidato
	if err := ValidateUI(uiCandidato); err != nil {
		return err
	}

	// Validar estado de votación
	votacion, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return err
	}
	
	// Validar si ya inicio la votación
	if err := ValidateBeforeStart(votacion); err != nil {
		return err
	}

	// Obtener y validar candidatos existentes
	candidatosVotacion, err := s.ObtenerCandidatos(ctx, uiVotacion)
	if err != nil {
		return err
	}

	if candidatosVotacion == nil {
		candidatosVotacion = &CandidatosVotacion{
			UIVotacion: uiVotacion,
			Candidatos: []Candidato{},
		}
	}
	

	// Validar candidato
	if err := ValidateCandidato(uiCandidato, candidatosVotacion); err != nil {
		return err
	}

	// Agregar candidato
	posicion := len(candidatosVotacion.Candidatos)
	candidatosVotacion.Candidatos = append(candidatosVotacion.Candidatos, Candidato{
		UI:       uiCandidato,
		Posicion: posicion,
	})

	candidatosActualizadosJSON, err := json.Marshal(candidatosVotacion)
	if err != nil {
		return fmt.Errorf("error serializando candidatos: %v", err)
	}

	return ctx.GetStub().PutState("CANDIDATOS_"+uiVotacion, candidatosActualizadosJSON)
}

// Funcion para eliminar un candidato
func (s *VotacionContract) EliminarCandidato(ctx contractapi.TransactionContextInterface, uiVotacion string, uiCandidato string) error {
	// Verificar si la votación aun no inicia
	votacion, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return err
	}

	if err := ValidateBeforeStart(votacion); err != nil {
		return err
	}

	// Obtener y validar candidatos existentes
	candidatosVotacion, err := s.ObtenerCandidatos(ctx, uiVotacion)
	if err != nil {
		return err
	}

	// Buscar y eliminar el candidato
	found := false
	for i, candidato := range candidatosVotacion.Candidatos {
		if candidato.UI == uiCandidato {
			// Eliminar el candidato de la lista
			candidatosVotacion.Candidatos = append(candidatosVotacion.Candidatos[:i], candidatosVotacion.Candidatos[i+1:]...)
			found = true
			break
		}
	}

	// Reordenar las posiciones de los candidatos restantes
	for i := range candidatosVotacion.Candidatos {
		candidatosVotacion.Candidatos[i].Posicion = i
	}	

	if !found {
		return ValidationError{
			Field:   "UI",
			Message: "el candidato no existe en esta votación",
			Code:    "CANDIDATE_NOT_FOUND",
		}
	}
	// Serializar la lista de candidatos actualizada
	candidatosActualizadosJSON, err := json.Marshal(candidatosVotacion)
	if err != nil {
		return fmt.Errorf("error serializando la lista de candidatos actualizada: %v", err)
	}
	// Guardar la lista de candidatos actualizada en el blockchain
	return ctx.GetStub().PutState("CANDIDATOS_"+uiVotacion, candidatosActualizadosJSON)
}

// ==================================================================
// =========================== Votos ================================
// ==================================================================
//Cambiar la forma de enviar voto


func (s *VotacionContract) Votar(ctx contractapi.TransactionContextInterface, uiVotacion string,idVotante string, voto interface{}) error {

    var votos []string

	// Manejar diferentes tipos de entrada
    switch v := voto.(type) {
    case string:
        // Si es string JSON, deserializar
        if v[0] == '[' {
            err := json.Unmarshal([]byte(v), &votos)
            if err != nil {
                return fmt.Errorf("error deserializando votos: %v", err)
            }
        } else {
            // Si es un solo voto como string
            votos = []string{v}
        }
    case []string:
        // Si ya es array (cuando se llama internamente)
        votos = v
    default:
        return fmt.Errorf("tipo de voto no soportado: %T", v)
    }



	// Validar estado de votación
	votacion, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return err
	}
	
	// Validar si la votación ha iniciado
	if err := ValidateVotationStarted(votacion); err != nil {
		return err
	}

	// Obtener la lista de votantes
	votantes, err := s.ObtenerVotantes(ctx,uiVotacion)
	
	if err := ValidateVotationStarted(votacion); err != nil {
		return err
	}
	aux :=false
	// Verificar si el votante ya está en la lista
	for _, existingID := range votantes.Votantes {
		if existingID == idVotante {
			aux = true
		}
	}

	if !aux {
		return ValidationError{
				Field:   "Vote",
				Message: "el votante no esta registrado para esta votación",
				Code:    "DUPLICATE_VOTER",
			}
	}

	votos2, err := s.ObtenerListaVotos(ctx,uiVotacion)
	for _, existingID := range votos2 {
		if existingID.Votante == idVotante {
			return ValidationError{
				Field:   "Vote",
				Message: "el votante ya registro su voto en esta votación",
				Code:    "DUPLICATE_VOTER",
			}
		}
	}


	//Validar largo del string voto
	if err := ValidateVotationStarted(votacion); err != nil {
		return err
	}

	// Obtener el certificado X509 del cliente
    cert, err := cid.GetX509Certificate(ctx.GetStub())
    if err != nil {
        return ValidationError{
			Field:   "Votante",
			Message: fmt.Sprintf("error obteniendo el certificado X509 del cliente: %v", err),
			Code:    "CERTIFICATE_ERROR",
		}
		
    }
    
    if cert == nil {
        return ValidationError{
			Field:   "Votante",
			Message: "el certificado X509 del cliente es nulo",
			Code:    "NULL_CERTIFICATE",
		}
    }
    
	// Obterner el nombre del votante desde los atributos del certificado
	


	// Validar votante único
	if voto, _ := ObtenerVoto(ctx, uiVotacion, idVotante); voto != nil {
		return ValidationError{
			Field:   "Votante",
			Message: "ya ha emitido su voto en esta votación",
			Code:    "DUPLICATE_VOTE",
		}
	}

	// Validar voto
	candidatosVotacion, err := s.ObtenerCandidatos(ctx, uiVotacion)
	if err != nil {
		return err
	}

	if candidatosVotacion == nil || len(candidatosVotacion.Candidatos) == 0 {
		return ValidationError{
			Field:   "Candidatos",
			Message: "no hay candidatos disponibles para esta votación",
			Code:    "NO_CANDIDATES",
		}
	}

	// Registrar voto
	votoStruct := Voto{
		Fecha:   time.Now().Format("2006-01-02 15:04:05"),
		Voto:    votos,
		Votante: idVotante,
	}
	compositeKey , err := ctx.GetStub().CreateCompositeKey("VOTO", []string{uiVotacion, idVotante})

	if err != nil {
		return err
	}


	votoJSON, err := json.Marshal(votoStruct)
	if err != nil {
		return fmt.Errorf("error serializando voto: %v", err)
	}

	return ctx.GetStub().PutState(compositeKey, votoJSON)
}

// ==================================================================
// ========================= Resultados =============================
// ==================================================================

func (s *VotacionContract) RegistrarResultados(ctx contractapi.TransactionContextInterface, uiVotacion string, resultado []string, totalVotos string) error {
	// Validar estado de votación
	votacion, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return err
	}
	
	if err := ValidateAfterEnd(votacion); err != nil {
		return err
	}
	// Registrar resultados
	resultados := Resultado{
		Resultado:  resultado,
		TotalVotos: totalVotos,
	}
	
	resultadosJSON, err := json.Marshal(resultados)
	if err != nil {
		return fmt.Errorf("error serializando los resultados: %v",
		err)
	}
	
	return ctx.GetStub().PutState("RESULTADOS"+uiVotacion, resultadosJSON)

}

// ==================================================================
// ==================== Votantes ==================================
// ==================================================================

func (s *VotacionContract) AgregarVotante(ctx contractapi.TransactionContextInterface, uiVotacion string, idVotante string) error {
	// Validar UIvotacion 
	if err := ValidateUI(uiVotacion); err != nil {
		return err
	}
	

	// Obtener la lista de votantes existente
	votersJSON, err := ctx.GetStub().GetState("VOTANTES_" + uiVotacion)
	if err != nil {
		return fmt.Errorf("error obteniendo los votantes: %v", err)
	}

	var voters Voters
	if votersJSON != nil {
		err = json.Unmarshal(votersJSON, &voters)
		if err != nil {
			return fmt.Errorf("error deserializando la lista de votantes: %v", err)
		}
	} else {
		voters = Voters{
			Votantes: []string{},
		}
	}

	// Verificar si el votante ya está en la lista
	for _, existingID := range voters.Votantes {
		if existingID == idVotante {
			return ValidationError{
				Field:   "Votante",
				Message: "el votante ya está registrado para esta votación",
				Code:    "DUPLICATE_VOTER",
			}
		}
	}

	// Agregar el nuevo votante a la lista
	voters.Votantes = append(voters.Votantes, idVotante)

	// Serializar y guardar la lista actualizada de votantes
	updatedVotersJSON, err := json.Marshal(voters)
	if err != nil {
		return fmt.Errorf("error serializando la lista de votantes actualizada: %v", err)
	}

	return ctx.GetStub().PutState("VOTANTES_"+uiVotacion, updatedVotersJSON)
}


// ==================================================================
// ==================== Funciones Auxiliares ========================
// ==================================================================

// (Mantén tus funciones ObtenerVotacion, ObtenerCandidatos, ObtenerVoto, etc. aquí)
// Solo asegúrate de que ObtenerVoto sea exportada (con mayúscula) para que pueda ser usada en validation.go

func ObtenerVoto(ctx contractapi.TransactionContextInterface, uiVotacion string, votante string) (*Voto, error) {
	votoJSON, err := ctx.GetStub().GetState(uiVotacion + "_" + votante)
	if err != nil {
		return nil, fmt.Errorf("no se pudo recuperar el voto: %v", err)
	}
	if votoJSON == nil {
		return nil, nil
	}

	var voto Voto
	err = json.Unmarshal(votoJSON, &voto)
	if err != nil {
		return nil, fmt.Errorf("error deserializando el voto: %v", err)
	}

	return &voto, nil
}

//Obtener resultados
func (s *VotacionContract) ObtenerResultados(ctx contractapi.TransactionContextInterface, uiVotacion string) (*Resultado, error) {
	
	ResultadosJSON, err := ctx.GetStub().GetState("RESULTADOS"+uiVotacion)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo los Resultados: %v", err)
	}
	if ResultadosJSON == nil {
		return nil, fmt.Errorf("la lista de Resultados para la votación %s no existe", uiVotacion)
	}
	var Resultado Resultado
	err = json.Unmarshal(ResultadosJSON, &Resultado)
	if err != nil {
		return nil, fmt.Errorf("error deserializando la lista de resultados: %v", err)
	}

	return &Resultado, nil
}

//Obtener lista de votos
func (s *VotacionContract) ObtenerListaVotos(ctx contractapi.TransactionContextInterface, uiVotacion string) ([]Voto, error) {
    // Verificar que la votación existe
    _, err := s.ObtenerVotacion(ctx, uiVotacion)
    if err != nil {
        return nil, err
    }
    
    // Buscar todos los votos usando partial composite key
    resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey("VOTO", []string{uiVotacion})
    if err != nil {
        return nil, fmt.Errorf("error obteniendo votos: %v", err)
    }
    defer resultsIterator.Close()
    
    var votos []Voto
    for resultsIterator.HasNext() {
        queryResponse, err := resultsIterator.Next()
        if err != nil {
            return nil, fmt.Errorf("error iterando votos: %v", err)
        }
        
        var voto Voto
        err = json.Unmarshal(queryResponse.Value, &voto)
        if err != nil {
            return nil, fmt.Errorf("error deserializando voto: %v", err)
        }
        votos = append(votos, voto)
    }
    
    return votos, nil
}

//obtener candidatos
func (s *VotacionContract) ObtenerCandidatos(ctx contractapi.TransactionContextInterface, uiVotacion string) (*CandidatosVotacion, error) {
	// Definir la clave de inicio y fin
	candidatosJSON, err := ctx.GetStub().GetState("CANDIDATOS_"+uiVotacion)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo los candidatos: %v", err)
	}
	if candidatosJSON == nil {
		return &CandidatosVotacion{
			UIVotacion: uiVotacion,
			Candidatos: []Candidato{},
		}, nil
	}
	var candidatos CandidatosVotacion
	err = json.Unmarshal(candidatosJSON, &candidatos)
	if err != nil {
		return nil, fmt.Errorf("error deserializando la lista de candidatos: %v", err)
	}

	return &candidatos, nil
}


//obtener la lista de votantes
func (s *VotacionContract) ObtenerVotantes(ctx contractapi.TransactionContextInterface, uiVotacion string) (*Voters, error) {
	votersJSON, err := ctx.GetStub().GetState("VOTANTES_" + uiVotacion)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo los votantes: %v", err)
	}
	if votersJSON == nil {
		return &Voters{
			Votantes: []string{},
		}, nil
	}
	var voters Voters
	err = json.Unmarshal(votersJSON, &voters)
	if err != nil {
		return nil, fmt.Errorf("error deserializando la lista de votantes: %v", err)
	}

	return &voters, nil
}

//Obtener una votación
func (s *VotacionContract) ObtenerVotacion(ctx contractapi.TransactionContextInterface, ui string) (*Votacion, error) {
	votacionJSON, err := ctx.GetStub().GetState(ui)
	if err != nil {
		return nil, fmt.Errorf("no se pudo recuperar la votación: %v", err)
	}
	if votacionJSON == nil {
		return nil, fmt.Errorf("votación con UI %s no encontrada", ui)
	}

	var votacion Votacion
	err = json.Unmarshal(votacionJSON, &votacion)
	if err != nil {
		return nil, fmt.Errorf("error deserializando la votación: %v", err)
	}

	return &votacion, nil
}


func main() {
	chaincode, err := contractapi.NewChaincode(&VotacionContract{})
	if err != nil {
		fmt.Printf("Error creando el chaincode: %v", err)
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error iniciando el chaincode: %v", err)
	}
}