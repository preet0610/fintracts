import React, { useState } from "react";
import axios from "axios";

const CreateBankAcc = () => {
  const user = JSON.parse(localStorage.getItem("user"));
  const [accountNumber, setAccountNumber] = useState("");
  const [accountHolder, setAccountHolder] = useState("");
  const [DOB, setDOB] = useState("");
  const [balance, setBalance] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      console.log(
        accountNumber,
        accountHolder,
        DOB,
        balance,
        user.ID,
        user.CentralBankID
      );
      //   // Make an API call to send the form data to the database
      const response = await axios.post(
        "http://localhost:3002/invoke",
        new URLSearchParams([
          ["", ""],
          ["channelid", "bankchannel"],
          ["chaincodeid", "bank"],
          ["function", "CreateBankAccount"],
          ["args", accountNumber],
          ["args", accountHolder],
          ["args", DOB],
          ["args", balance],
          ["args", user.ID],
          ["args", user.CentralBankID],
        ])
      );

      // Handle the response from the API call
      console.log(response.data);

      // Reset the form fields
      setAccountNumber("");
      setAccountHolder("");
      setDOB("");
      setBalance("");
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <h1>Create Bank Account</h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor="accountNumber">Account Number:</label>
        <input
          type="text"
          id="accountNumber"
          value={accountNumber}
          onChange={(e) => setAccountNumber(e.target.value)}
        />

        <label htmlFor="accountHolder">Account Holder:</label>
        <input
          type="text"
          id="accountHolder"
          value={accountHolder}
          onChange={(e) => setAccountHolder(e.target.value)}
        />

        <label htmlFor="DOB">DOB:</label>
        <input
          type="text"
          id="DOB"
          value={DOB}
          onChange={(e) => setDOB(e.target.value)}
        />

        <label htmlFor="balance">Balance:</label>
        <input
          type="number"
          id="balance"
          value={balance}
          onChange={(e) => setBalance(e.target.value)}
        />

        <button type="submit">Create Account</button>
      </form>
    </div>
  );
};

export default CreateBankAcc;
