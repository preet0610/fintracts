import React, { useEffect, useState } from "react";
import axios from "axios";
import Login from "./login";

const ViewBalance = () => {
  const user = JSON.parse(localStorage.getItem("user"));
  const role = localStorage.getItem("role");
  const [balance, setBalance] = useState(null);
  useEffect(() => {
    console.log(role);
    if (role === "Bank" || role === "CentralBank" || role === "ForexBank") {
      setBalance(user.Balance);
    } else {
      const fetchBalance = async () => {
        try {
          const response = await axios.get("http://localhost:3002/query", {
            params: {
              channelid: "bankchannel",
              chaincodeid: "bank",
              function: "GetAccountBalance",
              args: user.BankAccountID + user.BankID,
            },
          });
          console.log("Balance:", response.data);
          setBalance(Number(response.data.substring(9)));
        } catch (error) {
          console.error("Error fetching balance:", error);
        }
      };

      fetchBalance();
    }
  }, []);
  if (!user) {
    return <Login />;
  }
  return (
    <div>
      <h1>Account Balance</h1>
      {balance !== null ? (
        <p>Your balance is: {balance}</p>
      ) : (
        <p>Please contact bank to create account</p>
      )}
    </div>
  );
};

export default ViewBalance;
