package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math"

	// "regexp"
	"strconv"
	// "strings"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	
	"github.com/tuneinsight/lattigo/v6/core/rlwe"
	"github.com/tuneinsight/lattigo/v6/schemes/bgv"
)

// Estructura para un Candidato
type Candidato struct {
	UI       string `json:"UI"`
	ID       string `json:"ID"`
	Posicion int  `json:"Posicion"`
}
type CandidatosVotacion struct {
	UIVotacion string      `json:"UIVotacion"`
	Candidatos []Candidato `json:"Candidatos"`
}

// Estructura para un Voto
type Voto struct {
	Fecha   string   `json:"Fecha"`
	Voto    []string `json:"Voto"`
	Votante string   `json:"Votante"`
}

type ResultadoEncriptado struct {
	UIVotacion string   `json:"UIVotacion"`
	Resultado  []string `json:"Resultado"`
}

type Resultado struct {
	UIVotacion string   `json:"UIVotacion"`
	Resultado  []int `json:"Resultado"`
	TotalVotos int `json:"TotalVotos"`
}

type PruebaVotacion struct {
	Candidatos          []Candidato `json:"Candidatos"`
	VotosPueba          []Voto      `json:"VotosPrueba"`
	ResultadoEncriptado []string    `json:"ResultadoEncriptado"`
	Resultado           []int    `json:"Resultado"`
}

// Estructura para una Votación
type Votacion struct {
	UI        string `json:"UI"`
	Nombre    string `json:"Nombre"`
	Activo    bool   `json:"Activo"`
	Inicio    string `json:"Inicio"`
	Fin       string `json:"Fin"`
	Homorfica bool   `json:"Homorfica"`
}

// Smart Contract
type VotacionContract struct {
	contractapi.Contract
}

// Función Nueva Votación
func (s *VotacionContract) NuevaVotacion(ctx contractapi.TransactionContextInterface, ui string, nombre string, inicio string, fin string, homorfica bool) error {

	// Validar UI (debe ser una cadena alfanumérica de 10 caracteres)
	// Validar Nombre (debe ser una cadena no vacía)
	// Validar Inicio y Fin (deben ser fechas en formato YYYY-MM-DD y Fin debe ser posterior a Inicio, ambas en el futuro)
	// Validar Homorfica (debe ser un valor booleano)

	// Crear una nueva votación
	votacion := Votacion{
		UI:        ui,
		Nombre:    nombre,
		Activo:    false,
		Inicio:    inicio,
		Fin:       fin,
		Homorfica: homorfica,
	}
	
	// Serializar la votación a JSON
	votacionJSON, err := json.Marshal(votacion)
	if err != nil {
		return fmt.Errorf("error serializando la votación: %v", err)
	}

	// Guardar la votación en el blockchain
	return ctx.GetStub().PutState(ui, votacionJSON)

}

// Funcion para crear la lista de candidatos vacia
func (s *VotacionContract) CrearListaCandidatos(ctx contractapi.TransactionContextInterface, uiVotacion string) error {
	// Verificar que la votación exista
	votacion, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return err
	}

	// Crear la lista de candidatos vacía
	candidatos := CandidatosVotacion{
		UIVotacion: votacion.UI,
		Candidatos: []Candidato{},
	}

	// Serializar la lista de candidatos a JSON
	candidatosJSON, err := json.Marshal(candidatos)
	if err != nil {
		return fmt.Errorf("error serializando la lista de candidatos: %v", err)
	}

	// Guardar la lista de candidatos en el blockchain
	return ctx.GetStub().PutState("CANDIDATOS_"+uiVotacion, candidatosJSON)
}

