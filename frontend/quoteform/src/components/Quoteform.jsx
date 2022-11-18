import { useState } from "react";

export default function Quoteform(props) {
  var [quoteformFields, setquoteformFields] = useState({
    quotenumber: Math.floor(Math.random() * 999999999),
    effDate: new Date(),
    policyterm: 12,
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
        <div className="card-body">
          <form>
            <input
              className="form-control"
              name="quotenumber"
              placeholder="Quote Number"
              onChange={(event) => handleQuoteFormChange(event)}
              value={quoteformFields.quotenumber}
            />
            <input
              className="form-control"
              name="effDate"
              placeholder="Effective Date"
              onChange={(event) => handleQuoteFormChange(event)}
              value={quoteformFields.effDate}
            />
            <input
              className="form-control"
              type="number"
              min="6"
              name="policyterm"
              placeholder="Policy term"
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
