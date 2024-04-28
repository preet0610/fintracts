package chaincode

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)
const employerCollection = "employerCollection"
const employeeCollection = "employeeCollection"
const pendingContractCollection = "pendingContractCollection"
const ongoingContractCollection = "ongoingContractCollection"
const revokedContractCollection = "revokedContractCollection"
const pendingTransactionCollection = "pendingTransactionCollection"
// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
}

// Asset describes basic details of what makes up a simple asset
// Insert struct field in alphabetic order => to achieve determinism across languages
// golang keeps the order when marshal to json but doesn't order automatically
type Employee struct {
	ID             string `json:"ID"`
	Name           string `json:"Name"`
	Location       string `json:"Location"`
	DOB			   string `json:"DOB"`
	// Employers	  []Employer `json:"Employers"`
	BankID		  string `json:"BankID"`
	BankAccountID string `json:"BankAccountID"`
	Password	   string `json:"Password"`
}

type Contract struct {
	ID 		   string `json:"ID"`
	EmployeeID             string `json:"EmployeeID"`
	EmployerID     string `json:"EmployerID"`
	CTC 		  int `json:"CTC"`
	WorkingHours   int `json:"WorkingHours"`
	Designation	string `json:"Designation"`
	PaymentCycle	string `json:"PaymentCycle"`
	// Payments	  []Transaction `json:"Payments"`
}

type Transaction struct {
	PayerID             string `json:"PayerID"`
	PayeeID     string `json:"PayeeID"`
	PayerCentralBankID             string `json:"PayerCentralBankID"`
	PayeeCentralBankID     string `json:"PayeeCentralBankID"`
	AmountOriginal 		  float64 `json:"AmountOriginal"`
	AmountReceived 		  float64 `json:"AmountReceived"`
	PaymentDate   string `json:"PaymentDate"`
	TransactionId	string `json:"TransactionId"`
	ContractID	string `json:"ContractID"`
	Status	string `json:"Status"`
}

type Employer struct {
	ID             string `json:"ID"`
	Name           string `json:"Name"`
	Location       string `json:"Location"`
	Industry       string `json:"Industry"`
	// Employees	  []Employee `json:"Employees,omitempty" metadata:",optional"`
	BankID		  string `json:"BankID"`
	BankAccountID string `json:"BankAccountID"`
	Password	   string `json:"Password"`
}

// CreateEmployer issues a new Employer to the world state with given details.
func (s *SmartContract) CreateEmployer(ctx contractapi.TransactionContextInterface, id string, name string, location string, industry string, bankid string, bankaccountid string, pwd string) (string,error) {
	fmt.Printf("CreateEmployer: collection %v, ID %v", employerCollection, id)
	var employer = Employer{
		ID:             id,
		Name:           name,
		Location:       location,
		Industry:       industry,
		BankID:			bankid,
		BankAccountID:  bankaccountid,
		Password: 	 pwd,
	}
	employerJSON, err := json.Marshal(employer)
	if err != nil {
		return "", fmt.Errorf("failed to marshal employer into JSON: %v", err)
	}

	// Check if asset already exists
	assetAsBytes, err := ctx.GetStub().GetPrivateData(employerCollection, id)
	if err != nil {
		return "", fmt.Errorf("failed to get employer: %v", err)
	} else if assetAsBytes != nil {
		fmt.Println("Employer already exists: " + id)
		return "", fmt.Errorf("this employer already exists: " + id)
	}

	// Verify that the client is submitting request to peer in their organization
	// This is to ensure that a client from another org doesn't attempt to read or
	// write private data from this peer.
	err = verifyClientOrgMatchesPeerOrg(ctx)
	if err != nil {
		return "", fmt.Errorf("CreateEmployer cannot be performed: Error %v", err)
	}

	// Save asset to private data collection
	// Typical logger, logs to stdout/file in the fabric managed docker container, running this chaincode
	// Look for container name like dev-peer0.org1.example.com-{chaincodename_version}-xyz
	log.Printf("CreateEmployer Put: collection %v, ID %v", employerCollection, id)

	err = ctx.GetStub().PutPrivateData(employerCollection, id, employerJSON)
	if err != nil {
		return "", fmt.Errorf("failed to put employer into private data collecton: %v", err)
	}
	return employer.ID, nil
}
// CreateEmployer issues a new Employee to the world state with given details.
func (s *SmartContract) CreateEmployee(ctx contractapi.TransactionContextInterface, id string, name string, location string, dob string, bankid string, bankaccountid string, pwd string) error {
	var employee = Employee{
		ID:             id,
		Name:           name,
		Location:       location,
		DOB:       dob,
		BankID:			bankid,
		BankAccountID:  bankaccountid,
		Password: 	 pwd,
	}
	employeeJSON, err := json.Marshal(employee)
	if err != nil {
		return fmt.Errorf("failed to marshal employee into JSON: %v", err)
	}

	// Check if asset already exists
	assetAsBytes, err := ctx.GetStub().GetPrivateData(employeeCollection, id)
	if err != nil {
		return fmt.Errorf("failed to get employee: %v", err)
	} else if assetAsBytes != nil {
		fmt.Println("Employee already exists: " + id)
		return fmt.Errorf("this employee already exists: " + id)
	}

	// Verify that the client is submitting request to peer in their organization
	// This is to ensure that a client from another org doesn't attempt to read or
	// write private data from this peer.
	err = verifyClientOrgMatchesPeerOrg(ctx)
	if err != nil {
		return fmt.Errorf("CreateEmployee cannot be performed: Error %v", err)
	}

	// Save asset to private data collection
	// Typical logger, logs to stdout/file in the fabric managed docker container, running this chaincode
	// Look for container name like dev-peer0.org1.example.com-{chaincodename_version}-xyz
	log.Printf("CreateEmployee Put: collection %v, ID %v", employeeCollection, id)

	err = ctx.GetStub().PutPrivateData(employeeCollection, id, employeeJSON)
	if err != nil {
		return fmt.Errorf("failed to put asset into private data collecton: %v", err)
	}
	return nil
}

