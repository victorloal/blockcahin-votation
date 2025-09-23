package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// Estructura para almacenar la información de cada candidato
type Candidate struct {
	Name     string `json:"name"`
	Votes    int    `json:"votes"`
	IsActive bool   `json:"isActive"`
}

// Estructura para almacenar la información de la votación
type Vote struct {
	Voter  string `json:"voter"`
	CandidateName string `json:"candidate"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	// Inicialización vacía, por ahora no agregamos datos
	return nil
}

func (s *SmartContract) PostulateCandidate(ctx contractapi.TransactionContextInterface, name string) error {
	// Verifica si el candidato ya existe
	candidateAsBytes, err := ctx.GetStub().GetState(name)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if candidateAsBytes != nil {
		return fmt.Errorf("the candidate %s already exists", name)
	}

	// Si no existe, lo agregamos
	candidate := Candidate{
		Name:     name,
		Votes:    0,  // Inicialmente sin votos
		IsActive: true,
	}
	candidateJSON, err := json.Marshal(candidate)
	if err != nil {
		return fmt.Errorf("failed to marshal candidate: %v", err)
	}

	err = ctx.GetStub().PutState(name, candidateJSON)
	if err != nil {
		return fmt.Errorf("failed to put candidate into world state: %v", err)
	}

	return nil
}

func (s *SmartContract) VoteForCandidate(ctx contractapi.TransactionContextInterface, voter string, candidateName string) error {
	// Verificar si el candidato existe
	candidateAsBytes, err := ctx.GetStub().GetState(candidateName)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if candidateAsBytes == nil {
		return fmt.Errorf("the candidate %s does not exist", candidateName)
	}

	// Verificar que el candidato esté activo
	var candidate Candidate
	err = json.Unmarshal(candidateAsBytes, &candidate)
	if err != nil {
		return fmt.Errorf("failed to unmarshal candidate: %v", err)
	}
	if !candidate.IsActive {
		return fmt.Errorf("the candidate %s is no longer active", candidateName)
	}

	// Registrar el voto
	vote := Vote{
		Voter:        voter,
		CandidateName: candidateName,
	}

	// Almacenamos el voto en la blockchain
	voteJSON, err := json.Marshal(vote)
	if err != nil {
		return fmt.Errorf("failed to marshal vote: %v", err)
	}

	err = ctx.GetStub().PutState(fmt.Sprintf("vote-%s-%s", voter, candidateName), voteJSON)
	if err != nil {
		return fmt.Errorf("failed to put vote into world state: %v", err)
	}

	// Incrementar los votos del candidato
	candidate.Votes += 1
	candidateJSON, err := json.Marshal(candidate)
	if err != nil {
		return fmt.Errorf("failed to marshal updated candidate: %v", err)
	}

	err = ctx.GetStub().PutState(candidateName, candidateJSON)
	if err != nil {
		return fmt.Errorf("failed to update candidate in world state: %v", err)
	}

	return nil
}

func (s *SmartContract) GetCandidateVotes(ctx contractapi.TransactionContextInterface, candidateName string) (int, error) {
	candidateAsBytes, err := ctx.GetStub().GetState(candidateName)
	if err != nil {
		return 0, fmt.Errorf("failed to read from world state: %v", err)
	}
	if candidateAsBytes == nil {
		return 0, fmt.Errorf("the candidate %s does not exist", candidateName)
	}

	var candidate Candidate
	err = json.Unmarshal(candidateAsBytes, &candidate)
	if err != nil {
		return 0, fmt.Errorf("failed to unmarshal candidate: %v", err)
	}

	return candidate.Votes, nil
}
func (s *SmartContract) GetAllCandidates(ctx contractapi.TransactionContextInterface) ([]Candidate, error) {
	// Obtenemos todos los candidatos
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// Iteramos sobre los resultados
	var candidates []Candidate
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var candidate Candidate
		err = json.Unmarshal(queryResponse.Value, &candidate)
		if err != nil {
			return nil, err
		}

		candidates = append(candidates, candidate)
	}

	return candidates, nil
}
func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting chaincode: %s", err.Error())
	}
}