// Función para agregar un candidato
func (s *VotacionContract) AgregarCandidato(ctx contractapi.TransactionContextInterface, uiVotacion string, uiCandidato string, id string) error {
	// Verificar si la votación está activa
	votacion, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return err
	}
	if votacion.Activo {
		return errors.New("la votación está activa")
	}

	// Validar UI del candidato (debe ser una cadena alfanumérica de 10 caracteres)
	// Validar ID del candidato (debe ser una cadena no vacía)
	// Validar Posición (debe ser un entero positivo y único dentro de la votación)

	// Verificar que la posición no esté ya ocupada por otro candidato en la misma votación

	candidatosJSON, err := ctx.GetStub().GetState("CANDIDATOS_" + uiVotacion)
	if err != nil {
		return fmt.Errorf("error recuperando la lista de candidatos: %v", err)
	}
	if candidatosJSON == nil {
		return fmt.Errorf("la lista de candidatos para la votación %s no existe", uiVotacion)
	}

	var candidatosVotacion CandidatosVotacion
	err = json.Unmarshal(candidatosJSON, &candidatosVotacion)
	if err != nil {
		return fmt.Errorf("error deserializando la lista de candidatos: %v", err)
	}

	// Verificar que la posición no esté ya ocupada y que el UI del candidato no exista
	for _, candidato := range candidatosVotacion.Candidatos {
		if candidato.UI == uiCandidato {
			return fmt.Errorf("el candidato con UI %s ya existe en esta votación", uiCandidato)
		}
	}
	posicion := len(candidatosVotacion.Candidatos)

	// Actualizar la lista de candidatos
	candidatosVotacion.Candidatos = append(candidatosVotacion.Candidatos, Candidato{
		UI:       uiCandidato,
		ID:       id,
		Posicion: int(posicion),
	})

	// Serializar la lista de candidatos actualizada
	candidatosActualizadosJSON, err := json.Marshal(candidatosVotacion)
	if err != nil {
		return fmt.Errorf("error serializando la lista de candidatos actualizada: %v", err)
	}

	// Guardar la lista de candidatos actualizada en el blockchain
	return ctx.GetStub().PutState("CANDIDATOS_"+uiVotacion, candidatosActualizadosJSON)
}

// Funcion para probar la votacion
func (s *VotacionContract) PruebaVotacion(ctx contractapi.TransactionContextInterface, uiVotacion string, votos []int, publickey string, privada string) error {
	// Verificar que la votación exista
	votacion, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return err
	}
	// Verificar que la votación no esté activa
	if votacion.Activo {
		return errors.New("la votación está activa")
	}

	// Obtener la lista de candidatos desde el blockchain
	candidatosJSON, err := ctx.GetStub().GetState("CANDIDATOS_" + uiVotacion)
	if err != nil {
		return fmt.Errorf("error recuperando la lista de candidatos: %v", err)
	}
	if candidatosJSON == nil {
		return fmt.Errorf("la lista de candidatos para la votación %s no existe", uiVotacion)
	}

	var candidatosVotacion CandidatosVotacion
	err = json.Unmarshal(candidatosJSON, &candidatosVotacion)
	if err != nil {
		return fmt.Errorf("error deserializando la lista de candidatos: %v", err)
	}

	// Verificar que haya al menos un candidato
	if len(candidatosVotacion.Candidatos) == 0 {
		return fmt.Errorf("no hay candidatos registrados para la votación %s", uiVotacion)
	}

	// optener el nombre de quien realiza la prueba de votacion
	clienteID, err := ctx.GetClientIdentity().GetID()
	if err != nil {
		return fmt.Errorf("error obteniendo el ID del cliente: %v", err)
	}

	// Crear votos de prueba
	//Verificar si es una votacion homorfica
	var candidatos = len(candidatosVotacion.Candidatos)

	var votosPrueba []Voto
	var resultadoEncriptado []string
	var resultado []int
	if !votacion.Homorfica {
		votosPrueba, err = votoPrueba(votos, clienteID, candidatos)
		if err != nil {
			return fmt.Errorf("error creando votos de prueba: %v", err)
		}
		resultado, err = ContarVotos(votosPrueba, candidatos)
		if err != nil {
			return fmt.Errorf("error contando votos de prueba: %v", err)
		}
	} else {
		// Implementar creación de votos de prueba homomórficos
		votosPrueba, err = votoPruebaHomorfica(votos, clienteID, candidatos, publickey)
		if err != nil {
			return fmt.Errorf("error creando votos de prueba homorfica: %v", err)
		}
		resultadoEncriptado, err = calcularResultadoHomorfico(votosPrueba, publickey)
		if err != nil {
			return fmt.Errorf("error calculando el resultado homorfico: %v", err)
		}
		resultado, err = desencriptarResultadoHomorfico(resultadoEncriptado, privada)
		if err != nil {
			return fmt.Errorf("error desencriptando el resultado homorfico: %v", err)
		}

	}

	// Crear la estructura de prueba de votación
	prueba := PruebaVotacion{
		Candidatos:          candidatosVotacion.Candidatos,
		VotosPueba:          votosPrueba,
		ResultadoEncriptado: resultadoEncriptado,
		Resultado:           resultado,
	}

	// Serializar la prueba de votación a JSON
	pruebaJSON, err := json.Marshal(prueba)
	if err != nil {
		return fmt.Errorf("error serializando la prueba de votación: %v", err)
	}
	// Guardar la prueba de votación en el blockchain
	if err := ctx.GetStub().PutState("PRUEBA_"+uiVotacion, pruebaJSON); err != nil {
		return fmt.Errorf("error guardando la prueba de votación: %v", err)
	}
	return nil
}

