package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Define data structures for business entities

type Admin struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Anchor struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Participant struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Invoice struct {
	ID          string  `json:"id"`
	Participant string  `json:"participant"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"` // Paid or unpaid
}

type Program struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Rate         float64  `json:"rate"`
	Participants []string `json:"participants"`
}

type DiscountProposal struct {
	ID             string  `json:"id"`
	Anchor         string  `json:"anchor"`
	Participant    string  `json:"participant"`
	Amount         float64 `json:"amount"`
	Accepted       bool    `json:"accepted"`
	Rejected       bool    `json:"rejected"`
	Modified       bool    `json:"modified"`
	ProposalType   string  `json:"proposalType"`   // Top to down or down to top
	ProposalStatus string  `json:"proposalStatus"` // Pending, accepted, rejected
}

// SmartContract provides functions for managing the SCM system
type SmartContract struct {
	contractapi.Contract
}

// Init initializes the chaincode
func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	return nil
}

// Admin management functions

func (s *SmartContract) AddAdmin(ctx contractapi.TransactionContextInterface, id, username, password string) error {
	admin := Admin{
		ID:       id,
		Username: username,
		Password: password,
	}

	adminJSON, err := json.Marshal(admin)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, adminJSON)
}

func (s *SmartContract) GetAdmin(ctx contractapi.TransactionContextInterface, id string) (*Admin, error) {
	adminJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}
	if adminJSON == nil {
		return nil, fmt.Errorf("admin with ID %s does not exist", id)
	}

	var admin Admin
	err = json.Unmarshal(adminJSON, &admin)
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

// Anchor management functions

func (s *SmartContract) AddAnchor(ctx contractapi.TransactionContextInterface, id, username, password string) error {
	anchor := Anchor{
		ID:       id,
		Username: username,
		Password: password,
	}

	anchorJSON, err := json.Marshal(anchor)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, anchorJSON)
}

func (s *SmartContract) GetAnchor(ctx contractapi.TransactionContextInterface, id string) (*Anchor, error) {
	anchorJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}
	if anchorJSON == nil {
		return nil, fmt.Errorf("anchor with ID %s does not exist", id)
	}

	var anchor Anchor
	err = json.Unmarshal(anchorJSON, &anchor)
	if err != nil {
		return nil, err
	}

	return &anchor, nil
}

// Participant management functions

func (s *SmartContract) AddParticipant(ctx contractapi.TransactionContextInterface, id, username, password string) error {
	participant := Participant{
		ID:       id,
		Username: username,
		Password: password,
	}

	participantJSON, err := json.Marshal(participant)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, participantJSON)
}

func (s *SmartContract) GetParticipant(ctx contractapi.TransactionContextInterface, id string) (*Participant, error) {
	participantJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}
	if participantJSON == nil {
		return nil, fmt.Errorf("participant with ID %s does not exist", id)
	}

	var participant Participant
	err = json.Unmarshal(participantJSON, &participant)
	if err != nil {
		return nil, err
	}

	return &participant, nil
}

// Invoice management functions

func (s *SmartContract) AddInvoice(ctx contractapi.TransactionContextInterface, id, participantID string, amount float64) error {
	invoice := Invoice{
		ID:          id,
		Participant: participantID,
		Amount:      amount,
		Status:      "unpaid",
	}

	invoiceJSON, err := json.Marshal(invoice)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, invoiceJSON)
}

func (s *SmartContract) GetInvoice(ctx contractapi.TransactionContextInterface, id string) (*Invoice, error) {
	invoiceJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}
	if invoiceJSON == nil {
		return nil, fmt.Errorf("invoice with ID %s does not exist", id)
	}

	var invoice Invoice
	err = json.Unmarshal(invoiceJSON, &invoice)
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}

// Program management functions

func (s *SmartContract) CreateProgram(ctx contractapi.TransactionContextInterface, id, name, description string, rate float64, participants []string) error {
	program := Program{
		ID:           id,
		Name:         name,
		Description:  description,
		Rate:         rate,
		Participants: participants,
	}

	programJSON, err := json.Marshal(program)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, programJSON)
}

func (s *SmartContract) GetProgram(ctx contractapi.TransactionContextInterface, id string) (*Program, error) {
	programJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}
	if programJSON == nil {
		return nil, fmt.Errorf("program with ID %s does not exist", id)
	}

	var program Program
	err = json.Unmarshal(programJSON, &program)
	if err != nil {
		return nil, err
	}

	return &program, nil
}

// Discounting functions

func (s *SmartContract) CreateDiscountProposal(ctx contractapi.TransactionContextInterface, id, anchorID, participantID string, amount float64, proposalType string) error {
	proposal := DiscountProposal{
		ID:             id,
		Anchor:         anchorID,
		Participant:    participantID,
		Amount:         amount,
		Accepted:       false,
		Rejected:       false,
		Modified:       false,
		ProposalType:   proposalType,
		ProposalStatus: "pending",
	}

	proposalJSON, err := json.Marshal(proposal)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, proposalJSON)
}

func (s *SmartContract) GetDiscountProposal(ctx contractapi.TransactionContextInterface, id string) (*DiscountProposal, error) {
	proposalJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, err
	}
	if proposalJSON == nil {
		return nil, fmt.Errorf("discount proposal with ID %s does not exist", id)
	}

	var proposal DiscountProposal
	err = json.Unmarshal(proposalJSON, &proposal)
	if err != nil {
		return nil, err
	}

	return &proposal, nil
}

// Flow executed to initiate discounting

func (s *SmartContract) InitiateDiscountingFlow(ctx contractapi.TransactionContextInterface, adminID, anchorID, participantID string) error {
	admin, err := s.GetAdmin(ctx, adminID)
	if err != nil {
		return err
	}

	anchor, err := s.GetAnchor(ctx, anchorID)
	if err != nil {
		return err
	}

	participant, err := s.GetParticipant(ctx, participantID)
	if err != nil {
		return err
	}

	// Perform the steps in the flow as described in the requirements

	return nil
}