func (s *SmartContract) CreateContract(ctx contractapi.TransactionContextInterface, id string, employeeid string, employerid string, ctc int, workinghrs int, designation string, paymentcycle string) error {
	assetInput := Contract{
		ID:             id,
		EmployeeID:     employeeid,
		EmployerID:     employerid,
		CTC: 			ctc,
		WorkingHours:   workinghrs,
		Designation:	designation,
		PaymentCycle:	paymentcycle,
	}
	// Check if asset already exists
	assetAsBytes, err := ctx.GetStub().GetPrivateData(employeeCollection, assetInput.EmployeeID)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	} else if assetAsBytes == nil {
		fmt.Println("Employee does not exist: " + assetInput.EmployeeID)
		return fmt.Errorf("Employee does not exist: " + assetInput.EmployeeID)
	}

	assetAsBytes, err = ctx.GetStub().GetPrivateData(employerCollection, assetInput.EmployerID)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	} else if
	assetAsBytes == nil {
		fmt.Println("Employer does not exist: " + assetInput.EmployerID)
		return fmt.Errorf("Employer does not exist: " + assetInput.EmployerID)
	}

	assetAsBytes, err = ctx.GetStub().GetPrivateData(pendingContractCollection, assetInput.ID)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if assetAsBytes != nil {
		fmt.Println("Contract already exists: " + assetInput.ID)
		return fmt.Errorf("this contract id already exists: " + assetInput.ID)
	}

	// Verify that the client is submitting request to peer in their organization
	// This is to ensure that a client from another org doesn't attempt to read or
	// write private data from this peer.
	err = verifyClientOrgMatchesPeerOrg(ctx)
	if err != nil {
		return fmt.Errorf("CreateAsset cannot be performed: Error %v", err)
	}
	asset := Contract(assetInput)
	assetJSONasBytes, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal asset into JSON: %v", err)
	}
	log.Printf("CreateContract Put: collection %v, ID %v", pendingContractCollection, assetInput.ID)
	err = ctx.GetStub().PutPrivateData(pendingContractCollection, assetInput.ID, assetJSONasBytes)
	if err != nil {
		return fmt.Errorf("failed to put asset into private data collecton: %v", err)
	}
	return nil
}

