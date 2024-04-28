package chaincode

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)
const employerCollection = "employerCollection"
const employeeCollection = "employeeCollection"
const pendingTransactionCollection = "pendingTransactionCollection"
const pendingCentralBankToForexTransactionCollection = "pendingCentralBankToForexTransactionCollection"
const pendingForexToCentralBankTransactionCollection = "pendingForexToCentralBankTransactionCollection"
const pendingCentralBanktoBankTransactionCollection = "pendingCentralBankToBankTransactionCollection"
const completedTransactionCollection = "completedTransactionCollection"
const bankCollection = "bankCollection"
const bankAccountPrivateCollection = "bankAccountPrivateCollection"
const forexBankCollection = "forexBankCollection"
const exchangeRateCollection =	"exchangeRateCollection"
const centralBankCollection = "centralBankCollection"
const bankAccountCollection = "bankAccountCollection"
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

type BankAccount struct {
	ID             string `json:"ID"`
	BankID    string `json:"BankID"`
	CentralBankID       string `json:"CentralBankID"`
	Name           string `json:"Name"`
	DOB			string `json:"DOB"`
}

type BankAccountPrivate struct {
	CustomerID             string `json:"CustomerID"`
	BankID     string `json:"BankID"`
	Balance        float64 `json:"Balance"`
	// Transactions   []Transaction `json:"Transactions"`
}

type Bank struct {
	ID             string `json:"ID"`
	Name           string `json:"Name"`
	CentralBankID       string `json:"CentralBankID"`
	Balance		float64 `json:"Balance"`
	Password	   string `json:"Password"`
	// Customers	  []BankAccount `json:"Customers"`
}

type CentralBank struct {
	ID             string `json:"ID"`
	Location	   string `json:"Location"`
	// Banks		  []Bank `json:"Banks"`
	Currency	  string `json:"Currency"`
	Balance		float64 `json:"Balance"`
	Password	   string `json:"Password"`

}

type ForexBank struct {
	ID             string `json:"ID"`
	Balance		float64 `json:"Balance"`
	Password	   string `json:"Password"`
}

type ExchangeRate struct {
	FromCurrency             string `json:"FromCurrency"`
	ToCurrency     string `json:"ToCurrency"`
	Value 		  float64 `json:"Value"`
}

func (s *SmartContract) AddExchangeValue(ctx contractapi.TransactionContextInterface, fromCurrency string, toCurrency string, value float64) error {
	exchangeRate := ExchangeRate{
		FromCurrency: fromCurrency,
		ToCurrency: toCurrency,
		Value: value,
	}
	exchangeRateJSON, err := json.Marshal(exchangeRate)
	if err != nil {
		return fmt.Errorf("failed to marshal exchangeRate: %v", err)
	}
	exchangeRateAsBytes, err := ctx.GetStub().GetPrivateData(exchangeRateCollection, fromCurrency + toCurrency)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	} else if exchangeRateAsBytes != nil {
		fmt.Println("Asset already exists: " + fromCurrency + toCurrency)
		return fmt.Errorf("this asset already exists: " + fromCurrency + toCurrency)
	}
	err = verifyClientOrgMatchesPeerOrg(ctx)
	if err != nil {
		return fmt.Errorf("function cannot be performed: Error %v", err)
	}
	err = ctx.GetStub().PutPrivateData(exchangeRateCollection, fromCurrency + toCurrency, exchangeRateJSON)
	if err != nil {
		return fmt.Errorf("failed to put asset: %v", err)
	}
	return nil
}

func (s *SmartContract) CreateBank(ctx contractapi.TransactionContextInterface, id string, name string, centralBankID string, pwd string) error {
	bank := Bank{
		ID:             id,
		Name:           name,
		CentralBankID: centralBankID,
		Password: pwd,
	}
	bankJSON, err := json.Marshal(bank)
	if err != nil {
		return fmt.Errorf("failed to marshal bank: %v", err)
	}
	bankAsBytes, err := ctx.GetStub().GetPrivateData(bankCollection, id)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	} else if bankAsBytes != nil {
		fmt.Println("Asset already exists: " + id)
		return fmt.Errorf("this asset already exists: " + id)
	}
	err = verifyClientOrgMatchesPeerOrg(ctx)
	if err != nil {
		return fmt.Errorf("CreateAsset cannot be performed: Error %v", err)
	}
	err = ctx.GetStub().PutPrivateData(bankCollection, id, bankJSON)
	if err != nil {
		return fmt.Errorf("failed to put asset: %v", err)
	}
	return nil
}

