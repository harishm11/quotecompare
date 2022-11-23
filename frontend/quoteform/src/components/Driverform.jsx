import { useState } from "react";
import React from "react";
export default function Driverform(props) {
  const [driverformFields, setdriverformFields] = useState([
    {
      name: "",
      age: 0,
      experience: 0,
      course: "",
      incidentdate: "",
      incidenttype: "",
      maritalstatcode: "",
      licissuedt: new Date(),
      goodstudent: "",
      dateofbirth: new Date(),
      drveraddeddt: new Date(),
      occupation: "",
      pniind: "",
      relationtopni: "",
    },
  ]);

  const handleDriverFormChange = (event, index) => {
    let data = [...driverformFields];
    data[index][event.target.name] = event.target.value;
    setdriverformFields(data);
    props.getdriverData(data);
  };

  const addDriverFields = () => {
    let data = {
      name: "",
      age: 0,
      experience: 0,
      course: "",
      incidentdate: "",
      incidenttype: "",
      maritalstatcode: "",
      licissuedt: new Date(),
      goodstudent: "",
      dateofbirth: new Date(),
      drveraddeddt: new Date(),
      occupation: "",
      pniind: "",
      relationtopni: "",
    };
    setdriverformFields([...driverformFields, data]);
    props.getdriverData(data);
  };

  const removeDriverFields = (index) => {
    let data = [...driverformFields];
    data.splice(index, 1);
    setdriverformFields(data);
    props.getdriverData(data);
  };

  return (
    <>
      <div className="card">
        {driverformFields.map((form, index) => {
          return (
            <div className="card-body " key={index}>
              <h4>
                <label>Driver {index + 1}</label>
              </h4>

              <label>Name</label>
              <input
                className="form-control form-control-sm"
                name="name"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.name}
                required
              />
              {/* <label>Age</label>
                <input
                     className="form-control form-control-sm"
                  type="number"
                  min="15"
                  name="age"
                  placeholder="Age"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.age}
                /> */}
              <label>Driving Experience</label>
              <input
                className="form-control form-control-sm"
                type="number"
                name="experience"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.experience}
              />
              <label>Driving Course</label>
              <input
                className="form-control form-control-sm"
                name="course"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.course}
              />
              <label>Inident Date</label>
              <input
                className="form-control form-control-sm"
                name="incidentdate"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.incidentdate}
              />
              <label>Incident Type</label>
              <input
                className="form-control form-control-sm"
                name="incidenttype"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.incidenttype}
              />
              <label>Marital status</label>
              <input
                className="form-control form-control-sm"
                name="maritalstatcode"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.maritalstatcode}
                required
              />
              <label>License Issue Date</label>
              <input
                className="form-control form-control-sm"
                name="licissuedt"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.licissuedt}
              />
              <label>Good Student?</label>
              <input
                className="form-control form-control-sm"
                name="goodstudent"
                placeholder="Good Student Indicator"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.goodstudent}
              />
              <label>Date of Birth</label>
              <input
                className="form-control form-control-sm"
                name="dateofbirth"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.dateofbirth}
                required
              />
              {/* <label>Driver Added Date</label>
                <input
                     className="form-control form-control-sm"
                  name="drveraddeddt"
                  placeholder="Driver Added Date"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.drveraddeddt}
                /> */}
              <label>Occupation</label>
              <input
                className="form-control form-control-sm"
                name="occupation"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.occupation}
              />
              <label>PNI?</label>
              <input
                className="form-control form-control-sm"
                name="pniind"
                placeholder="PNI Indicator"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.pniind}
              />
              <label>Relationship to PNI</label>
              <input
                className="form-control form-control-sm"
                name="relationtopni"
                onChange={(event) => handleDriverFormChange(event, index)}
                value={form.relationtopni}
              />
              <button
                className="btn btn-outline-dark my-2 my-sm-0"
                onClick={() => removeDriverFields(index)}
              >
                Remove Driver
              </button>
            </div>
          );
        })}
      </div>

      <button
        className="btn btn-outline-dark my-2 my-sm-0"
        onClick={addDriverFields}
      >
        Add Driver
      </button>
    </>
  );
}
