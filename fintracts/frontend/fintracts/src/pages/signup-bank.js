import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const BankSignUp = () => {
  const navigate = useNavigate();
  const [id, setId] = useState("");
  const [name, setName] = useState("");
  const [centralBankId, setCentralBankId] = useState("");
  const [password, setPassword] = useState("");

  const handleSignUp = async () => {
    // Perform sign-up logic here
    // You can use the values of firstName, lastName, email, and password to create a new user account
    console.log(id, name, centralBankId, password);
    try {
      const response = await axios.post(
        "http://localhost:3002/invoke",
        new URLSearchParams([
          ["", ""],
          ["channelid", "bankchannel"],
          ["chaincodeid", "bank"],
          ["function", "CreateBank"],
          ["args", id],
          ["args", name],
          ["args", centralBankId],
          ["args", password],
        ])
      );
      console.log("Response:", response);
      alert("Bank signed up successfully");
      navigate("/login");
    } catch (error) {
      alert("Error signing up Bank");
      console.error("Error:", error);
    }
  };

  return (
    <div>
      <h2>Bank Sign Up</h2>
      <form>
        <label>
          Bank ID:
          <input
            type="text"
            value={id}
            onChange={(e) => setId(e.target.value)}
          />
        </label>
        <br />
        <label>
          Password:
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </label>
        <br />
        <label>
          Name:
          <input
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </label>
        <br />
        <label>
          Central Bank ID:
          <input
            type="text"
            value={centralBankId}
            onChange={(e) => setCentralBankId(e.target.value)}
          />
        </label>
        <br />
        <button type="button" onClick={handleSignUp}>
          Sign Up
        </button>
      </form>
    </div>
  );
};

export default BankSignUp;
