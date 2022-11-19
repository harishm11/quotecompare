import { useState } from "react";
import React from "react";
export default function Vehicleform(props) {
  const [vehicleformFields, setvehicleformFields] = useState([
    {
      vehyear: 2001,
      vehmake: "",
      vehmodel: "",
      annualMileage: 10000,
      grgZip: "",
      vehicleusage: "",
    },
  ]);

  const handleVehcileFormChange = (event, index) => {
    let data = [...vehicleformFields];
    data[index][event.target.name] = event.target.value;
    setvehicleformFields(data);
    props.getvehicleData(data);
  };

  const addVehicleFields = () => {
    let data = {
      vehyear: 2001,
      vehmake: "",
      vehmodel: "",
      annualMileage: 10000,
      grgZip: "",
      vehicleusage: "",
    };
    setvehicleformFields([...vehicleformFields, data]);
    props.getvehicleData(data);
  };

  const removeVehicleFields = (index) => {
    let data = [...vehicleformFields];
    data.splice(index, 1);
    setvehicleformFields(data);
    props.getvehicleData(data);
  };

  return (
    <>
      <div className="card">
        <div className="card-body">
          <h3>Vehicles</h3>

          {vehicleformFields.map((form, index) => {
            return (
              <div key={index}>
                <label>Model Year</label>
                <input
                  className="form-control"
                  type="number"
                  min="1901"
                  name="vehyear"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.year}
                />
                <label>Make</label>
                <input
                  className="form-control"
                  name="vehmake"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.make}
                />
                <label>Model</label>
                <input
                  className="form-control"
                  name="vehmodel"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.model}
                />
                <label>Annual Mileage</label>
                <input
                  className="form-control"
                  type="number"
                  min="3000"
                  name="annualMileage"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.annualMileage}
                />
                <label>Garaging Zip Code</label>
                <input
                  className="form-control"
                  name="grgZip"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.grgZip}
                />
                <label>Vehicle Usage</label>
                <input
                  className="form-control"
                  name="vehicleusage"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.vehicleusage}
                />
                <button
                  className="btn btn-outline-success my-2 my-sm-0"
                  onClick={addVehicleFields}
                >
                  Add
                </button>
                <button
                  className="btn btn-outline-success my-2 my-sm-0"
                  onClick={() => removeVehicleFields(index)}
                >
                  Remove
                </button>
              </div>
            );
          })}
        </div>
      </div>
    </>
  );
}
