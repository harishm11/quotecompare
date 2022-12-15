import { useState } from "react";
import "../App.css";
import Driverform from "../components/Driverform";
import Quoteform from "../components/Quoteform";
import Vehicleform from "../components/Vehicleform";
import { useNavigate } from "react-router-dom";
import React from "react";

function CreateQuote() {
  const [driverformFields, setdriverformFields] = useState();
  const [vehicleformFields, setvehicleformFields] = useState();
  const [quoteformFields, setquoteformFields] = useState();

  const getdriverData = (driverData) => {
    setdriverformFields(driverData);
  };
  const [showQuoteButton, setshowQuoteButton] = useState(true);
  const [showDriverButton, setshowDriverButton] = useState(false);
  const [showVehicleButton, setshowVehicleButton] = useState(false);
  const [showQuote, setshowQuote] = useState(false);
  const [showDriver, setshowDriver] = useState(false);
  const [showVehicle, setshowVehicle] = useState(false);

  const getvehicleData = (vehicleData) => {
    setvehicleformFields(vehicleData);
  };

  const getquoteData = (quoteData) => {
    setquoteformFields(quoteData);
  };

  const navigate = useNavigate();
  const handleSubmit = (event) => {
    event.preventDefault();
    console.time("start");
    const url = "http://localhost:8000/quoteApi/rating";
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        quoteformFields,
        driverformFields,
        vehicleformFields,
      }),
    };
    console.log(requestOptions.body);
    fetch(url, requestOptions)
      .then((response) => response.json())
      .then((data) => navigate("/Results", { state: { data } }))
      .then(console.time("end"))
      .catch((error) => console.log("Form submit error", error));
  };

  return (
    <div style={{ width: "600px", height: "600px" }} className="container">
      <div className="form-group col-md-45">
        {showQuoteButton && (
          <div className="form-group col-md-45">
            <input type="text" name="Zipcode" placeholder="Zipcode" />
            <button
              className=" btn btn-outline-dark my-2 my-sm-0"
              onClick={() => {
                setshowQuote(true);
                setshowQuoteButton(false);
                setshowDriverButton(true);
              }}
            >
              Create Quote
            </button>
          </div>
        )}
      </div>
      <div className="form-group col-md-45">
        {showQuote && <Quoteform getquoteData={getquoteData}></Quoteform>}
      </div>

      {showDriverButton && (
        <button
          className="btn btn-outline-dark my-2 my-sm-0"
          onClick={() => {
            setshowDriver(true);
            setshowDriverButton(false);
            setshowVehicleButton(true);
          }}
        >
          Add Driver Details
        </button>
      )}

      <div className="form-group col-md-45">
        {showDriver && <Driverform getdriverData={getdriverData}></Driverform>}
      </div>
      {showVehicleButton && (
        <button
          className="btn btn-outline-dark my-2 my-sm-0"
          onClick={() => {
            setshowVehicle(true);
            setshowVehicleButton(false);
          }}
        >
          Add Vehicle Details
        </button>
      )}
      <div className="form-group col-md-45">
        {showVehicle && (
          <Vehicleform getvehicleData={getvehicleData}></Vehicleform>
        )}
      </div>
      <div className="form-group col-md-6">
        {showQuote && showDriver && showVehicle && (
          <button
            className="btn btn-outline-dark my-2 my-sm-0"
            onClick={handleSubmit}
          >
            Quote
          </button>
        )}
      </div>
    </div>
  );
}

export default CreateQuote;
