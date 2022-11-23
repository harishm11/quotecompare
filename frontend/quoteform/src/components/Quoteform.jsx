import { useState } from "react";
import React from "react";
export default function Quoteform(props) {
  var [quoteformFields, setquoteformFields] = useState({
    quotenumber: Math.floor(Math.random() * 999999999),
    effDate: new Date(),
    policyterm: "6",
    AutoUmbrellaInd: "",
    AutoHomeInd: "",
    AutoHomeLifeInd: "",
    AutoLifeInd: "",
    AutoRenterInd: "",
    AutoRenterLifeInd: "",
  });

  const handleQuoteFormChange = (event) => {
    let data = quoteformFields;
    data[event.target.name] = event.target.value;
    setquoteformFields(data);
    props.getquoteData(data);
  };

  return (
    <>
      <div className="card">
        <div className="card-body ">
          <form>
            <label>Quote Number</label>
            <input
              className="form-control form-control-sm"
              name="quotenumber"
              onChange={(event) => handleQuoteFormChange(event)}
              value={quoteformFields.quotenumber}
            />
            <label>Effective Date</label>
            <input
              className="form-control form-control-sm"
              name="effDate"
              onChange={(event) => handleQuoteFormChange(event)}
              value={quoteformFields.effDate}
            />
            <label>Policy term</label>
            <input
              className="form-control form-control-sm"
              type="number"
              min="6"
              name="policyterm"
              onChange={(event) => handleQuoteFormChange(event)}
              value={quoteformFields.policyterm}
            />
            {/* <label>
              Line of business:
                <select value={form.lob} onChange={event => handleQuoteFormChange(event, index)}>
                  <option value="Auto">Auto</option>
                  <option value="Home">Home</option>
                  <option value="Umbrella">Umbrella</option>
                </select>
              </label> */}
          </form>
        </div>
      </div>
    </>
  );
}
