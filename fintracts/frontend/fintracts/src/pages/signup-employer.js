import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const EmployerSignUp = () => {
  const navigate = useNavigate();
  const [usrname, setUsrname] = useState("");
  const [Name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [country, setCountry] = useState("");
  const [industry, setIndustry] = useState("");
  const [bankID, setBankID] = useState("");
  const [bankAccountNumber, setBankAccountNumber] = useState("");

  const handleSignUp = async () => {
    // Perform sign-up logic here
    console.log("Signing up employer...");
    console.log("Name:", Name);
    console.log("Email:", usrname);
    console.log("Password:", password);
    // Add your sign-up logic here
    try {
      const response = await axios.post(
        "http://localhost:3000/invoke",
        new URLSearchParams([
          ["", ""],
          ["channelid", "bankchannel"],
          ["chaincodeid", "emp"],
          ["function", "CreateEmployer"],
          ["args", usrname],
          ["args", Name],
          ["args", country],
          ["args", industry],
          ["args", bankID],
          ["args", bankAccountNumber],
          ["args", password],
        ])
      );
      console.log("Response:", response);
      alert("Employer signed up successfully");
      navigate("/login");
    } catch (error) {
      alert("Error signing up employer");
      console.error("Error:", error);
    }
  };

  return (
    <div>
      <h2>Employer Sign Up</h2>
      <form>
        <label>
          Username:
          <input
            type="text"
            value={usrname}
            onChange={(e) => setUsrname(e.target.value)}
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
            value={Name}
            onChange={(e) => setName(e.target.value)}
          />
        </label>
        <br />
        <label>
          Country:
          <input
            type="text"
            value={country}
            onChange={(e) => setCountry(e.target.value)}
          />
        </label>
        <br />
        <label>
          Industry:
          <input
            type="text"
            value={industry}
            onChange={(e) => setIndustry(e.target.value)}
          />
        </label>
        <br />
        <label>
          Bank ID:
          <input
            type="text"
            value={bankID}
            onChange={(e) => setBankID(e.target.value)}
          />
        </label>
        <br />
        <label>
          Bank Account Number:
          <input
            type="text"
            value={bankAccountNumber}
            onChange={(e) => setBankAccountNumber(e.target.value)}
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

export default EmployerSignUp;
