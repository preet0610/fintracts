package web

import (
	"fmt"
	"net/http"
)

// Query handles chaincode query requests.
func (setup OrgSetup) Query(w http.ResponseWriter, r *http.Request) {
	switch origin := r.Header.Get("Origin"); origin {
	case "http://localhost:3001":
		w.Header().Set("Access-Control-Allow-Origin", origin)
		case "http://172.23.0.77:3001":
	}
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	fmt.Println("Received Query request")
	queryParams := r.URL.Query()
	chainCodeName := queryParams.Get("chaincodeid")
	channelID := queryParams.Get("channelid")
	function := queryParams.Get("function")
	args := r.URL.Query()["args"]
	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)
	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	evaluateResponse, err := contract.EvaluateTransaction(function, args...)
	if err != nil {
		// fmt.Fprintf(w, "Error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error evaluating transaction"))
		return
	}
	fmt.Fprintf(w, "Response: %s", evaluateResponse)
}
