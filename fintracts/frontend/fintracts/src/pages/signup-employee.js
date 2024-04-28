import React, { useState } from "react";
import axios from "axios";
import { useNavigate } from "react-router-dom";

const EmployeeSignUp = () => {
  const [usrname, setUsrname] = useState("");
  const [Name, setName] = useState("");
  const [country, setCountry] = useState("");
  const [dob, setDob] = useState("");
  const [bankID, setBankID] = useState("");
  const [bankAccountNumber, setBankAccountNumber] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();
  const handleSignUp = async () => {
    // Perform sign-up logic here
    // You can use the values of firstName, lastName, email, and password to create a new employee account
    try {
      const response = await axios.post(
        "http://localhost:3000/invoke",
        new URLSearchParams([
          ["", ""],
          ["channelid", "bankchannel"],
          ["chaincodeid", "emp"],
          ["function", "CreateEmployee"],
          ["args", usrname],
          ["args", Name],
          ["args", country],
          ["args", dob],
          ["args", bankID],
          ["args", bankAccountNumber],
          ["args", password],
        ])
      );
      console.log("Response:", response);
      alert("Employee signed up successfully");
      navigate("/login");
    } catch (error) {
      alert("Error signing up employee");
      console.error("Error:", error);
    }
  };

  return (
    <div>
      <h2>Employee Sign Up</h2>
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
          DOB:
          <input
            type="text"
            value={dob}
            onChange={(e) => setDob(e.target.value)}
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

export default EmployeeSignUp;
