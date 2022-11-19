import { useLocation } from "react-router-dom";
// import axios from "axios";
// import fileDownload from "js-file-download";
import React from "react";
export default function ShowResults() {
  const location = useLocation();
  return (
    <>
      <ul className="list-group">
        <li className="list-group-item">
          <h3>Policy Premium = {location.state.data.Amount}</h3>
        </li>
      </ul>
      {/* <div>VehicleDetails = {location.state.data.VehDetails}</div> */}
      <ul className="list-group">
        {Object.entries(location.state.data.VehDetails).map(([key, veh], i) => (
          <li className="list-group-item" key={i}>
            <span className="input-label">
              Vehicle-{i + 1} Premium: {veh.Amount}
            </span>
            <ul className="list-group">
              {Object.entries(veh.CvgDetails).map(([key, cvg], i) => (
                <li className="list-group-item" key={i}>
                  <span className="input-label">
                    {cvg.CoverageCode} Coverage Premium: {cvg.Amount}
                  </span>
                </li>
              ))}
            </ul>
          </li>
        ))}
      </ul>
      <button
        onClick={() => {
          window.open("http://localhost:8000/download/output.csv");
          // this.handleDownload(
          //   "http://localhost:8000/download/output.csv",
          //   "output.csv"
          // );
        }}
      >
        Download Worksheet
      </button>
    </>
  );
}

// handleDownload = (url, filename) => {
//   axios
//     .get(url, {
//       responseType: "blob",
//     })
//     .then((res) => {
//       fileDownload(res.data, filename);
//     });
// };
