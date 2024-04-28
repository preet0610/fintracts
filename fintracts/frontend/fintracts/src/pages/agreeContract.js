import React, { useState, useEffect } from "react";
import axios from "axios";
import Login from "./login";
const AgreeContract = () => {
  const user = window.localStorage.getItem("user");
  const userJson = JSON.parse(user);
  const [contracts, setContracts] = useState([]);
  useEffect(() => {
    const fetchContracts = async () => {
      try {
        const response = await axios.get("http://localhost:3000/query", {
          params: {
            channelid: "bankchannel",
            chaincodeid: "emp",
            function: "QueryPendingContractByEmployee",
            args: userJson["ID"],
          },
        });
        setContracts([...JSON.parse(response.data.substring(9))]);
        console.log(contracts);
      } catch (error) {
        console.error("Error fetching contracts:", error);
      }
    };
    fetchContracts();
  }, []);
  if (!window.localStorage.getItem("isLoggedIn")) {
    return <Login />;
  }

  return (
    <>
      <h3>Pending Contracts</h3>
      <div>
        Contract ID | Employer ID | Designation | Working Hours | CTC | Payment
        Cycle |
      </div>
      {contracts.map((contract) => {
        return (
          <>
            <div>
              {contract.ID} | {contract.EmployerID} | {contract.Designation} |{" "}
              {contract.WorkingHours} | {contract.CTC} | {contract.PaymentCycle}{" "}
              |
              <button
                onClick={async () => {
                  try {
                    const response = await axios.post(
                      "http://localhost:3000/invoke",
                      new URLSearchParams([
                        ["", ""],
                        ["channelid", "bankchannel"],
                        ["chaincodeid", "emp"],
                        ["function", "RejectContract"],
                        ["args", contract.ID],
                      ])
                    );
                    console.log("Response:", response);
                    alert("Contract Rejected successfully");
                    window.location.reload();
                  } catch (error) {
                    alert("Error rejecting contract");
                    console.error("Error:", error);
                  }
                }}
              >
                Reject
              </button>
              <button
                onClick={async () => {
                  try {
                    const response = await axios.post(
                      "http://localhost:3000/invoke",
                      new URLSearchParams([
                        ["", ""],
                        ["channelid", "bankchannel"],
                        ["chaincodeid", "emp"],
                        ["function", "AgreeToContract"],
                        ["args", contract.ID],
                      ])
                    );
                    console.log("Response:", response);
                    alert("Contract Agreed successfully");
                    window.location.reload();
                  } catch (error) {
                    alert("Error agreeing contract");
                    console.error("Error:", error);
                  }
                }}
              >
                Agree
              </button>
            </div>
          </>
        );
      })}
    </>
  );
};

export default AgreeContract;