func (s *SmartContract) CreateBankAccount(ctx contractapi.TransactionContextInterface, id string, name string, dob string, bal float64, bankid string, centralbankid string) error {
	bankAccount := BankAccount{
		ID:             id,
		Name:           name,
		CentralBankID: centralbankid,
		DOB: dob,
		BankID: bankid,
	}
	bankAccountPrivateDetails := BankAccountPrivate{
		CustomerID: id,
		BankID: bankid,
		Balance: bal,
	}
	bankAccountJSON, err := json.Marshal(bankAccount)
	if err != nil {
		return fmt.Errorf("failed to marshal bankAccount: %v", err)
	}
	bankAccountAsBytes, err := ctx.GetStub().GetPrivateData(bankAccountCollection, id+bankid)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	} else if bankAccountAsBytes != nil {
		fmt.Println("Asset already exists: " + id)
		return fmt.Errorf("this asset already exists: " + id)
	}
	bankAccountPrivateJSON, err := json.Marshal(bankAccountPrivateDetails)
	if err != nil {
		return fmt.Errorf("failed to marshal bankAccountPrivateDetails: %v", err)
	}
	bankAccountPrivateAsBytes, err := ctx.GetStub().GetPrivateData(bankAccountPrivateCollection, id+bankid)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	} else if bankAccountPrivateAsBytes != nil {
		fmt.Println("Asset already exists: " + id)
		return fmt.Errorf("this asset already exists: " + id)
	}

	err = verifyClientOrgMatchesPeerOrg(ctx)
	if err != nil {
		return fmt.Errorf("CreateAsset cannot be performed: Error %v", err)
	}
	err = ctx.GetStub().PutPrivateData(bankAccountCollection, id+bankid, bankAccountJSON)
	if err != nil {
		return fmt.Errorf("failed to put asset: %v", err)
	}
	err = ctx.GetStub().PutPrivateData(bankAccountPrivateCollection, id+bankid, bankAccountPrivateJSON)
	if err != nil {
		return fmt.Errorf("failed to put asset: %v", err)
	}
	return nil
}

func (s *SmartContract) CreateCentralBank(ctx contractapi.TransactionContextInterface, id string, location string, currency string, pwd string) error {
	centralBank := CentralBank{
		ID:             id,
		Location:       location,
		Currency: currency,
		Password: pwd,
	}
	centralBankJSON, err := json.Marshal(centralBank)
	if err != nil {
		return fmt.Errorf("failed to marshal centralBank: %v", err)
	}
	centralBankAsBytes, err := ctx.GetStub().GetPrivateData(centralBankCollection, id)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	} else if centralBankAsBytes != nil {
		fmt.Println("Asset already exists: " + id)
		return fmt.Errorf("this asset already exists: " + id)
	}
	err = verifyClientOrgMatchesPeerOrg(ctx)
	if err != nil {
		return fmt.Errorf("CreateAsset cannot be performed: Error %v", err)
	}
	err = ctx.GetStub().PutPrivateData(centralBankCollection, id, centralBankJSON)
	if err != nil {
		return fmt.Errorf("failed to put asset: %v", err)
	}
	return nil
}

func (s *SmartContract) CreateForexBank(ctx contractapi.TransactionContextInterface, id string, bal float64, pwd string) error {
	forexBank := ForexBank{
		ID:             id,
		Balance: bal,
		Password: pwd,
	}
	forexBankJSON, err := json.Marshal(forexBank)
	if err != nil {
		return fmt.Errorf("failed to marshal forexBank: %v", err)
	}
	forexBankAsBytes, err := ctx.GetStub().GetPrivateData(forexBankCollection, id)
	if err != nil {
		return fmt.Errorf("failed to get asset: %v", err)
	} else if forexBankAsBytes != nil {
		fmt.Println("Asset already exists: " + id)
		return fmt.Errorf("this asset already exists: " + id)
	}
	err = verifyClientOrgMatchesPeerOrg(ctx)
	if err != nil {
		return fmt.Errorf("CreateAsset cannot be performed: Error %v", err)
	}
	err = ctx.GetStub().PutPrivateData(forexBankCollection, id, forexBankJSON)
	if err != nil {
		return fmt.Errorf("failed to put asset: %v", err)
	}
	return nil
}

func (s *SmartContract) BankToCentralBankTransaction(ctx contractapi.TransactionContextInterface, transactionID string) error {
	// Get the transaction from the pendingTransaction collection
	transactionAsBytes, err := ctx.GetStub().GetPrivateData(pendingTransactionCollection, transactionID)
	if err != nil {
		return fmt.Errorf("failed to get transaction from pendingTransaction collection: %v", err)
	}
	if transactionAsBytes == nil {
		return fmt.Errorf("transaction not found in pendingTransaction collection: %s", transactionID)
	}

	// Unmarshal the transaction
	var transaction Transaction
	err = json.Unmarshal(transactionAsBytes, &transaction)
	if err != nil {
		return fmt.Errorf("failed to unmarshal transaction: %v", err)
	}

	

	

	// Send money from the payer's bank account to the central bank
	err = sendMoneyToCentralBank(ctx, transaction.PayerID, transaction.PayerCentralBankID, transaction.AmountReceived)
	if err != nil {
		return fmt.Errorf("failed to send money to central bank: %v", err)
	}

	// Update the transaction status
	transaction.Status = "Reached Central Bank of your country"
	transactionAsBytes, err = json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction: %v", err)
	}
	
	if transaction.PayeeCentralBankID != transaction.PayerCentralBankID {
		err = ctx.GetStub().PutPrivateData(pendingCentralBankToForexTransactionCollection, transactionID, transactionAsBytes)
		if err != nil {
			return fmt.Errorf("failed to put transaction in pendingCentralBankToForexTransaction collection: %v", err)
		}
	} else {
		err = ctx.GetStub().PutPrivateData(pendingCentralBanktoBankTransactionCollection, transactionID, transactionAsBytes)
		if err != nil {
			return fmt.Errorf("failed to put transaction in pendingCentralBankToBankTransaction collection: %v", err)
		}
	}

	// Delete the transaction from the pendingTransaction collection
	err = ctx.GetStub().DelPrivateData(pendingTransactionCollection, transactionID)
	if err != nil {
		return fmt.Errorf("failed to delete transaction from pendingTransaction collection: %v", err)
	}

	return nil
}

