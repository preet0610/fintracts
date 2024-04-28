import React, { useState } from "react";
import Login from "./login";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const CreateContract = () => {
  const navigate = useNavigate();
  const user = window.localStorage.getItem("user");
  const userJson = JSON.parse(user);
  const [id, setId] = useState("");
  const [employeeid, setEmployeeId] = useState("");
  const [ctc, setCtc] = useState("");
  const [designation, setDesignation] = useState("");
  const [workinghours, setWorkingHours] = useState("");
  const [paymentcycle, setPaymentCycle] = useState("Monthly");
  if (!window.localStorage.getItem("isLoggedIn")) {
    return <Login />;
  }
  const handleIdChange = (e) => {
    setId(e.target.value);
  };
  const handleEmployeeIdChange = (e) => {
    setEmployeeId(e.target.value);
  };
  const handleCtcChange = (e) => {
    setCtc(e.target.value);
  };
  const handleDesignationChange = (e) => {
    setDesignation(e.target.value);
  };
  const handleWorkingHoursChange = (e) => {
    setWorkingHours(e.target.value);
  };
  const handlePaymentCycleChange = (e) => {
    setPaymentCycle(e.target.value);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    // Handle form submission here
    console.log("Contract ID:", id);
    console.log("Employee ID:", employeeid);
    console.log("Employer ID:", userJson["ID"]);
    console.log("CTC:", ctc);
    console.log("Designation:", designation);
    console.log("Working Hours:", workinghours);
    console.log("Payment Cycle:", paymentcycle);
    // You can send the form data to an API or perform any other actions
    try {
      const response = await axios.post(
        "http://localhost:3000/invoke",
        new URLSearchParams([
          ["", ""],
          ["channelid", "bankchannel"],
          ["chaincodeid", "emp"],
          ["function", "CreateContract"],
          ["args", id],
          ["args", employeeid],
          ["args", userJson["ID"]],
          ["args", ctc],
          ["args", workinghours],
          ["args", designation],
          ["args", paymentcycle],
        ])
      );
      console.log("Response:", response);
      alert("Contract created successfully");
      navigate("/home");
    } catch (error) {
      alert("Error creating contract");
      console.error("Error:", error);
    }
  };

  return (
    <div>
      <h1>Create Contract</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Contract ID:
          <input
            type="text"
            name="field1"
            value={id}
            onChange={handleIdChange}
          />
        </label>
        <br />
        <label>
          Employee ID:
          <input
            type="text"
            name="field2"
            value={employeeid}
            onChange={handleEmployeeIdChange}
          />
        </label>
        <br />
        <label>
          CTC:
          <input
            type="integer"
            name="field1"
            value={ctc}
            onChange={handleCtcChange}
          />
        </label>
        <br />
        <label>
          Designation:
          <input
            type="text"
            name="field1"
            value={designation}
            onChange={handleDesignationChange}
          />
        </label>
        <br />
        <label>
          Working Hours:
          <input
            type="integer"
            name="field1"
            value={workinghours}
            onChange={handleWorkingHoursChange}
          />
        </label>
        <br />
        <label>
          Payment Cycle:
          <select value={paymentcycle} onChange={handlePaymentCycleChange}>
            <option value="Monthly">Monthly</option>
            <option value="Quarterly">Quarterly</option>
            <option value="Half-Yearly">Half Yearly</option>
            <option value="Yearly">Yearly</option>
            <option value="One-Time">One Time</option>
          </select>
        </label>
        <br />
        <button type="submit">Submit</button>
      </form>
    </div>
  );
};

export default CreateContract;
