import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const Signup = () => {
  const history = useNavigate();
  const [accountType, setAccountType] = useState("");

  const handleAccountTypeSelection = (type) => {
    setAccountType(type);
    // Redirect to different pages based on account type
    if (type === "employer") {
      history("/signup/employer");
    } else if (type === "employee") {
      history("/signup/employee");
    } else if (type === "bank") {
      history("/signup/bank");
    } else if (type === "centralbank") {
      history("/signup/centralbank");
    } else if (type === "forex") {
      history("/signup/forex");
    } else {
      console.error("Invalid account type selected");
    }
  };

  return (
    <div>
      <div>
        <h1>Signup</h1>
        <p>Select account type:</p>
        <button onClick={() => handleAccountTypeSelection("employer")}>
          Employer
        </button>
        <button onClick={() => handleAccountTypeSelection("employee")}>
          Employee
        </button>
        <button onClick={() => handleAccountTypeSelection("bank")}>Bank</button>
        <button onClick={() => handleAccountTypeSelection("centralbank")}>
          Central Bank
        </button>
        <button onClick={() => handleAccountTypeSelection("forex")}>
          Foreign Exchange Bank
        </button>
      </div>
      <div>
        <p>Already have an account?</p>
        <button onClick={() => history("/login")}>Login</button>
      </div>
    </div>
  );
};

export default Signup;