func sendMoneyToCentralBank(ctx contractapi.TransactionContextInterface, bankAccountID string, centralBankID string, amount float64) error {
	// Get the central bank's bank account from the bank account collection
	centralbankAsBytes, err := ctx.GetStub().GetPrivateData(centralBankCollection, centralBankID)
	if err != nil {
		return fmt.Errorf("failed to get central bank account from bank account collection: %v", err)
	}
	if centralbankAsBytes == nil {
		return fmt.Errorf("central bank account not found in central bank collection: %s", centralBankID)
	}

	// Unmarshal the central bank
	var centralBank CentralBank
	err = json.Unmarshal(centralbankAsBytes, &centralBank)
	if err != nil {
		return fmt.Errorf("failed to unmarshal central bank account: %v", err)
	}
	userBankAccount, err := ctx.GetStub().GetPrivateData(bankAccountPrivateCollection, bankAccountID)
	if err != nil {
		return fmt.Errorf("failed to get bank account from bank account collection: %v", err)
	}
	if userBankAccount == nil {
		return fmt.Errorf("bank account not found in bank account collection: %s", bankAccountID)
	}

	// Unmarshal the bank account
	var bankAccount BankAccountPrivate
	err = json.Unmarshal(userBankAccount, &bankAccount)
	if err != nil {
		return fmt.Errorf("failed to unmarshal bank account: %v", err)
	}
	// Check if the payer's bank account has sufficient balance
	if bankAccount.Balance < amount {
		return fmt.Errorf("insufficient balance in payer's bank account")
	}

	// Deduct the amount from the payer's bank account
	bankAccount.Balance -= amount
	bankAccountAsBytes, err := json.Marshal(bankAccount)
	if err != nil {
		return fmt.Errorf("failed to marshal payer's bank account: %v", err)
	}

	err = ctx.GetStub().PutPrivateData(bankAccountPrivateCollection, bankAccountID, bankAccountAsBytes)
	if err != nil {
		return fmt.Errorf("failed to update payer's bank account: %v", err)
	}

	// Add the amount to the central bank's bank account
	centralBank.Balance += amount
	centralBankAccountAsBytes, err := json.Marshal(centralBank)
	if err != nil {
		return fmt.Errorf("failed to marshal central bank account: %v", err)
	}

	err = ctx.GetStub().PutPrivateData(bankAccountCollection, centralBank.ID, centralBankAccountAsBytes)
	if err != nil {
		return fmt.Errorf("failed to update central bank account: %v", err)
	}

	return nil
}

func (s *SmartContract) CentralBankToForexBankTransaction(ctx contractapi.TransactionContextInterface, transactionID string, centralBankId string) error {
	// Get the transaction from the pendingTransaction collection
	transactionAsBytes, err := ctx.GetStub().GetPrivateData(pendingCentralBankToForexTransactionCollection, transactionID)
	if err != nil {
		return fmt.Errorf("failed to get transaction from pendingCentralBankToForexTransaction collection: %v", err)
	}
	if transactionAsBytes == nil {
		return fmt.Errorf("transaction not found in pendingCentralBankToForexTransaction collection: %s", transactionID)
	}

	// Unmarshal the transaction
	var transaction Transaction
	err = json.Unmarshal(transactionAsBytes, &transaction)
	if err != nil {
		return fmt.Errorf("failed to unmarshal transaction: %v", err)
	}

	// Get the central bank's bank account from the bank account collection
	centralBankAsBytes, err := ctx.GetStub().GetPrivateData(centralBankCollection, centralBankId)
	if err != nil {
		return fmt.Errorf("failed to get central bank from central bank collection: %v", err)
	}
	if centralBankAsBytes == nil {
		return fmt.Errorf("central bank not found in central bank collection: %s", centralBankId)
	}

	// Unmarshal the central bank
	var centralBank CentralBank
	err = json.Unmarshal(centralBankAsBytes, &centralBank)
	if err != nil {
		return fmt.Errorf("failed to unmarshal central bank: %v", err)
	}

	// Get the forex bank's bank account from the bank account collection
	forexBankAsBytes, err := ctx.GetStub().GetPrivateData(forexBankCollection, "ForexBank")
	if err != nil {
		return fmt.Errorf("failed to get forex bank from forex bank collection: %v", err)
	}
	if forexBankAsBytes == nil {
		return fmt.Errorf("forex bank not found in forex bank collection: %s", "ForexBank")
	}

	// Unmarshal the forex bank
	var forexBank ForexBank
	err = json.Unmarshal(forexBankAsBytes, &forexBank)
	if err != nil {
		return fmt.Errorf("failed to unmarshal forex bank: %v", err)
	}

	// Send money from the central bank's bank account to the forex bank
	transaction.AmountReceived, err = sendMoneyToForexBank(ctx, centralBank.ID, forexBank.ID, transaction.AmountReceived)
	if err != nil {
		return fmt.Errorf("failed to send money to forex bank: %v", err)
	}
	// Update the transaction status
	transaction.Status = "Reached Forex Bank"
	transactionAsBytes, err = json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction: %v", err)
	}
	err = ctx.GetStub().PutPrivateData(pendingForexToCentralBankTransactionCollection, transactionID, transactionAsBytes)
	if err != nil {
		return fmt.Errorf("failed to put transaction in pendingForexToCentralBankTransaction collection: %v", err)
	}
	// Delete the transaction from the pendingTransaction collection
	err = ctx.GetStub().DelPrivateData(pendingCentralBankToForexTransactionCollection, transactionID)
	if err != nil {
		return fmt.Errorf("failed to delete transaction from pendingCentralBankToForexTransaction collection: %v", err)
	}
	return nil
}

