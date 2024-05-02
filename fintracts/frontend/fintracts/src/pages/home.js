import React from "react";
import { useNavigate } from "react-router-dom";
import Login from "./login";
import ContractsEmployer from "./contractsEmployer";
import ContractsEmployee from "./contractsEmployee";
import ForexVals from "./forexVals";

const Home = () => {
  const user = window.localStorage.getItem("user");
  const userJson = JSON.parse(user);
  const role = window.localStorage.getItem("role");
  const navigate = useNavigate();
  if (!window.localStorage.getItem("isLoggedIn")) {
    return <Login />;
  }
  const Logout = () => {
    window.localStorage.removeItem("isLoggedIn");
    window.localStorage.removeItem("user");
    window.localStorage.removeItem("role");
    navigate("/");
  };
  if (role === "Employer" || role === "Employee") {
    return (
      <div>
        <div className="navbar">
          {/* Navigation bar content */}
          <a href="/home">Home</a>
          <a href="/about">About</a>
          <a href="/contact">Contact</a>
          <a href="/veiw-balance">View Balance</a>
          {userJson["Industry"] ? (
            <a href="/create-contract">Create Contract</a>
          ) : (
            <a href="/agree-to-contract">New Contracts</a>
          )}
          <a href="/revoked-contracts">Revoked Contracts</a>
          <a onClick={() => Logout()}>Logout</a>
        </div>
        <h1>Welcome {userJson["Name"]}!</h1>
        <div>
          {
            /* Display contracts based on user role */
            userJson["Industry"] ? <ContractsEmployer /> : <ContractsEmployee />
          }
        </div>
        <button onClick={() => Logout()}>Log Out</button>
      </div>
    );
  } else {
    return (
      <div>
        <div className="navbar">
          {/* Navigation bar content */}
          <a href="/home">Home</a>
          <a href="/about">About</a>
          <a href="/contact">Contact</a>
          <a href="/veiw-balance">View Balance</a>
          {role === "ForexBank" ? (
            <a href="/add-exchange-rate">Add Exchange Rate</a>
          ) : role === "CentralBank" ? (
            <a href="/central-bank-pending-txns">Pending Transactions</a>
          ) : (
            <a href="/bank-pending-txns">Pending Transactions</a>
          )}
          {role === "Bank" ? (
            <a href="/create-bank-account">Create Bank Account</a>
          ) : role === "ForexBank" ? (
            <a href="/forex-pending-txns">Pending Transactions</a>
          ) : null}
          <a onClick={() => Logout()}>Logout</a>
        </div>
        <h1>Welcome {userJson["Name"]}!</h1>
        <div></div>
        {role === "ForexBank" ? <ForexVals /> : null}
        {role === "CentralBank" ? <h2>Central Bank</h2> : null}
        {role === "Bank" ? <h2>Bank</h2> : null}
        <button onClick={() => Logout()}>Log Out</button>
      </div>
    );
  }
};

export default Home;
