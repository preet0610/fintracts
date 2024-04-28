import React, { useState, useEffect } from "react";
import axios from "axios";
import "../App.css";
function ContractsEmployee() {
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
            function: "QueryContractByEmployee",
            args: userJson["ID"],
          },
        });
        console.log("Response:", JSON.parse(response.data.substring(9)));
        setContracts([...JSON.parse(response.data.substring(9))]);
      } catch (error) {
        console.error("Error:", error);
      }
    };
    fetchContracts();
  }, []);

  return (
    <>
      <h3>Ongoing Contracts</h3>
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
                        ["function", "RevokeContract"],
                        ["args", contract.ID],
                      ])
                    );
                    console.log("Response:", response);
                    alert("Contract revoked successfully");
                    window.location.reload();
                  } catch (error) {
                    alert("Error revoking contract");
                    console.error("Error:", error);
                  }
                }}
              >
                Revoke
              </button>
            </div>
          </>
        );
      })}
    </>
  );
}

export default ContractsEmployee;
