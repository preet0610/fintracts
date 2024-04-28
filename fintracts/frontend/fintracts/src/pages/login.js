import React, { useState } from "react";
import "../App.css";
import axios from "axios";
import { useNavigate } from "react-router-dom";
const Login = () => {
  const history = useNavigate();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState("Employer");

  const handleUsernameChange = (e) => {
    setUsername(e.target.value);
  };

  const handlePasswordChange = (e) => {
    setPassword(e.target.value);
  };

  const handleRoleChange = (e) => {
    setRole(e.target.value);
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    window.localStorage.setItem("isLoggedIn", "true");
    // Add your login logic here
    console.log("Username:", username);
    console.log("Password:", password);
    console.log("Role:", role);
    if (role === "Employer" || role === "Employee") {
      try {
        const response = await axios.get("http://localhost:3000/query", {
          params: {
            channelid: "bankchannel",
            chaincodeid: "emp",
            function: "Read" + role,
            args: username,
          },
        });
        console.log("Response:", response);
        if (JSON.parse(response.data.substring(9))["Password"] === password) {
          window.localStorage.setItem("user", response.data.substring(9));
          window.localStorage.setItem("role", role);
          alert("Login successful");
          history("/home");
        } else {
          alert("Invalid credentials");
        }
      } catch (error) {
        alert("Error logging in");
        console.error("Error:", error);
      }
    } else {
      try {
        const response = await axios.get("http://localhost:3002/query", {
          params: {
            channelid: "bankchannel",
            chaincodeid: "bank",
            function: "Read" + role,
            args: username,
          },
        });
        console.log("Response:", response);
        if (JSON.parse(response.data.substring(9))["Password"] === password) {
          window.localStorage.setItem("user", response.data.substring(9));
          window.localStorage.setItem("role", role);
          alert("Login successful");
          history("/home");
        } else {
          alert("Invalid credentials");
        }
      } catch (error) {
        alert("Error logging in");
        console.error("Error:", error);
      }
    }
  };

  return (
    <div className="App">
      <h1>Login</h1>
      <form onSubmit={handleSubmit}>
        <label>
          Username:
          <input type="text" value={username} onChange={handleUsernameChange} />
        </label>
        <br />
        <label>
          Password:
          <input
            type="password"
            value={password}
            onChange={handlePasswordChange}
          />
        </label>
        <br />
        <label>
          <select value={role} onChange={handleRoleChange}>
            <option value="Employer">Employer</option>
            <option value="Employee">Employee</option>
            <option value="Bank">Bank</option>
            <option value="CentralBank">Central Bank</option>
            <option value="ForexBank">Forex Bank</option>
          </select>
        </label>
        <br />
        <button type="submit">Login</button>
      </form>
      <div>
        <p>New User?</p>
        <button onClick={() => history("/signup")}>Register</button>
      </div>
    </div>
  );
};

export default Login;