func votoPrueba(votos []int, clientId string, c int) ([]Voto, error) {
	var v []Voto
	fechaActual := time.Now().Format("2006-01-02 15:04:05")

	// Crear una plantilla de voto en blanco como []string para coincidir con Voto.Voto
	votoBlanco := make([]string, c)
	for i := 0; i < c; i++ {
		votoBlanco[i] = "0"
	}

	for _, voto := range votos {
		idx := int(voto)
		if idx < 0 || idx >= c {
			return nil, fmt.Errorf("índice de voto fuera de rango: %d", idx)
		}

		// Crear una copia independiente para cada voto para evitar mutaciones compartidas
		votoCopy := make([]string, c)
		copy(votoCopy, votoBlanco)
		votoCopy[idx] = "1"

		v = append(v, Voto{
			Fecha:   fechaActual,
			Voto:    votoCopy,
			Votante: clientId,
		})
	}

	return v, nil
}

func desencriptarResultadoHomorfico(resultadoEncriptado []string, priv string) ([]int, error) {
	// var resultadosDesencriptados []int
	// privateKeyBytes, err := base64.StdEncoding.DecodeString(priv)
	// if err != nil {
	// 	return nil, fmt.Errorf("error decodificando la clave privada: %v", err)
	// }

	// // Deserializar la clave privada
	// privateKey := new(crypto.PrivateKey)
	// err = json.Unmarshal(privateKeyBytes, privateKey)
	// if err != nil {
	// 	return nil, fmt.Errorf("error deserializando la clave privada: %v", err)
	// }

	// // Desencriptar cada uno de los resultados encriptados
	// for _, cStr := range resultadoEncriptado {
	// 	cipher := new(big.Int)
	// 	if _, ok := cipher.SetString(cStr, 10); !ok {
	// 		return nil, fmt.Errorf("ciphertext inválido (no es entero base 10)")
	// 	}

	// 	// Desencriptar el ciphertext
	// 	plain, err := privateKey.Decrypt(cipher)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("error desencriptando: %v", err)
	// 	}

	// 	// Convertir el mensaje desencriptado a un número
	// 	m := new(big.Int).SetBytes(plain)
	// 	resultadosDesencriptados = append(resultadosDesencriptados, int(m.Int64()))
	// }

	// return resultadosDesencriptados, nil
	return nil, nil
}

func calcularResultadoHomorfico(votosPrueba []Voto, pub string) ([]string, error) {
	// // Inicializamos la clave pública
	// publicKeyBytes, err := base64.StdEncoding.DecodeString(pub)
	// if err != nil {
	// 	return nil, fmt.Errorf("error decodificando la clave publica: %v", err)
	// }

	// publicKey := new(crypto.PublicKey)
	// err = json.Unmarshal(publicKeyBytes, publicKey)
	// if err != nil {
	// 	return nil, fmt.Errorf("error deserializando la clave publica: %v", err)
	// }

	// // Inicializamos el acumulador
	// cols := len(votosPrueba[0].Voto)
	// acc := make([]*big.Int, cols)
	// for i := 0; i < cols; i++ {
	// 	acc[i] = big.NewInt(0) // Usamos big.Int en lugar de []byte
	// }

	// // Sumamos los votos cifrados
	// for _, voto := range votosPrueba {
	// 	for j, cStr := range voto.Voto {
	// 		cipher := new(big.Int)
	// 		if _, ok := cipher.SetString(cStr, 10); !ok {
	// 			return nil, fmt.Errorf("ciphertext inválido (no es entero base 10)")
	// 		}

	// 		// Sumamos los votos homomórficamente
	// 		acc[j].Add(acc[j], cipher)
	// 	}
	// }

	// // Convertimos el resultado acumulado a string
	// resultadosEncriptados := make([]string, cols)
	// for i, res := range acc {
	// 	resultadosEncriptados[i] = res.String() // Convertimos a string
	// }

	// return resultadosEncriptados, nil
	return nil, nil
}


