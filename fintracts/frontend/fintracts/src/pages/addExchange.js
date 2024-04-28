import React, { useState } from "react";
import axios from "axios";

const AddExchange = () => {
  const [exchangeData, setExchangeData] = useState({
    currencyFrom: "",
    currencyTo: "",
    rate: "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setExchangeData({
      ...exchangeData,
      [name]: value,
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    // Perform API call to send exchangeData to the database
    // You can use libraries like axios to make the API call
    // Example: axios.post('/api/exchange', exchangeData)
    // Handle the response and any errors accordingly
    // Reset the form after successful submission
    console.log(exchangeData);
    try {
      const response = await axios.post(
        "http://localhost:3002/invoke",
        new URLSearchParams([
          ["", ""],
          ["channelid", "bankchannel"],
          ["chaincodeid", "bank"],
          ["function", "AddExchangeValue"],
          ["args", exchangeData.currencyFrom],
          ["args", exchangeData.currencyTo],
          ["args", exchangeData.rate],
        ])
      );
      console.log("Response:", response);
    } catch (error) {
      alert("Error adding exchange rate");
      console.error("Error:", error);
    }
    alert("Exchange rate added successfully");
    setExchangeData({
      currencyFrom: "",
      currencyTo: "",
      rate: "",
    });
  };

  return (
    <div>
      <h1>Add Exchange Rate</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="currency">Original Currency:</label>
          <input
            type="text"
            id="currency"
            name="currencyFrom"
            value={exchangeData.currencyFrom}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="currency">Final Currency:</label>
          <input
            type="text"
            id="currency"
            name="currencyTo"
            value={exchangeData.currencyTo}
            onChange={handleChange}
          />
        </div>
        <div>
          <label htmlFor="rate">Rate:</label>
          <input
            type="number"
            id="rate"
            name="rate"
            value={exchangeData.rate}
            onChange={handleChange}
          />
        </div>
        <button type="submit">Add Exchange Rate</button>
      </form>
    </div>
  );
};

export default AddExchange;
