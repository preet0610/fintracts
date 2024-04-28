import React, { useState, useEffect } from "react";
import axios from "axios";

const BankPending = () => {
  const [pendingTransactions, setPendingTransactions] = useState([]);
  const user = JSON.parse(localStorage.getItem("user"));

  useEffect(() => {
    // Fetch pending transactions from the backend API
    const fetchPendingTransactions = async () => {
      try {
        const response = await axios.get("http://localhost:3000/query", {
          params: {
            channelid: "bankchannel",
            chaincodeid: "emp",
            function: "ReadPendingTransactionsByBankID",
            args: user.ID,
          },
        });
        console.log("Pending transactions:", response.data);
        setPendingTransactions(JSON.parse(response.data.substring(9)));
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
        <ul>
          {pendingTransactions.map((transaction) => (
            <li key={transaction.id}>
              {/* Display transaction details */}
              <p>Transaction ID: {transaction.TransactionId}</p>
              <p>Sender: {transaction.PayerID}</p>
              <p>Amount: {transaction.AmountReceived}</p>
              <p>Recipient: {transaction.PayerCentralBankID}</p>
              <button
                onClick={async () => {
                  // Approve the transaction
                  const response = await axios.post(
                    "http://localhost:3002/invoke",
                    new URLSearchParams([
                      ["", ""],
                      ["channelid", "bankchannel"],
                      ["chaincodeid", "bank"],
                      ["function", "AddPendingTransaction"],
                      ["args", transaction.TransactionId],
                      ["args", transaction.PayerID],
                      ["args", transaction.AmountReceived],
                      ["args", transaction.PayerCentralBankID],
                      ["args", transaction.PayeeID],
                      ["args", transaction.AmountOriginal],
                      ["args", transaction.PayeeCentralBankID],
                      ["args", transaction.PaymentDate],
                      ["args", transaction.Status],
                      ["args", transaction.ContractID],
                    ])
                  );
                  console.log("Response:", response);
                  console.log(JSON.stringify(transaction));
                  // const response2 = await axios.post(
                  //   "http://localhost:3000/invoke",
                  //   new URLSearchParams([
                  //     ["", ""],
                  //     ["channelid", "bankchannel"],
                  //     ["chaincodeid", "emp"],
                  //     ["function", "DeletePendingTransaction"],
                  //     ["args", transaction.TransactionId],
                  //   ])
                  // );
                  // console.log("Response:", response2);
                  // window.location.reload();
                  console.log(
                    "Approving transaction:",
                    transaction.TransactionId
                  );
                  const response3 = await axios.post(
                    "http://localhost:3002/invoke",
                    new URLSearchParams([
                      ["", ""],
                      ["channelid", "bankchannel"],
                      ["chaincodeid", "bank"],
                      ["function", "BankToCentralBankTransaction"],
                      ["args", transaction.TransactionId],
                    ])
                  );
                  //   console.log("Response:", response3);
                }}
              >
                Approve
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default BankPending;