func ContarVotos(votosPrueba []Voto, c int) ([]int, error) {

	resultado := make([]int, c)
	for i := 0; i < c; i++ {
		resultado[i] = 0
	}

	for _, voto := range votosPrueba {

		for idx, v := range voto.Voto {
			// Si es 1 incrementar resultado en la posicion correspondiente
			if v == "1" {
				// Incrementar el resultado en la posición correspondiente
				if idx < len(resultado) {
					resultado[idx]++
				} else {
					return nil, fmt.Errorf("índice de voto fuera de rango: %d", idx)
				}
			}
		}

	}
	return resultado, nil
}


func votoPruebaHomorfica(votos []int, clienteID string, candidatos int, pubKeyBase64 string) ([]Voto, error) {
	var votosCifrados []Voto
	fechaActual := time.Now().Format("2006-01-02 15:04:05")

	// 1. Decodificar la clave pública desde base64
	publicKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyBase64)
	if err != nil {
		return nil, fmt.Errorf("error decodificando la clave pública: %v", err)
	}

	// 2. Configurar parámetros (deben ser los mismos que usaste para generar las claves)
	params, err := bgv.NewParametersFromLiteral(bgv.ParametersLiteral{
		LogN: 1,
		Q:    []uint64{0x10000000001, 0x10000000001},
		P:    []uint64{0x10000000001},
	})
	if err != nil {
		// Si falla, usar parámetros de ejemplo
		return nil, fmt.Errorf("error creando parámetros BGV: %v", err)
	}

	// 3. Crear instancia del esquema BGV
	encoder := bgv.NewEncoder(params)
	encryptor := rlwe.NewEncryptor(params, &rlwe.PublicKey{})

	// 4. Deserializar la clave pública
	publicKey := rlwe.NewPublicKey(params)
	if err := publicKey.UnmarshalBinary(publicKeyBytes); err != nil {
		return nil, fmt.Errorf("error deserializando la clave pública: %v", err)
	}

	// Actualizar el encryptor con la clave pública real
	encryptor = rlwe.NewEncryptor(params, publicKey)

	// 5. Procesar cada voto
	for _, voto := range votos {
		// Validar que el voto esté en rango
		if voto < 0 || voto >= candidatos {
			return nil, fmt.Errorf("índice de voto fuera de rango: %d (candidatos: %d)", voto, candidatos)
		}

		// Crear vector de voto one-hot
		votoVector := make([]uint64, candidatos)
		for i := 0; i < candidatos; i++ {
			if i == voto {
				votoVector[i] = 1
			} else {
				votoVector[i] = 0
			}
		}

		// Cifrar cada posición del vector
		cifradoVoto := make([]string, candidatos)
		for i := 0; i < candidatos; i++ {
			// Codificar el valor (0 o 1) en un polinomio
			plaintext := bgv.NewPlaintext(params, params.MaxLevel())
			encoder.Encode([]uint64{votoVector[i]}, plaintext)

			// Cifrar el plaintext
			ciphertext, err := encryptor.EncryptNew(plaintext)
			if err != nil {
				return nil, fmt.Errorf("error cifrando voto para candidato %d: %v", i, err)
			}

			// Serializar ciphertext a bytes
			ciphertextBytes, err := ciphertext.MarshalBinary()
			if err != nil {
				return nil, fmt.Errorf("error serializando ciphertext: %v", err)
			}

			// Convertir a base64 para almacenamiento
			cifradoVoto[i] = base64.StdEncoding.EncodeToString(ciphertextBytes)
		}

		// Crear el objeto Voto
		votosCifrados = append(votosCifrados, Voto{
			Fecha:   fechaActual,
			Voto:    cifradoVoto,
			Votante: clienteID,
		})
	}

	return votosCifrados, nil
}