func sendMoneyToForexBank(ctx contractapi.TransactionContextInterface, centralBankID string, forexBankID string, amount float64) (float64,error) {
	// Get the central bank's bank account from the bank account collection
	centralBankAsBytes, err := ctx.GetStub().GetPrivateData(centralBankCollection, centralBankID)
	if err != nil {
		return 0.0, fmt.Errorf("failed to get central bank account from central bank collection: %v", err)
	}
	if centralBankAsBytes == nil {
		return 0.0, fmt.Errorf("central bank account not found in central bank collection: %s", centralBankID)
	}

	// Unmarshal the central bank
	var centralBank CentralBank
	err = json.Unmarshal(centralBankAsBytes, &centralBank)
	if err != nil {
		return 0.0, fmt.Errorf("failed to unmarshal central bank account: %v", err)
	}

	// Get the forex bank's bank account from the bank account collection
	forexBankAsBytes, err := ctx.GetStub().GetPrivateData(forexBankCollection, forexBankID)
	if err != nil {
		return 0.0, fmt.Errorf("failed to get forex bank account from forex bank collection: %v", err)
	}
	if forexBankAsBytes == nil {
		return 0.0, fmt.Errorf("forex bank account not found in forex bank collection: %s", forexBankID)
	}

	// Unmarshal the forex bank
	var forexBank ForexBank
	err = json.Unmarshal(forexBankAsBytes, &forexBank)
	if err != nil {
		return 0.0, fmt.Errorf("failed to unmarshal forex bank account: %v", err)
	}

	// Check if the central bank's bank account has sufficient balance
	if centralBank.Balance < amount {
		return 0.0, fmt.Errorf("insufficient balance in central bank's bank account")
	}

	// Deduct the amount from the central bank's bank account
	centralBank.Balance -= amount
	centralBankAsBytes, err = json.Marshal(centralBank)
	if err != nil {
		return 0.0, fmt.Errorf("failed to marshal central bank account: %v", err)
	}

	err = ctx.GetStub().PutPrivateData(centralBankCollection, centralBank.ID, centralBankAsBytes)
	if err != nil {
		return 0.0, fmt.Errorf("failed to update central bank account: %v", err)
	}

	// Add the amount to the forex bank's bank account in USD
	exchangeRateAsBytes, err := ctx.GetStub().GetPrivateData(exchangeRateCollection, centralBank.Currency+"USD")
	if err != nil {
		return 0.0, fmt.Errorf("failed to get exchange rate from exchange rate collection: %v", err)
	}
	if exchangeRateAsBytes == nil {
		return 0.0, fmt.Errorf("exchange rate not found in exchange rate collection: %s", centralBank.Currency+"USD")
	}

	// Unmarshal the exchange rate
	var exchangeRate ExchangeRate
	err = json.Unmarshal(exchangeRateAsBytes, &exchangeRate)
	if err != nil {
		return 0.0, fmt.Errorf("failed to unmarshal exchange rate: %v", err)
	}

	amountInUSD := amount * exchangeRate.Value
	forexBank.Balance += amountInUSD
	forexBankAsBytes, err = json.Marshal(forexBank)
	if err != nil {
		return 0.0, fmt.Errorf("failed to marshal forex bank account: %v", err)
	}

	err = ctx.GetStub().PutPrivateData(forexBankCollection, forexBank.ID, forexBankAsBytes)
	if err != nil {
		return 0.0, fmt.Errorf("failed to update forex bank account: %v", err)
	}

	return amountInUSD, nil
}

