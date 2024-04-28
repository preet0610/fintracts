import React, { useState, useEffect } from "react";
import axios from "axios";

const ForexVals = () => {
  const [forexValues, setForexValues] = useState([]);

  useEffect(() => {
    // Fetch forex values from the database
    const fetchForexValues = async () => {
      try {
        const response = await axios.get("http://localhost:3002/query", {
          params: {
            channelid: "bankchannel",
            chaincodeid: "bank",
            function: "GetAllExchangeValues",
          },
        });
        console.log("Forex values:", response.data);
        setForexValues(JSON.parse(response.data.substring(9)));
      } catch (error) {
        console.error("Error fetching forex values:", error);
      }
    };

    fetchForexValues();
  }, []);

  return (
    <div>
      <h1>Foreign Exchange Values</h1>
      <table>
        <thead>
          <tr>
            <th>Original Currency</th>
            <th>Converted Currency</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
          {forexValues.map((forex) => (
            <tr key={forex.FromCurrency}>
              <td>{forex.FromCurrency}</td>
              <td>{forex.ToCurrency}</td>
              <td>{forex.Value}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ForexVals;