func (s *VotacionContract) ContarVotos(ctx contractapi.TransactionContextInterface, uiVotacion string) ([]int, error) {
	// Obtener la votación
	votacion, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return nil, err
	}

	// Comprobar que la votación ha finalizado
	fechaFin, _ := time.Parse("2006-01-02", votacion.Fin)
	if time.Now().Before(fechaFin) {
		return nil, errors.New("la votación aún no ha finalizado")
	}

	// Contar votos para cada candidato (dependiendo de la encriptación homomórfica)
	var votos []int
	if votacion.Homorfica {
		// Implementar conteo homomórfico (encriptación no desencriptada)
	} else {
		// Implementar conteo de votos normal
	}

	return votos, nil
}

// Función para votar
func (s *VotacionContract) Votar(ctx contractapi.TransactionContextInterface, uiVotacion string, votante string, voto []string) error {
	// Verificar si la votación está activa
	votacion, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return err
	}
	if !votacion.Activo {
		return errors.New("la votación no está activa")
	}

	// Verificar que el votante no haya votado antes
	votoExistente, err := s.ObtenerVoto(ctx,uiVotacion, votante)
	if err == nil && votoExistente != nil {
		return errors.New("el votante ya ha emitido su voto")
	}

	// Crear y registrar el voto
	votoStruct := Voto{
		Fecha:   time.Now().Format("2006-01-02 15:04:05"),
		Voto:    voto,
		Votante: votante,
	}

	// Serializar el voto
	votoJSON, err := json.Marshal(votoStruct)
	if err != nil {
		return fmt.Errorf("error serializando el voto: %v", err)
	}

	// Guardar el voto en el blockchain
	return ctx.GetStub().PutState(uiVotacion+ "_" +votante, votoJSON)
}

//Activar votacion
func (s *VotacionContract) ActivarVotacion(ctx contractapi.TransactionContextInterface, ui string) error {
	// Obtener la votación
	votacion, err := s.ObtenerVotacion(ctx, ui)
	if err != nil {
		return err
	}

	// Verificar que la votación no esté ya activa
	if votacion.Activo {
		return errors.New("la votación ya está activa")
	}

	// Verificar que la fecha actual esté dentro del rango de la votación
	fechaInicio, err := time.Parse("2006-01-02", votacion.Inicio)
	if err != nil {
		return fmt.Errorf("formato de fecha de inicio inválido: %v", err)
	}
	fechaFin, err := time.Parse("2006-01-02", votacion.Fin)
	if err != nil {
		return fmt.Errorf("formato de fecha de fin inválido: %v", err)
	}
	fechaActual := time.Now()
	if fechaActual.Before(fechaInicio) || fechaActual.After(fechaFin) {
		return errors.New("la votación no puede ser activada fuera del rango de fechas")
	}

	// Activar la votación
	votacion.Activo = true

	// Serializar la votación actualizada
	votacionJSON, err := json.Marshal(votacion)
	if err != nil {
		return fmt.Errorf("error serializando la votación: %v", err)
	}

	// Guardar la votación actualizada en el blockchain
	return ctx.GetStub().PutState(ui, votacionJSON)
}

//Terminar votacion
func (s *VotacionContract) TerminarVotacion(ctx contractapi.TransactionContextInterface, ui string) error {
	// Obtener la votación
	votacion, err := s.ObtenerVotacion(ctx, ui)
	if err != nil {
		return err
	}

	// Verificar que la votación esté activa
	if !votacion.Activo {
		return errors.New("la votación no está activa")
	}

	// Verificar que la fecha actual sea posterior a la fecha de fin
	fechaFin, err := time.Parse("2006-01-02", votacion.Fin)
	if err != nil {
		return fmt.Errorf("formato de fecha de fin inválido: %v", err)
	}

	fechaActual := time.Now()
	if fechaActual.Before(fechaFin) {
		return errors.New("la votación no puede ser terminada antes de la fecha de fin")
	}
	// Terminar la votación
	votacion.Activo = false
	// Serializar la votación actualizada
	votacionJSON, err := json.Marshal(votacion)
	if err != nil {
		return fmt.Errorf("error serializando la votación: %v", err)
	}
	// Guardar la votación actualizada en el blockchain
	return ctx.GetStub().PutState(ui, votacionJSON)
}