func (s *SmartContract) AgreeToContract(ctx contractapi.TransactionContextInterface, contractID string) error {
	assetAsBytes, err := ctx.GetStub().GetPrivateData(pendingContractCollection, contractID)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if assetAsBytes == nil {
		fmt.Println("Contract does not exist: " + contractID)
		return fmt.Errorf("Contract does not exist: " + contractID)
	}
	contract := Contract{}
	err = json.Unmarshal(assetAsBytes, &contract)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	assetJSONasBytes, err := json.Marshal(contract)
	if err != nil {
		return fmt.Errorf("failed to marshal asset into JSON: %v", err)
	}

	err = ctx.GetStub().DelPrivateData(pendingContractCollection, contractID)
	if err != nil {
		return fmt.Errorf("failed to delete contract from private data collection: %v", err)
	}

	log.Printf("AgreeToContract Put: collection %v, ID %v", ongoingContractCollection, contractID)
	err = ctx.GetStub().PutPrivateData(ongoingContractCollection, contractID, assetJSONasBytes)
	if err != nil {
		return fmt.Errorf("failed to put asset into private data collecton: %v", err)
	}
	return nil
}

func (s *SmartContract) RejectContract(ctx contractapi.TransactionContextInterface, contractID string) error {
	assetAsBytes, err := ctx.GetStub().GetPrivateData(pendingContractCollection, contractID)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if assetAsBytes == nil {
		fmt.Println("Contract does not exist: " + contractID)
		return fmt.Errorf("Contract does not exist: " + contractID)
	}
	contract := Contract{}
	err = json.Unmarshal(assetAsBytes, &contract)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	assetJSONasBytes, err := json.Marshal(contract)
	if err != nil {
		return fmt.Errorf("failed to marshal asset into JSON: %v", err)
	}

	err = ctx.GetStub().DelPrivateData(pendingContractCollection, contractID)
	if err != nil {
		return fmt.Errorf("failed to delete contract from private data collection: %v", err)
	}

	log.Printf("RejectContract Put: collection %v, ID %v", revokedContractCollection, contractID)
	err = ctx.GetStub().PutPrivateData(revokedContractCollection, contractID, assetJSONasBytes)
	if err != nil {
		return fmt.Errorf("failed to put asset into private data collecton: %v", err)
	}
	return nil
}

func (s *SmartContract) RevokeContract(ctx contractapi.TransactionContextInterface, contractID string) error {
	assetAsBytes, err := ctx.GetStub().GetPrivateData(ongoingContractCollection, contractID)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if assetAsBytes == nil {
		fmt.Println("Contract does not exist: " + contractID)
		return fmt.Errorf("Contract does not exist: " + contractID)
	}
	contract := Contract{}
	err = json.Unmarshal(assetAsBytes, &contract)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	assetJSONasBytes, err := json.Marshal(contract)
	if err != nil {
		return fmt.Errorf("failed to marshal asset into JSON: %v", err)
	}

	err = ctx.GetStub().DelPrivateData(ongoingContractCollection, contractID)
	if err != nil {
		return fmt.Errorf("failed to delete contract from private data collection: %v", err)
	}

	log.Printf("RevokeContract Put: collection %v, ID %v", revokedContractCollection, contractID)
	err = ctx.GetStub().PutPrivateData(revokedContractCollection, contractID, assetJSONasBytes)
	if err != nil {
		return fmt.Errorf("failed to put asset into private data collecton: %v", err)
	}
	return nil
}

