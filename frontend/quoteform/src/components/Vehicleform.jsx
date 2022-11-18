import { useState } from "react";

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
                <input
                  className="form-control"
                  type="number"
                  min="1901"
                  name="vehyear"
                  placeholder="Year"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.year}
                />
                <input
                  className="form-control"
                  name="vehmake"
                  placeholder="Make"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.make}
                />
                <input
                  className="form-control"
                  name="vehmodel"
                  placeholder="Model"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.model}
                />
                <input
                  className="form-control"
                  type="number"
                  min="3000"
                  name="annualMileage"
                  placeholder="Annual Mileage"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.annualMileage}
                />
                <input
                  className="form-control"
                  name="grgZip"
                  placeholder="Garaging Zip Code"
                  onChange={(event) => handleVehcileFormChange(event, index)}
                  value={form.grgZip}
                />
                <input
                  className="form-control"
                  name="vehicleusage"
                  placeholder="Vehicle Usage"
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