//Obtener prueba de votacion
func (s *VotacionContract) ObtenerPruebaVotacion(ctx contractapi.TransactionContextInterface, ui string) (*PruebaVotacion, error) {
	pruebaJSON, err := ctx.GetStub().GetState("PRUEBA_" + ui)
	if err != nil {
		return nil, fmt.Errorf("no se pudo recuperar la prueba de votación: %v", err)
	}
	if pruebaJSON == nil {
		return nil, fmt.Errorf("prueba de votación con UI %s no encontrada", ui)
	}
	var prueba PruebaVotacion
	err = json.Unmarshal(pruebaJSON, &prueba)
	if err != nil {
		return nil, fmt.Errorf("error deserializando la prueba de votación: %v", err)
	}
	return &prueba, nil
}

//Obtener lista de votos
func (s *VotacionContract) ObtenerListaVotos(ctx contractapi.TransactionContextInterface, uiVotacion string) ([]Voto, error) {
	// Obtener la votación para verificar que existe
	_, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return nil, err
	}
	// Obtener todos los votos asociados a la votación
	resultsIterator, err := ctx.GetStub().GetStateByRange(uiVotacion+"_", uiVotacion+"_"+ strconv.Itoa(math.MaxInt)) 
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

//Calcular Resultados
func (s *VotacionContract) CalcularResultados(ctx contractapi.TransactionContextInterface, uiVotacion string, privada string,publickey string) error {	
	// Obtener la votación
	votacion, err := s.ObtenerVotacion(ctx, uiVotacion)
	if err != nil {
		return err
	}
	// Comprobar que la votación ha finalizado
	fechaFin, _ := time.Parse("2006-01-02", votacion.Fin)
	if time.Now().Before(fechaFin) {
		return errors.New("la votación aún no ha finalizado")
	}
	// obtener la lista de votos
	votos, err := s.ObtenerListaVotos(ctx, uiVotacion)
	if err != nil {
		return err
	}
	// obtener la lista de candidatos
	candidatos, err := s.obtenerCandidatos(ctx, uiVotacion)
	if err != nil {
		return err
	}
	// Contar votos para cada candidato (dependiendo de la encriptación homomórfica)
	// Crear votos de prueba
	//Verificar si es una votacion homorfica
	var can = len(candidatos.Candidatos)
	var resultadoEncriptado []string
	var resultado []int

	if !votacion.Homorfica {
		resultado, err = ContarVotos(votos, can)
		if err != nil {
			return fmt.Errorf("error contando votos de prueba: %v", err)
		}
	} else {
		
		resultadoEncriptado, err = calcularResultadoHomorfico(votos, publickey)
		if err != nil {
			return fmt.Errorf("error calculando el resultado homorfico: %v", err)
		}
		resultado, err = desencriptarResultadoHomorfico(resultadoEncriptado, privada)
		if err != nil {
			return fmt.Errorf("error desencriptando el resultado homorfico: %v", err)
		}

	}
	resultadoEncriptado2 := ResultadoEncriptado{
		UIVotacion: uiVotacion,
		Resultado:  resultadoEncriptado,
	}

	resultado2 := Resultado{
		UIVotacion: uiVotacion,
		Resultado:  resultado,
		TotalVotos: int(len(votos)),
	}

	// Serializar la prueba de votación a JSON
	resultadoEncriptadoJSON, err := json.Marshal(resultadoEncriptado2)
	if err != nil {
		return fmt.Errorf("error serializando la prueba de votación: %v", err)
	}
	// Guardar la prueba de votación en el blockchain
	if err := ctx.GetStub().PutState("RESULTADO_ENCRIPTADO_"+uiVotacion, resultadoEncriptadoJSON); err != nil {
		return fmt.Errorf("error guardando la prueba de votación: %v", err)
	}

	resultadoJSON, err := json.Marshal(resultado2)
	if err != nil {
		return fmt.Errorf("error serializando la prueba de votación: %v", err)
	}
	// Guardar la prueba de votación en el blockchain
	if err := ctx.GetStub().PutState("RESULTADO_"+uiVotacion, resultadoJSON); err != nil {
		return fmt.Errorf("error guardando la prueba de votación: %v", err)
	}

	return nil


}