func (s *SmartContract) MakePayment(ctx contractapi.TransactionContextInterface, contractID string) error {
	assetAsBytes, err := ctx.GetStub().GetPrivateData(ongoingContractCollection, contractID)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if assetAsBytes == nil {
		fmt.Println("Contract does not exist: " + contractID)
		return fmt.Errorf("Contract does not exist: " + contractID)
	}
	contract := Contract{}
	err = json.Unmarshal(assetAsBytes, &contract)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	employerAsBytes, err := ctx.GetStub().GetPrivateData(employerCollection, contract.EmployerID)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if employerAsBytes == nil {
		fmt.Println("Employer does not exist: " + contract.EmployerID)
		return fmt.Errorf("Employer does not exist: " + contract.EmployerID)
	}

	employer := Employer{}
	err = json.Unmarshal(employerAsBytes, &employer)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	employeeAsBytes, err := ctx.GetStub().GetPrivateData(employeeCollection, contract.EmployeeID)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	}
	if employeeAsBytes == nil {
		fmt.Println("Employee does not exist: " + contract.EmployeeID)
		return fmt.Errorf("Employee does not exist: " + contract.EmployeeID)
	}

	employee := Employee{}
	err = json.Unmarshal(employeeAsBytes, &employee)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}
	
	var amt = float64(contract.CTC) / 12
	switch contract.PaymentCycle {
		case "One-Time":
			amt = float64(contract.CTC)
		case "Monthly":
			amt = float64(contract.CTC) / 12
		case "Quarterly":
			amt = float64(contract.CTC) / 4
		case "Half-Yearly":
			amt = float64(contract.CTC) / 2
		case "Yearly":
			amt = float64(contract.CTC)
		default:
			amt = float64(contract.CTC) / 12
	}
	//error somewhere in this block
	transaction := Transaction{}
	transaction.TransactionId = newTansaction();
	transaction.PayerID = employer.BankAccountID+employer.BankID
	transaction.PayeeID = employee.BankAccountID+employee.BankID
	transaction.PayerCentralBankID = employer.Location+"CentralBank"
	transaction.PayeeCentralBankID = employee.Location+"CentralBank"
	transaction.AmountOriginal = amt
	transaction.AmountReceived = amt
	transaction.PaymentDate = time.Now().Format("2006-01-02")
	transaction.ContractID = contractID
	transaction.Status = "Pending from Bank to Central Bank"
	// contract.Payments = append(contract.Payments, transaction)
	assetJSONasBytes, err := json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to marshal asset into JSON: %v", err)
	}

	log.Printf("MakePayment Put: collection %v, ID %v", pendingTransactionCollection, transaction.TransactionId)
	err = ctx.GetStub().PutPrivateData(pendingTransactionCollection, transaction.TransactionId, assetJSONasBytes)
	if err != nil {
		return fmt.Errorf("failed to put asset into private data collecton: %v", err)
	}
	return nil
}

func (s *SmartContract) DeletePendingTransaction(ctx contractapi.TransactionContextInterface, transactionID string) error {
	err := ctx.GetStub().DelPrivateData(pendingTransactionCollection, transactionID)
	if err != nil {
		return fmt.Errorf("failed to delete asset: %v", err)
	}
	return nil
}

func (s *SmartContract) ReadPendingTransactionsByBankID(ctx contractapi.TransactionContextInterface, bankID string) ([]*Transaction, error) {
	
	queryString := fmt.Sprintf(`{"selector":{"$and":[{"PayerID":{"$regex":"%s$"}}]}}`, bankID)
	return s.getQueryResultForQueryStringPendingTxn(ctx, queryString)
}

func (s *SmartContract) getQueryResultForQueryStringPendingTxn(ctx contractapi.TransactionContextInterface, queryString string) ([]*Transaction, error) {

	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(pendingTransactionCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []*Transaction{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset *Transaction

		err = json.Unmarshal(response.Value, &asset)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
		}

		results = append(results, asset)
	}
	return results, nil
}

// ReadAsset reads the information from collection
func (s *SmartContract) ReadEmployer(ctx contractapi.TransactionContextInterface, assetID string) (*Employer, error) {

	log.Printf("ReadAsset: collection %v, ID %v", employerCollection, assetID)
	assetJSON, err := ctx.GetStub().GetPrivateData(employerCollection, assetID) //get the asset from chaincode state
	if err != nil {
		return nil, fmt.Errorf("failed to read asset: %v", err)
	}

	// No Asset found, return empty response
	if assetJSON == nil {
		log.Printf("%v does not exist in collection %v", assetID, employerCollection)
		return nil, nil
	}

	var asset *Employer
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return asset, nil

}

