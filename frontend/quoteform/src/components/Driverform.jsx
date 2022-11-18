import { useState } from "react";

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
        <div className="card-body">
          <h3>Drivers</h3>
          {driverformFields.map((form, index) => {
            return (
              <div key={index}>
                <label>Name</label>
                <input
                  className="form-control"
                  name="name"
                  placeholder="Name"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.name}
                />
                {/* <label>Age</label>
                <input
                  className="form-control"
                  type="number"
                  min="15"
                  name="age"
                  placeholder="Age"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.age}
                /> */}
                <label>Driving Experience</label>
                <input
                  className="form-control"
                  type="number"
                  name="experience"
                  placeholder="Experience"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.experience}
                />
                <label>Driving Course</label>
                <input
                  className="form-control"
                  name="course"
                  placeholder="Course"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.course}
                />
                <label>Inident Date</label>
                <input
                  className="form-control"
                  name="incidentdate"
                  placeholder="Incident Date"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.incidentdate}
                />
                <label>Incident Type</label>
                <input
                  className="form-control"
                  name="incidenttype"
                  placeholder="Incident Type"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.incidenttype}
                />
                <label>Marital status</label>
                <input
                  className="form-control"
                  name="maritalstatcode"
                  placeholder="Marital status code"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.maritalstatcode}
                />
                <label>License Issue Date</label>
                <input
                  className="form-control"
                  name="licissuedt"
                  placeholder="License Issue Date"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.licissuedt}
                />
                <label>Good Student?</label>
                <input
                  className="form-control"
                  name="goodstudent"
                  placeholder="Good Student Indicator"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.goodstudent}
                />
                <label>Date of Birth</label>
                <input
                  className="form-control"
                  name="dateofbirth"
                  placeholder="Date of Birth"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.dateofbirth}
                />
                {/* <label>Driver Added Date</label>
                <input
                  className="form-control"
                  name="drveraddeddt"
                  placeholder="Driver Added Date"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.drveraddeddt}
                /> */}
                <label>Occupation</label>
                <input
                  className="form-control"
                  name="occupation"
                  placeholder="Occupation"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.occupation}
                />
                <label>PNI?</label>
                <input
                  className="form-control"
                  name="pniind"
                  placeholder="PNI Indicator"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.pniind}
                />
                <label>Relationship to PNI</label>
                <input
                  className="form-control"
                  name="relationtopni"
                  placeholder="Relationship to PNI"
                  onChange={(event) => handleDriverFormChange(event, index)}
                  value={form.relationtopni}
                />
                <button
                  className="btn btn-outline-success my-2 my-sm-0"
                  onClick={addDriverFields}
                >
                  Add
                </button>
                <button
                  className="btn btn-outline-success my-2 my-sm-0"
                  onClick={() => removeDriverFields(index)}
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
