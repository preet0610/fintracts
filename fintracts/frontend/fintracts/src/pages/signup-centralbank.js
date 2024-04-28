import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const CentralBankSignup = () => {
  // const [username, setUsername] = useState("");
  const navigate = useNavigate();
  const [password, setPassword] = useState("");
  const [location, setLocation] = useState("");
  const [Currency, setCurrency] = useState("");

  const handleSignUp = async () => {
    console.log("Location: ", location);
    console.log("Password: ", password);
    console.log("Currency: ", Currency);

    try {
      const response = await axios.post(
        "http://localhost:3002/invoke",
        new URLSearchParams([
          ["", ""],
          ["channelid", "bankchannel"],
          ["chaincodeid", "bank"],
          ["function", "CreateCentralBank"],
          ["args", location + "CentralBank"],
          ["args", location],
          ["args", Currency],
          ["args", password],
        ])
      );
      console.log("Response:", response);
      alert("Central Bank signed up successfully");
      navigate("/login");
    } catch (error) {
      alert("Error signing up Central Bank");
      console.error("Error:", error);
    }
  };

  return (
    <div>
      <h2>Central Bank Sign Up</h2>
      <form>
        <label>
          Country:
          <input
            type="text"
            value={location}
            onChange={(e) => setLocation(e.target.value)}
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
          Currency:
          <input
            type="text"
            value={Currency}
            onChange={(e) => setCurrency(e.target.value)}
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

export default CentralBankSignup;
