import { useLocation } from "react-router-dom";

export default function ShowResults() {
  const location = useLocation();
  return (
    <>
      <div className="d-flex justify-content-center p-3 mb-2 bg-success text-white">
        Premium = {location.state.data.Amount}
      </div>
      {/* <div>VehicleDetails = {location.state.data.VehDetails}</div> */}
      {Object.entries(location.state.data.VehDetails).map(([key, veh], i) => (
        <li className="travelcompany-input" key={i}>
          <span className="input-label">
            Vehicle{i + 1} Premium: {veh.Amount}
          </span>
          {Object.entries(veh.CvgDetails).map(([key, cvg], i) => (
            <li className="travelcompany-input" key={i}>
              <span className="input-label">
                cvg{i + 1} Premium: {cvg.Amount}
              </span>
            </li>
          ))}
        </li>
      ))}
    </>
  );
}
