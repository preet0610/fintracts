import React, { useState, useEffect } from "react";
import axios from "axios";

const ForexPending = () => {
  const [pendingTransactions, setPendingTransactions] = useState([]);
  const user = JSON.parse(localStorage.getItem("user"));

  useEffect(() => {
    // Fetch pending transactions from the backend API
    const fetchPendingTransactions = async () => {
      try {
        const response = await axios.get("http://localhost:3002/query", {
          params: {
            channelid: "bankchannel",
            chaincodeid: "bank",
            function: "ReadPendingForexTransactions",
            args: user.ID,
          },
        });
        console.log("Pending transactions to Forex:", response.data);
        setPendingTransactions(JSON.parse(response.data.substring(9)));
        console.log(user.ID);
      } catch (error) {
        console.error("Error fetching pending transactions:", error);
      }
    };

    fetchPendingTransactions();
  }, []);

  return (
    <div>
      <h1>Pending Transactions</h1>
      {pendingTransactions.length === 0 ? (
        <p>No pending transactions</p>
      ) : (
        <div>
          <ul>
            {pendingTransactions.map((transaction) => (
              <li key={transaction.id}>
                {/* Display transaction details */}
                <p>Transaction ID: {transaction.TransactionId}</p>
                <p>Sender: {transaction.PayerID}</p>
                <p>Amount: {transaction.AmountReceived}</p>
                <p>Recipient: {transaction.PayeeCentralBankID}</p>
                <button
                  onClick={async () => {
                    // Approve the transaction
                    console.log(
                      "Approving transaction:",
                      transaction.TransactionId
                    );
                    try {
                      const response3 = await axios.post(
                        "http://localhost:3002/invoke",
                        new URLSearchParams([
                          ["", ""],
                          ["channelid", "bankchannel"],
                          ["chaincodeid", "bank"],
                          ["function", "ForexToCentralBankTransaction"],
                          ["args", transaction.TransactionId],
                        ])
                      );
                      console.log("Response:", response3);
                    } catch (error) {
                      console.error("Error approving transaction:", error);
                    }
                  }}
                >
                  Approve
                </button>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default ForexPending;
