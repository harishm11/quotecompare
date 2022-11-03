import { useState } from 'react';
import React from 'react';
// import DatePicker from "react-datepicker";
import './App.css';

function App() {
  const [quoteformFields, setquoteformFields] = useState(
    [{ quotenumber :Math.floor(Math.random() * 999999999) ,effDate:'',policyterm:6
  ,AutoUmbrellaInd :'',
	AutoHomeInd       :'',
	AutoHomeLifeInd  :'',
	AutoLifeInd       :'',
	AutoRenterInd     :'',
	AutoRenterLifeInd :''}]
  )

  
  var quotenumber = quoteformFields[0].quotenumber
  const effDate = new Date(quoteformFields[0].effDate)
  const policyterm = 6

  

  const [driverformFields, setdriverformFields] = useState([
    { name: '', age: 0 ,experience:0, course: '',incidentdate:'', incidenttype:'',maritalstatcode:''
    ,licissuedt:new Date(),goodstudent:'',dateofbirth:new Date(),drveraddeddt:new Date(),occupation:'',pniind:'',relationtopni:''},
  ])
  
  const [vehicleformFields, setvehicleformFields] = useState([
    { vehyear:2001, vehmake: '' ,vehmodel: '', annualMileage: 10000,grgZip:'',vehicleusage:''},
  ])

  
  const handleQuoteFormChange = (event, index) => {
    let data = [...quoteformFields];
    data[index][event.target.name] = event.target.value;
    setquoteformFields(data);
    
  }

  const handleDriverFormChange = (event, index) => {
    let data = [...driverformFields];
    data[index][event.target.name] = event.target.value;
    setdriverformFields(data);
  }

  const handleVehcileFormChange = (event, index) => {
    let data = [...vehicleformFields];
    data[index][event.target.name] = event.target.value;
    setvehicleformFields(data);
  }
  const handleSubmit = event => {
    event.preventDefault();
    console.log("from submit")
    console.log(new Date().toISOString())
    const url = 'http://localhost:8000/quoteApi/rating'
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({quotenumber,effDate,policyterm, driverformFields,vehicleformFields})
    };
    console.log(requestOptions.body)
    fetch(url, requestOptions)
        .then(response =>response.json())
        //.then(data => console.log(data) )
        //.then(data => alert("premium = " + data))
        .then(console.log("After response"))
        .then(data => console.log(data))
        //.then(window.location.reload(true))
        .then(console.log(new Date().toISOString()))
        .catch(error => console.log('Form submit error', error))

  };
  
  const addDriverFields = () => {
    let object = {
      name: '', age: 0 ,experience:0, course: '',incidentdate:'', incidenttype:'',maritalstatcode:''
    ,licissuedt:new Date(),goodstudent:'',dateofbirth:new Date(),drveraddeddt:new Date(),occupation:'',pniind:'',relationtopni:''
    }
    setdriverformFields([...driverformFields, object])
  }

  const removeDriverFields = (index) => {
    let data = [...driverformFields];
    data.splice(index, 1)
    setdriverformFields(data)
  }

  const addVehicleFields = () => {
    let object = {
       vehyear:2001, vehmake: '' ,vehmodel: '', annualMileage: 10000,grgZip:'',vehicleusage:''
    }
    setvehicleformFields([...vehicleformFields, object])
  }

  const removeVehicleFields = (index) => {
    let data = [...vehicleformFields];
    data.splice(index, 1)
    setvehicleformFields(data)
  }
  return (
    <div className="App">
      {quoteformFields.map((form,index) => {
          return (
            <div key={index}>
              <div>
                <h1>Quote</h1>
              </div>
              <input
                name='quotenumber'
                placeholder='Quote Number'
                onChange={event => handleQuoteFormChange(event, index)}
                value={form.quotenumber}
              />   
              <input
                name='effDate'
                placeholder='Effective Date'
                onChange={event => handleQuoteFormChange(event, index)}
                value={form.effDate}
              />
              <input
                type='number'
                min ='6'
                name='policyterm'
                placeholder='Policy term'
                onChange={event => handleQuoteFormChange(event, index)}
                value={form.policyterm}
              />
              {/* <label>
              Line of business:
                <select value={form.lob} onChange={event => handleQuoteFormChange(event, index)}>
                  <option value="Auto">Auto</option>
                  <option value="Home">Home</option>
                  <option value="Umbrella">Umbrella</option>
                </select>
              </label> */}
            </div>
          )
        })}
        <div>
                <h1>Drivers</h1>
        </div>
        {driverformFields.map((form, index) => {
          return (
            
            <div key={index}>
              
              <input
                name='name'
                placeholder='Name'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.name}
              />
              <input
                type='number'
                min = '15'
                name='age'
                placeholder='Age'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.age}
              />
              <input
                type='number'
                name='experience'
                placeholder='Experience'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.experience}
              />
              <input
                name='course'
                placeholder='Course'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.course}
              />
              <input
                name='incidentdate'
                placeholder='Incident Date'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.incidentdate}
              />
              <input
                name='incidenttype'
                placeholder='Incident Type'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.incidenttype}
              />
              <input
                name='maritalstatcode'
                placeholder='Marital status code'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.maritalstatcode}
              />
              <input
                name='licissuedt'
                placeholder='License Issue Date'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.licissuedt}
              />
              <input
                name='goodstudent'
                placeholder='Good Student Indicator'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.goodstudent}
              />
              <input
                name='dateofbirth'
                placeholder='Date of Birth'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.dateofbirth}
              />
              <input
                name='drveraddeddt'
                placeholder='Driver Added Date'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.drveraddeddt}
              />
              <input
                name='occupation'
                placeholder='Occupation'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.occupation}
              />
              <input
                name='pniind'
                placeholder='PNI Indicator'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.pniind}
              />
              <input
                name='relationtopni'
                placeholder='Relationship to PNI'
                onChange={event => handleDriverFormChange(event, index)}
                value={form.relationtopni}
              />
              <button onClick={addDriverFields}>Add</button>
              <button onClick={() => removeDriverFields(index)}>Remove</button>
            </div>
          )
        })}


        <div>
                <h1>Vehicles</h1>
        </div>
        {vehicleformFields.map((form, index) => {
          return (
            
            <div key={index}>
              
              <input
                type='number'
                min = '1901'
                name='vehyear'
                placeholder='Year'
                onChange={event => handleVehcileFormChange(event, index)}
                value={form.year}
              />
              <input
                name='vehmake'
                placeholder='Make'
                onChange={event => handleVehcileFormChange(event, index)}
                value={form.make}
              />
              <input
                name='vehmodel'
                placeholder='Model'
                onChange={event => handleVehcileFormChange(event, index)}
                value={form.model}
              />
              <input
                type='number'
                min='3000'
                name='annualMileage'
                placeholder='Annual Mileage'
                onChange={event => handleVehcileFormChange(event, index)}
                value={form.annualMileage}
              />
              <input
                name='grgZip'
                placeholder='Garaging Zip Code'
                onChange={event => handleVehcileFormChange(event, index)}
                value={form.grgZip}
              />
              <input
                name='vehicleusage'
                placeholder='Vehicle Usage'
                onChange={event => handleVehcileFormChange(event, index)}
                value={form.vehicleusage}
              />
              <button onClick={addVehicleFields}>Add</button>
              <button onClick={() => removeVehicleFields(index)}>Remove</button>
            </div>
          )
        })}
      {/* </form> */}
      <br />
      <button onClick={handleSubmit}>Submit</button>
    </div>
  );
}

export default App;