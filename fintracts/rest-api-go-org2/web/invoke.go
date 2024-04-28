package web

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// Invoke handles chaincode invoke requests.
func (setup *OrgSetup) Invoke(w http.ResponseWriter, r *http.Request) {
	switch origin := r.Header.Get("Origin"); origin {
	case "http://localhost:3001":
		w.Header().Set("Access-Control-Allow-Origin", origin)
		case "http://172.23.0.77:3001":
	}
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	fmt.Println("Received Invoke request")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %s", err)
		fmt.Println("Error parsing form")
		w.Write([]byte("Error parsing form"))
		return
	}
	chainCodeName := r.FormValue("chaincodeid")
	channelID := r.FormValue("channelid")
	function := r.FormValue("function")
	args := r.Form["args"]
	// transient := r.Form["transient"]
	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)
	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	txn_proposal, err := contract.NewProposal(function, client.WithArguments(args...))
	if err != nil {
		// fmt.Fprintf(w, "Error creating txn proposal: %s", err)
		fmt.Println("Error creating txn proposal")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating txn proposal"))
		return
	}
	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		// fmt.Fprintf(w, "Error endorsing txn: %s", err)
		fmt.Println("Error endorsing txn")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error endorsing txn"))
		return
	}
	txn_committed, err := txn_endorsed.Submit()
	if err != nil {
		// fmt.Fprintf(w, "Error submitting transaction: %s", err)
		fmt.Println("Error submitting transaction")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error submitting transaction"))
		return
	}
	fmt.Fprintf(w, "Transaction ID : %s Response: %s", txn_committed.TransactionID(), txn_endorsed.Result())
	fmt.Println(txn_endorsed.Result())
}
