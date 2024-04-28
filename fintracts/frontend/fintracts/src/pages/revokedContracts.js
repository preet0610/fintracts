import React, { useEffect, useState } from "react";
import axios from "axios";
import Login from "./login";

const RevokedContracts = () => {
  const user = window.localStorage.getItem("user");
  const userJson = JSON.parse(user);
  const [revokedContracts, setRevokedContracts] = useState([]);
  var func = userJson["Industry"]
    ? "QueryRevokedContractByEmployer"
    : "QueryRevokedContractByEmployee";
  useEffect(() => {
    const fetchContracts = async () => {
      try {
        const response = await axios.get("http://localhost:3000/query", {
          params: {
            channelid: "bankchannel",
            chaincodeid: "emp",
            function: userJson["Industry"]
              ? "QueryRevokedContractByEmployer"
              : "QueryRevokedContractByEmployee",
            args: userJson["ID"],
          },
        });
        setRevokedContracts([...JSON.parse(response.data.substring(9))]);
        console.log(userJson["ID"]);
        // console.log(revokedContracts);
      } catch (error) {
        console.error("Error fetching contracts:", error);
      }
    };
    fetchContracts();
  }, []);
  if (!window.localStorage.getItem("isLoggedIn")) {
    return <Login />;
  }
  console.log(revokedContracts);
  return (
    <>
      <h3>Revoked Contracts</h3>
      <div>
        Contract ID | Employer ID | Employee ID | Designation | Working Hours |
        CTC | Payment Cycle |
      </div>
      {revokedContracts.map((contract) => {
        return (
          <>
            <div>
              {contract.ID} | {contract.EmployerID} | {contract.EmployeeID} |{" "}
              {contract.Designation} | {contract.WorkingHours} | {contract.CTC}{" "}
              | {contract.PaymentCycle} |
            </div>
          </>
        );
      })}
    </>
  );
};

export default RevokedContracts;
