import React, { useState, useEffect } from "react";
import axios from "axios";

const CentralBankPending = () => {
  const [pendingTransactionsToForex, setPendingTransactionsToForex] = useState(
    []
  );
  const [pendingTransactionsToBank, setPendingTransactionsToBank] = useState(
    []
  );
  const user = JSON.parse(localStorage.getItem("user"));

  useEffect(() => {
    // Fetch pending transactions from the backend API
    const fetchPendingTransactionsToForex = async () => {
      try {
        const response = await axios.get("http://localhost:3002/query", {
          params: {
            channelid: "bankchannel",
            chaincodeid: "bank",
            function: "ReadPendingTransactionsToForexByCentralBankID",
            args: user.ID,
          },
        });
        console.log("Pending transactions to Forex:", response.data);
        setPendingTransactionsToForex(JSON.parse(response.data.substring(9)));
        console.log(user.ID);
      } catch (error) {
        console.error("Error fetching pending transactions:", error);
      }
    };

    fetchPendingTransactionsToForex();
    const fetchPendingTransactionsToBank = async () => {
      try {
        const response = await axios.get("http://localhost:3002/query", {
          params: {
            channelid: "bankchannel",
            chaincodeid: "bank",
            function: "ReadPendingTransactionsToBankByCentralBankID",
            args: user.ID,
          },
        });
        console.log("Pending transactions to Bank:", response.data);
        setPendingTransactionsToBank(JSON.parse(response.data.substring(9)));
        console.log(user.ID);
      } catch (error) {
        console.error("Error fetching pending transactions:", error);
      }
    };

    fetchPendingTransactionsToBank();
  }, []);

  return (
    <div>
      <h1>Pending Transactions</h1>
      {pendingTransactionsToForex.length + pendingTransactionsToBank.length ===
      0 ? (
        <p>No pending transactions</p>
      ) : (
        <div>
          <ul>
            {pendingTransactionsToForex.map((transaction) => (
              <li key={transaction.id}>
                {/* Display transaction details */}
                <p>Transaction ID: {transaction.TransactionId}</p>
                <p>Sender: {transaction.PayerID}</p>
                <p>Amount: {transaction.AmountReceived}</p>
                <p>Recipient: Forex Bank</p>
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
                          ["function", "CentralBankToForexBankTransaction"],
                          ["args", transaction.TransactionId],
                          ["args", transaction.PayerCentralBankID],
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
          <ul>
            {pendingTransactionsToBank.map((transaction) => (
              <li key={transaction.id}>
                {/* Display transaction details */}
                <p>Transaction ID: {transaction.TransactionId}</p>
                <p>Sender: {transaction.PayerID}</p>
                <p>Amount: {transaction.AmountReceived}</p>
                <p>Recipient: {transaction.PayeeID}</p>
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
                          ["function", "CentralBankToBankTransaction"],
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

export default CentralBankPending;
