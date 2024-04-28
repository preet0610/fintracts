import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const ForexBankSignUp = () => {
  const navigate = useNavigate();
  const [userName, setUserName] = useState("");
  const [balance, setBalance] = useState("");
  const [password, setPassword] = useState("");

  const handleSignUp = async () => {
    // Implement your sign-up logic here
    // You can use the values of firstName, lastName, email, and password
    // to send a request to your backend API or perform any other necessary actions
    console.log(userName, balance, password);
    try {
      const response = await axios.post(
        "http://localhost:3002/invoke",
        new URLSearchParams([
          ["", ""],
          ["channelid", "bankchannel"],
          ["chaincodeid", "bank"],
          ["function", "CreateForexBank"],
          ["args", userName],
          ["args", balance],
          ["args", password],
        ])
      );
      console.log("Response:", response);
      alert("Forex Bank signed up successfully");
      navigate("/login");
    } catch (error) {
      alert("Error signing up Forex Bank");
      console.error("Error:", error);
    }
  };

  return (
    <div>
      <h1>Forex Bank Sign Up</h1>
      <form>
        <label>
          Username:
          <input
            type="text"
            value={userName}
            onChange={(e) => setUserName(e.target.value)}
          />
        </label>
        <br />
        <label>
          Balance:
          <input
            type="integer"
            value={balance}
            onChange={(e) => setBalance(e.target.value)}
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
        <button type="button" onClick={handleSignUp}>
          Sign Up
        </button>
      </form>
    </div>
  );
};

export default ForexBankSignUp;