func (s *SmartContract) ReadEmployee(ctx contractapi.TransactionContextInterface, assetID string) (*Employee, error) {

	log.Printf("ReadAsset: collection %v, ID %v", employeeCollection, assetID)
	assetJSON, err := ctx.GetStub().GetPrivateData(employeeCollection, assetID) //get the asset from chaincode state
	if err != nil {
		return nil, fmt.Errorf("failed to read asset: %v", err)
	}

	// No Asset found, return empty response
	if assetJSON == nil {
		log.Printf("%v does not exist in collection %v", assetID, employeeCollection)
		return nil, nil
	}

	var asset *Employee
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return asset, nil

}
func (s *SmartContract) ReadPendingTransaction(ctx contractapi.TransactionContextInterface, assetID string) (*Transaction, error) {

	log.Printf("ReadPendingTransaction: collection %v, ID %v", pendingTransactionCollection, assetID)
	assetJSON, err := ctx.GetStub().GetPrivateData(pendingTransactionCollection, assetID) //get the asset from chaincode state
	if err != nil {
		return nil, fmt.Errorf("failed to read asset: %v", err)
	}

	// No Asset found, return empty response
	if assetJSON == nil {
		log.Printf("%v does not exist in collection %v", assetID, employeeCollection)
		return nil, nil
	}

	var asset *Transaction
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return asset, nil

}



func (s *SmartContract) QueryContractByEmployer(ctx contractapi.TransactionContextInterface, employer string) ([]*Contract, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"EmployerID\":\"%v\"}}", employer)

	queryResults, err := s.getQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}

func (s *SmartContract) QueryContractByEmployee(ctx contractapi.TransactionContextInterface, employee string) ([]*Contract, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"EmployeeID\":\"%v\"}}", employee)

	queryResults, err := s.getQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}

func (s *SmartContract) QueryPendingContractByEmployee(ctx contractapi.TransactionContextInterface, employee string) ([]*Contract, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"EmployeeID\":\"%v\"}}", employee)

	queryResults, err := s.getQueryResultForQueryStringPending(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}
func (s *SmartContract) QueryPendingContractByEmployer(ctx contractapi.TransactionContextInterface, employee string) ([]*Contract, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"EmployerID\":\"%v\"}}", employee)

	queryResults, err := s.getQueryResultForQueryStringPending(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}

func (s *SmartContract) QueryRevokedContractByEmployer(ctx contractapi.TransactionContextInterface, employee string) ([]*Contract, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"EmployerID\":\"%v\"}}", employee)

	queryResults, err := s.getQueryResultForQueryStringRevoked(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}
func (s *SmartContract) QueryRevokedContractByEmployee(ctx contractapi.TransactionContextInterface, employee string) ([]*Contract, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"EmployeeID\":\"%v\"}}", employee)

	queryResults, err := s.getQueryResultForQueryStringRevoked(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}

func verifyClientOrgMatchesPeerOrg(ctx contractapi.TransactionContextInterface) error {
	clientMSPID, err := ctx.GetClientIdentity().GetMSPID()
	if err != nil {
		return fmt.Errorf("failed getting the client's MSPID: %v", err)
	}
	peerMSPID, err := shim.GetMSPID()
	if err != nil {
		return fmt.Errorf("failed getting the peer's MSPID: %v", err)
	}

	if clientMSPID != peerMSPID {
		return fmt.Errorf("client from org %v is not authorized to read or write private data from an org %v peer", clientMSPID, peerMSPID)
	}

	return nil
}

func (s *SmartContract) getQueryResultForQueryStringPending(ctx contractapi.TransactionContextInterface, queryString string) ([]*Contract, error) {

	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(pendingContractCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []*Contract{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset *Contract

		err = json.Unmarshal(response.Value, &asset)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
		}

		results = append(results, asset)
	}
	return results, nil
}
func (s *SmartContract) getQueryResultForQueryStringRevoked(ctx contractapi.TransactionContextInterface, queryString string) ([]*Contract, error) {

	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(revokedContractCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []*Contract{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset *Contract

		err = json.Unmarshal(response.Value, &asset)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
		}

		results = append(results, asset)
	}
	return results, nil
}

// getQueryResultForQueryString executes the passed in query string.
func (s *SmartContract) getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*Contract, error) {

	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(ongoingContractCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []*Contract{}

	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset *Contract

		err = json.Unmarshal(response.Value, &asset)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
		}

		results = append(results, asset)
	}
	return results, nil
}

func newTansaction() string {
	h := fnv.New32a()
	h.Write([]byte(time.Now().String()))
	return fmt.Sprintf("%d",h.Sum32())
}