//obtener candidatos
func (s *VotacionContract) obtenerCandidatos(ctx contractapi.TransactionContextInterface, uiVotacion string) (*CandidatosVotacion, error) {
	// Definir la clave de inicio y fin
	candidatosJSON, err := ctx.GetStub().GetState("CANDIDATOS_"+uiVotacion)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo los candidatos: %v", err)
	}
	if candidatosJSON == nil {
		return nil, fmt.Errorf("la lista de candidatos para la votación %s no existe", uiVotacion)
	}
	var candidatos CandidatosVotacion
	err = json.Unmarshal(candidatosJSON, &candidatos)
	if err != nil {
		return nil, fmt.Errorf("error deserializando la lista de candidatos: %v", err)
	}

	return &candidatos, nil
}
//Obtener resultados encriptados
func (s *VotacionContract) ObtenerResultadosEncriptados(ctx contractapi.TransactionContextInterface, uiVotacion string) (*ResultadoEncriptado, error) {
	
	ResultadosJSON, err := ctx.GetStub().GetState("RESULTADOS_ENCRIPTADOS"+uiVotacion)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo los ResultadoE: %v", err)
	}
	if ResultadosJSON == nil {
		return nil, fmt.Errorf("la lista de ResultadoE para la votación %s no existe", uiVotacion)
	}
	var Resultado ResultadoEncriptado
	err = json.Unmarshal(ResultadosJSON, &Resultado)
	if err != nil {
		return nil, fmt.Errorf("error deserializando la lista de resultadosE: %v", err)
	}

	return &Resultado, nil
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

// Obtener una votación
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

