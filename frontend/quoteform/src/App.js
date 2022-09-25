import { useState } from 'react';
import React from 'react';
// import DatePicker from "react-datepicker";
import './App.css';

function App() {
  const [quoteformFields, setquoteformFields] = useState(
    [{ quotenumber :Math.floor(Math.random() * 999999999) ,effDate:'',policyterm:6}]
  )

  
  var quotenumber = quoteformFields[0].quotenumber
  const effDate = new Date(quoteformFields[0].effDate)
  const policyterm = 6

  

  const [driverformFields, setdriverformFields] = useState([
    { name: '', age: 0 ,experience:0, course: '',incidentdate:'', incidenttype:'',maritalstatcode:''},
  ])
  const [vehicleformFields, setvehicleformFields] = useState([
    { vehyear:1987, vehmake: '' ,vehmodel: '', annualMileage: 10000,grgZip:'',vehicleusage:''},
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
    
    const url = 'http://localhost:8000/quoteApi/quote'
    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({quotenumber,effDate,policyterm, driverformFields,vehicleformFields})
    };
   
    fetch(url, requestOptions)
        .then(response =>response.json())
        .then(data => alert("premium = " + data))
        //.then(window.location.reload(true))
        .catch(error => console.log('Form submit error', error))

  };
  
  const addDriverFields = () => {
    let object = {
      name: '', age: 0 ,experience:0, course: '',incidentdate:'', incidenttype:'',maritalstatcode:''
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
       vehyear:1987, vehmake: '' ,vehmodel: '', annualMileage: 10000,grgZip:'',vehicleusage:''
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