func (s *SmartContract) ForexToCentralBankTransaction(ctx contractapi.TransactionContextInterface, transactionID string, forexBankId string) error {
	// Get the transaction from the pendingTransaction collection
	transactionAsBytes, err := ctx.GetStub().GetPrivateData(pendingForexToCentralBankTransactionCollection, transactionID)
	if err != nil {
		return fmt.Errorf("failed to get transaction from pendingForexToCentralBankTransaction collection: %v", err)
	}
	if transactionAsBytes == nil {
		return fmt.Errorf("transaction not found in pendingForexToCentralBankTransaction collection: %s", transactionID)
	}

	// Unmarshal the transaction
	var transaction Transaction
	err = json.Unmarshal(transactionAsBytes, &transaction)
	if err != nil {
		return fmt.Errorf("failed to unmarshal transaction: %v", err)
	}

	// Get employee's bank account from the employee collection
	employeeAsBytes, err := ctx.GetStub().GetPrivateData(employeeCollection, transaction.PayeeID)
	if err != nil {
		return fmt.Errorf("failed to get employee from employee collection: %v", err)
	}
	if employeeAsBytes == nil {
		return fmt.Errorf("employee not found in employee collection: %s", transaction.PayeeID)
	}

	// Unmarshal the employee
	var employee Employee
	err = json.Unmarshal(employeeAsBytes, &employee)
	if err != nil {
		return fmt.Errorf("failed to unmarshal employee: %v", err)
	}

	// Get the central bank's bank account from the bank account collection
	centralBankAsBytes, err := ctx.GetStub().GetPrivateData(centralBankCollection, employee.Location+"CentralBank")
	if err != nil {
		return fmt.Errorf("failed to get central bank from central bank collection: %v", err)
	}
	if centralBankAsBytes == nil {
		return fmt.Errorf("central bank not found in central bank collection: %s", employee.Location+"CentralBank")
	}

	// Unmarshal the central bank
	var centralBank CentralBank
	err = json.Unmarshal(centralBankAsBytes, &centralBank)
	if err != nil {
		return fmt.Errorf("failed to unmarshal central bank: %v", err)
	}

	// Get the forex bank's bank account from the bank account collection
	forexBankAsBytes, err := ctx.GetStub().GetPrivateData(forexBankCollection, forexBankId)
	if err != nil {
		return fmt.Errorf("failed to get forex bank from forex bank collection: %v", err)
	}
	if forexBankAsBytes == nil {
		return fmt.Errorf("forex bank not found in forex bank collection: %s", forexBankId)
	}

	// Unmarshal the forex bank
	var forexBank Bank
	err = json.Unmarshal(forexBankAsBytes, &forexBank)
	if err != nil {
		return fmt.Errorf("failed to unmarshal forex bank: %v", err)
	}

	// Send money from the forex bank's bank account to the central bank
	transaction.AmountReceived, err = sendMoneyFromForexToCentralBank(ctx, forexBank.ID, centralBank.ID, transaction.AmountReceived)
	if err != nil {
		return fmt.Errorf("failed to send money to central bank: %v", err)
	}

	// Update the transaction status
	transaction.Status = "Reached Central Bank of employee's country"
	transactionAsBytes, err = json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction: %v", err)
	}

	err = ctx.GetStub().PutPrivateData(pendingCentralBanktoBankTransactionCollection, transactionID, transactionAsBytes)
	if err != nil {
		return fmt.Errorf("failed to put transaction in pendingCentralBanktoBankTransaction collection: %v", err)
	}

	// Delete the transaction from the pendingTransaction collection
	err = ctx.GetStub().DelPrivateData(pendingForexToCentralBankTransactionCollection, transactionID)
	if err != nil {
		return fmt.Errorf("failed to delete transaction from pendingForexToCentralBankTransaction collection: %v", err)
	}

	return nil
}