// Obtener un voto
func (s *VotacionContract) ObtenerVoto(ctx contractapi.TransactionContextInterface, uiVotacion string,votante string) (*Voto, error) {
	votoJSON, err := ctx.GetStub().GetState(uiVotacion+ "_" +votante)
	if err != nil {
		return nil, fmt.Errorf("no se pudo recuperar el voto: %v", err)
	}
	if votoJSON == nil {
		return nil, nil // El votante no ha votado
	}

	var voto Voto
	err = json.Unmarshal(votoJSON, &voto)
	if err != nil {
		return nil, fmt.Errorf("error deserializando el voto: %v", err)
	}

	return &voto, nil
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
































// ValidarDato valida diferentes tipos de datos según reglas predefinidas
// func ValidarDato(dato interface{}, tipo string, reglas map[string]interface{}) error {
// 	switch tipo {
// 	case "string":
// 		return validarString(dato, reglas)
// 	case "int":
// 		return validarInt(dato, reglas)
// 	case "fecha":
// 		return validarFecha(dato, reglas)
// 	case "bool":
// 		return validarBool(dato)
// 	default:
// 		return fmt.Errorf("tipo de dato no soportado: %s", tipo)
// 	}
// }

// validarString valida campos de tipo string
// func validarString(dato interface{}, reglas map[string]interface{}) error {
// 	str, ok := dato.(string)
// 	if !ok {
// 		return fmt.Errorf("se esperaba un string, se recibió: %T", dato)
// 	}

// 	Validar requerido
// 	if requerido, ok := reglas["requerido"].(bool); ok && requerido && strings.TrimSpace(str) == "" {
// 		return fmt.Errorf("el campo es requerido")
// 	}

// 	Validar longitud mínima
// 	if min, ok := reglas["min"].(int); ok && len(str) < min {
// 		return fmt.Errorf("longitud mínima de %d caracteres", min)
// 	}

// 	Validar longitud máxima
// 	if max, ok := reglas["max"].(int); ok && len(str) > max {
// 		return fmt.Errorf("longitud máxima de %d caracteres", max)
// 	}

// 	Validar patrón regex
// 	if patron, ok := reglas["patron"].(string); ok && patron != "" {
// 		matched, err := regexp.MatchString(patron, str)
// 		if err != nil {
// 			return fmt.Errorf("error en patrón de validación: %v", err)
// 		}
// 		if !matched {
// 			return fmt.Errorf("el formato no es válido")
// 		}
// 	}

// 	Validar opciones específicas
// 	if opciones, ok := reglas["opciones"].([]string); ok {
// 		valido := false
// 		for _, opcion := range opciones {
// 			if str == opcion {
// 				valido = true
// 				break
// 			}
// 		}
// 		if !valido {
// 			return fmt.Errorf("valor no permitido. Opciones válidas: %v", opciones)
// 		}
// 	}

// 	return nil
// }

// validarInt valida campos de tipo entero
// func validarInt(dato interface{}, reglas map[string]interface{}) error {
// 	var num int
// 	switch v := dato.(type) {
// 	case int:
// 		num = v
// 	case string:
// 		n, err := strconv.Atoi(v)
// 		if err != nil {
// 			return fmt.Errorf("se esperaba un número entero")
// 		}
// 		num = n
// 	case float64:
// 		num = int(v)
// 	default:
// 		return fmt.Errorf("se esperaba un número entero, se recibió: %T", dato)
// 	}

// 	Validar valor mínimo
// 	if min, ok := reglas["min"].(int); ok && num < min {
// 		return fmt.Errorf("valor mínimo permitido: %d", min)
// 	}

// 	Validar valor máximo
// 	if max, ok := reglas["max"].(int); ok && num > max {
// 		return fmt.Errorf("valor máximo permitido: %d", max)
// 	}

// 	Validar opciones específicas
// 	if opciones, ok := reglas["opciones"].([]int); ok {
// 		valido := false
// 		for _, opcion := range opciones {
// 			if num == opcion {
// 				valido = true
// 				break
// 			}
// 		}
// 		if !valido {
// 			return fmt.Errorf("valor no permitido. Opciones válidas: %v", opciones)
// 		}
// 	}

// 	return nil
// }

// validarFecha valida campos de tipo fecha
// func validarFecha(dato interface{}, reglas map[string]interface{}) error {
// 	var fecha time.Time
// 	var err error

// 	switch v := dato.(type) {
// 	case string:
// 		Intentar diferentes formatos de fecha
// 		formatos := []string{"2006-01-02", "02/01/2006", "2006-01-02T15:04:05", time.RFC3339}
// 		for _, formato := range formatos {
// 			fecha, err = time.Parse(formato, v)
// 			if err == nil {
// 				break
// 			}
// 		}
// 		if err != nil {
// 			return fmt.Errorf("formato de fecha inválido. Use YYYY-MM-DD")
// 		}
// 	case time.Time:
// 		fecha = v
// 	default:
// 		return fmt.Errorf("se esperaba una fecha, se recibió: %T", dato)
// 	}

// 	Validar fecha mínima
// 	if min, ok := reglas["min"].(time.Time); ok && fecha.Before(min) {
// 		return fmt.Errorf("la fecha no puede ser anterior a %s", min.Format("2006-01-02"))
// 	}

// 	Validar fecha máxima
// 	if max, ok := reglas["max"].(time.Time); ok && fecha.After(max) {
// 		return fmt.Errorf("la fecha no puede ser posterior a %s", max.Format("2006-01-02"))
// 	}

// 	Validar que no sea futura (si se especifica)
// 	if soloPasado, ok := reglas["soloPasado"].(bool); ok && soloPasado && fecha.After(time.Now()) {
// 		return fmt.Errorf("la fecha no puede ser futura")
// 	}

// 	Validar que no sea pasada (si se especifica)
// 	if soloFuturo, ok := reglas["soloFuturo"].(bool); ok && soloFuturo && fecha.Before(time.Now()) {
// 		return fmt.Errorf("la fecha no puede ser pasada")
// 	}

// 	return nil
// }

// validarBool valida campos booleanos
// func validarBool(dato interface{}) error {
// 	switch v := dato.(type) {
// 	case bool:
// 		return nil
// 	case string:
// 		if v == "true" || v == "false" || v == "1" || v == "0" {
// 			return nil
// 		}
// 		return fmt.Errorf("se esperaba un valor booleano (true/false)")
// 	default:
// 		return fmt.Errorf("se esperaba un valor booleano, se recibió: %T", dato)
// 	}
// }