func sendMoneyFromForexToCentralBank(ctx contractapi.TransactionContextInterface, forexBankID string, centralBankID string, amount float64) (float64, error) {
	// Get the central bank's bank account from the bank account collection
	centralBankAsBytes, err := ctx.GetStub().GetPrivateData(centralBankCollection, centralBankID)
	if err != nil {
		return 0.0, fmt.Errorf("failed to get central bank account from central bank collection: %v", err)
	}
	if centralBankAsBytes == nil {
		return 0.0, fmt.Errorf("central bank account not found in central bank collection: %s", centralBankID)
	}

	// Unmarshal the central bank
	var centralBank CentralBank
	err = json.Unmarshal(centralBankAsBytes, &centralBank)
	if err != nil {
		return 0.0, fmt.Errorf("failed to unmarshal central bank account: %v", err)
	}

	// Get the forex bank's bank account from the bank account collection
	forexBankAsBytes, err := ctx.GetStub().GetPrivateData(forexBankCollection, forexBankID)
	if err != nil {
		return 0.0, fmt.Errorf("failed to get forex bank account from forex bank collection: %v", err)
	}
	if forexBankAsBytes == nil {
		return 0.0, fmt.Errorf("forex bank account not found in forex bank collection: %s", forexBankID)
	}

	// Unmarshal the forex bank
	var forexBank ForexBank
	err = json.Unmarshal(forexBankAsBytes, &forexBank)
	if err != nil {
		return 0.0, fmt.Errorf("failed to unmarshal forex bank account: %v", err)
	}

	// Check if the forex bank's bank account has sufficient balance
	if forexBank.Balance < amount {
		return 0.0, fmt.Errorf("insufficient balance in forex bank's bank account")
	}

	// exchange rate from USD to central bank's currency
	exchangeRateAsBytes, err := ctx.GetStub().GetPrivateData(exchangeRateCollection, "USD"+centralBank.Currency)
	if err != nil {
		return 0.0, fmt.Errorf("failed to get exchange rate from exchange rate collection: %v", err)
	}
	if exchangeRateAsBytes == nil {
		return 0.0, fmt.Errorf("exchange rate not found in exchange rate collection: %s", "USD"+centralBank.Currency)
	}

	// Unmarshal the exchange rate
	var exchangeRate ExchangeRate
	err = json.Unmarshal(exchangeRateAsBytes, &exchangeRate)
	if err != nil {
		return 0.0, fmt.Errorf("failed to unmarshal exchange rate: %v", err)
	}

	amountInCentralBankCurrency := amount * exchangeRate.Value
	

	// Deduct 99% of the amount from the forex bank's bank account
	// 1% is the fee charged by the forex bank
	forexBank.Balance -= 99*amountInCentralBankCurrency/100
	forexBankAsBytes, err = json.Marshal(forexBank)
	if err != nil {
		return 0.0, fmt.Errorf("failed to marshal forex bank account: %v", err)
	}

	err = ctx.GetStub().PutPrivateData(forexBankCollection, forexBank.ID, forexBankAsBytes)
	if err != nil {
		return 0.0, fmt.Errorf("failed to update forex bank account: %v", err)
	}

	centralBank.Balance += 99*amountInCentralBankCurrency/100
	centralBankAsBytes, err = json.Marshal(centralBank)
	if err != nil {
		return 0.0, fmt.Errorf("failed to marshal central bank account: %v", err)
	}

	err = ctx.GetStub().PutPrivateData(centralBankCollection, centralBank.ID, centralBankAsBytes)
	if err != nil {
		return 0.0, fmt.Errorf("failed to update central bank account: %v", err)
	}

	return 99*amountInCentralBankCurrency/100, nil
}

func (s *SmartContract) CentralBankToBankTransaction(ctx contractapi.TransactionContextInterface, transactionID string) error {
	// Get the transaction from the pendingTransaction collection
	transactionAsBytes, err := ctx.GetStub().GetPrivateData(pendingCentralBanktoBankTransactionCollection, transactionID)
	if err != nil {
		return fmt.Errorf("failed to get transaction from pendingCentralBanktoBankTransaction collection: %v", err)
	}
	if transactionAsBytes == nil {
		return fmt.Errorf("transaction not found in pendingCentralBanktoBankTransaction collection: %s", transactionID)
	}

	// Unmarshal the transaction
	var transaction Transaction
	err = json.Unmarshal(transactionAsBytes, &transaction)
	if err != nil {
		return fmt.Errorf("failed to unmarshal transaction: %v", err)
	}

	// Get the employee's bank account from the employee collection
	employeeAsBytes, err := ctx.GetStub().GetPrivateData(employeeCollection, transaction.PayeeID)
	if err != nil {
		return fmt.Errorf("failed to get employee from employee collection: %v", err)
	}
	if employeeAsBytes == nil {
		return fmt.Errorf("employee not found in employee collection: %s", transaction.PayeeID)
	}

	// Unmarshal the employee
	var employee Employee
	err = json.Unmarshal(employeeAsBytes, &employee)
	if err != nil {
		return fmt.Errorf("failed to unmarshal employee: %v", err)
	}

	// Get the employee's bank account from the bank account collection
	bankAccountAsBytes, err := ctx.GetStub().GetPrivateData(bankAccountCollection, employee.BankAccountID+employee.BankID)
	if err != nil {
		return fmt.Errorf("failed to get bank account from bank account collection: %v", err)
	}
	if bankAccountAsBytes == nil {
		return fmt.Errorf("bank account not found in bank account collection: %s", employee.BankAccountID)
	}

	// Unmarshal the bank account
	var bankAccount BankAccount
	err = json.Unmarshal(bankAccountAsBytes, &bankAccount)
	if err != nil {
		return fmt.Errorf("failed to unmarshal bank account: %v", err)
	}

	// Send money from the central bank's bank account to the bank
	err = sendMoneyToBank(ctx, employee.Location+"CentralBank", bankAccount.ID+bankAccount.BankID, transaction.AmountReceived)
	if err != nil {
		return fmt.Errorf("failed to send money to bank: %v", err)
	}

	// Update the transaction status
	transaction.Status = "Completed"
	transactionAsBytes, err = json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction: %v", err)
	}

	err = ctx.GetStub().PutPrivateData(completedTransactionCollection, transactionID, transactionAsBytes)
	if err != nil {
		return fmt.Errorf("failed to put transaction in completedTransaction collection: %v", err)
	}

	// Delete the transaction from the pendingTransaction collection
	err = ctx.GetStub().DelPrivateData(pendingCentralBanktoBankTransactionCollection, transactionID)
	if err != nil {
		return fmt.Errorf("failed to delete transaction from pendingCentralBanktoBankTransaction collection: %v", err)
	}

	return nil
}

func sendMoneyToBank(ctx contractapi.TransactionContextInterface, centralBankID string, bankAccountID string, amount float64) error {
	// Get the central bank's bank account from the bank account collection
	centralBankAsBytes, err := ctx.GetStub().GetPrivateData(centralBankCollection, centralBankID)
	if err != nil {
		return fmt.Errorf("failed to get central bank account from central bank collection: %v", err)
	}
	if centralBankAsBytes == nil {
		return fmt.Errorf("central bank account not found in central bank collection: %s", centralBankID)
	}

	// Unmarshal the central bank
	var centralBank CentralBank
	err = json.Unmarshal(centralBankAsBytes, &centralBank)
	if err != nil {
		return fmt.Errorf("failed to unmarshal central bank account: %v", err)
	}

	// Get the bank's bank account from the bank account collection
	bankAccountAsBytes, err := ctx.GetStub().GetPrivateData(bankAccountPrivateCollection, bankAccountID)
	if err != nil {
		return fmt.Errorf("failed to get bank account from bank account collection: %v", err)
	}
	if bankAccountAsBytes == nil {
		return fmt.Errorf("bank account not found in bank account collection: %s", bankAccountID)
	}

	// Unmarshal the bank account
	var bankAccount BankAccountPrivate
	err = json.Unmarshal(bankAccountAsBytes, &bankAccount)
	if err != nil {
		return fmt.Errorf("failed to unmarshal bank account: %v", err)
	}

	// Check if the central bank's bank account has sufficient balance
	if centralBank.Balance < amount {
		return fmt.Errorf("insufficient balance in central bank's bank account")
	}

	// Deduct the amount from the central bank's bank account
	// 10% is the income tax charged by the central bank
	centralBank.Balance -= amount*90/100
	centralBankAsBytes, err = json.Marshal(centralBank)
	if err != nil {
		return fmt.Errorf("failed to marshal central bank account: %v", err)
	}

	err = ctx.GetStub().PutPrivateData(centralBankCollection, centralBank.ID, centralBankAsBytes)
	if err != nil {
		return fmt.Errorf("failed to update central bank account: %v", err)
	}

	// Add the amount to the bank's bank account
	bankAccount.Balance += 90*amount/100
	bankAccountAsBytes, err = json.Marshal(bankAccount)
	if err != nil {
		return fmt.Errorf("failed to marshal bank account: %v", err)
	}

	err = ctx.GetStub().PutPrivateData(bankAccountPrivateCollection, bankAccountID, bankAccountAsBytes)
	if err != nil {
		return fmt.Errorf("failed to update bank account: %v", err)
	}

	return nil
}

func (s *SmartContract) QueryTransactionByEmployer(ctx contractapi.TransactionContextInterface, employer string) ([]*Transaction, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"PayerID\":\"%v\"}}", employer)

	queryResults, err := s.getQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}
func (s *SmartContract) QueryTransactionByEmployee(ctx contractapi.TransactionContextInterface, employee string) ([]*Transaction, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"PayeeID\":\"%v\"}}", employee)

	queryResults, err := s.getQueryResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}
func (s *SmartContract) QueryBankByCentralBank(ctx contractapi.TransactionContextInterface, centralbank string) ([]*Bank, error) {

	queryString := fmt.Sprintf("{\"selector\":{\"CentralBankID\":\"%v\"}}", centralbank)

	queryResults, err := s.getBankResultForQueryString(ctx, queryString)
	if err != nil {
		return nil, err
	}
	return queryResults, nil
}

// getQueryResultForQueryString executes the passed in query string.
func (s *SmartContract) getQueryResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*Transaction, error) {
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

	resultsIterator, err = ctx.GetStub().GetPrivateDataQueryResult(pendingCentralBankToForexTransactionCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
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

	resultsIterator, err = ctx.GetStub().GetPrivateDataQueryResult(pendingForexToCentralBankTransactionCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
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

	resultsIterator, err = ctx.GetStub().GetPrivateDataQueryResult(pendingCentralBanktoBankTransactionCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
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
func (s *SmartContract) getBankResultForQueryString(ctx contractapi.TransactionContextInterface, queryString string) ([]*Bank, error) {
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(bankCollection, queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []*Bank{}
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset *Bank
		err = json.Unmarshal(response.Value, &asset)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
		}
		results = append(results, asset)
	}

	return results, nil
}

func (s *SmartContract) ReadBank(ctx contractapi.TransactionContextInterface, assetID string) (*Bank, error) {

	log.Printf("ReadAsset: collection %v, ID %v", bankCollection, assetID)
	assetJSON, err := ctx.GetStub().GetPrivateData(bankCollection, assetID) //get the asset from chaincode state
	if err != nil {
		return nil, fmt.Errorf("failed to read asset: %v", err)
	}

	// No Asset found, return empty response
	if assetJSON == nil {
		log.Printf("%v does not exist in collection %v", assetID, bankCollection)
		return nil, nil
	}

	var asset *Bank
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return asset, nil

}
func (s *SmartContract) ReadCentralBank(ctx contractapi.TransactionContextInterface, assetID string) (*CentralBank, error) {

	log.Printf("ReadAsset: collection %v, ID %v", centralBankCollection, assetID)
	assetJSON, err := ctx.GetStub().GetPrivateData(centralBankCollection, assetID) //get the asset from chaincode state
	if err != nil {
		return nil, fmt.Errorf("failed to read asset: %v", err)
	}

	// No Asset found, return empty response
	if assetJSON == nil {
		log.Printf("%v does not exist in collection %v", assetID, centralBankCollection)
		return nil, nil
	}

	var asset *CentralBank
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return asset, nil

}
func (s *SmartContract) ReadForexBank(ctx contractapi.TransactionContextInterface, assetID string) (*ForexBank, error) {

	log.Printf("ReadAsset: collection %v, ID %v", forexBankCollection, assetID)
	assetJSON, err := ctx.GetStub().GetPrivateData(forexBankCollection, assetID) //get the asset from chaincode state
	if err != nil {
		return nil, fmt.Errorf("failed to read asset: %v", err)
	}

	// No Asset found, return empty response
	if assetJSON == nil {
		log.Printf("%v does not exist in collection %v", assetID, forexBankCollection)
		return nil, nil
	}

	var asset *ForexBank
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return asset, nil

}

func (s *SmartContract) GetAllExchangeValues(ctx contractapi.TransactionContextInterface) ([]*ExchangeRate, error) {
	resultsIterator, err := ctx.GetStub().GetPrivateDataByRange(exchangeRateCollection, "", "")
	if err != nil {
		return nil, fmt.Errorf("failed to get exchange rate result for query string: %v", err)
	}
	defer resultsIterator.Close()

	results := []*ExchangeRate{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to iterate over exchange rate query results: %v", err)
		}

		exchangeRate := new(ExchangeRate)
		err = json.Unmarshal(queryResponse.Value, exchangeRate)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal exchange rate query result: %v", err)
		}

		results = append(results, exchangeRate)
	}

	return results, nil
}

func (s *SmartContract) AddPendingTransaction(ctx contractapi.TransactionContextInterface, transactionID string,payerid string,amtrcv float64,payercbid string,payeeid string,amtorg float64, payeecbid string, date string, stat string, contractid string) error {
	assetJSON,err := ctx.GetStub().GetPrivateData(pendingTransactionCollection, transactionID)
	if err != nil {
		return fmt.Errorf("failed to read from pendingTransaction collection: %v", err)
	}
	if assetJSON != nil {
		return fmt.Errorf("transaction already exists in pendingTransaction collection: %s", transactionID)
	}

	transaction := Transaction{
		TransactionId: transactionID,
		PayerID: payerid,
		AmountReceived: amtrcv,
		PayerCentralBankID: payercbid,
		PayeeID: payeeid,
		AmountOriginal: amtorg,
		PayeeCentralBankID: payeecbid,
		PaymentDate: date,
		Status: stat,
		ContractID: contractid,
	}

	transactionJSON, err := json.Marshal(transaction)
	if err != nil {
		return err
	}


	err = ctx.GetStub().PutPrivateData(pendingTransactionCollection, transactionID, []byte(transactionJSON))
	if err != nil {
		return fmt.Errorf("failed to put transaction in pendingTransaction collection: %v", err)
	}

	return nil
}

func (s *SmartContract) GetAccountBalance(ctx contractapi.TransactionContextInterface, accountID string) (float64, error) {
	accountAsBytes, err := ctx.GetStub().GetPrivateData(bankAccountPrivateCollection, accountID)
	if err != nil {
		return 0.0, fmt.Errorf("failed to get account from bank account collection: %v", err)
	}
	if accountAsBytes == nil {
		return 0.0, fmt.Errorf("account not found in bank account collection: %s", accountID)
	}

	var account BankAccountPrivate
	err = json.Unmarshal(accountAsBytes, &account)
	if err != nil {
		return 0.0, fmt.Errorf("failed to unmarshal account: %v", err)
	}

	return account.Balance, nil
}

func (s *SmartContract) ReadPendingTransactionsByCentralBankID(ctx contractapi.TransactionContextInterface, bankID string) ([]*Transaction, error) {
	queryString1 := fmt.Sprintf("{\"selector\":{\"PayerCentralBankID\":\"%v\"}}", bankID)
	queryString2 := fmt.Sprintf("{\"selector\":{\"PayeeCentralBankID\":\"%v\"}}", bankID)
	// queryString := fmt.Sprintf(`{"selector":{"$and":[{"PayerCentralBankID":{"$regex":"%s$"}}]}}`, bankID)
	return s.getQueryResultForQueryStringCBPendingTxn(ctx, queryString1, queryString2)
}

func (s *SmartContract) getQueryResultForQueryStringCBPendingTxn(ctx contractapi.TransactionContextInterface, queryString1 string, queryString2 string) ([]*Transaction, error) {

	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(pendingCentralBankToForexTransactionCollection, queryString1)
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
	resultsIterator2, err := ctx.GetStub().GetPrivateDataQueryResult(pendingCentralBanktoBankTransactionCollection, queryString2)
	if err != nil {
		return nil, err
	}
	defer resultsIterator2.Close()

	for resultsIterator2.HasNext() {
		response, err := resultsIterator2.Next()
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