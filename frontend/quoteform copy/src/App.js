import CreateQuote from "./pages/CreateQuote";
import { BrowserRouter as Router, Routes, Route , Link} from "react-router-dom";
import ShowResults from "./pages/ShowResults";
import React  from 'react';

export default function App() {


  return (

    <Router>
      <div className="App">
        <div className="navbar navbar-expand-lg navbar-dark bg-dark">  
           <ul className="navbar-nav mr-sm-2">
              <Link className="nav-link" to="/">Home</Link>
              <Link className="nav-link" to="/Quote">Quote</Link>
              <Link className="nav-link" to="/new">New</Link>             
           </ul>
        <div className="logo">QuoteCompare</div>
        </div>
        <Routes>
            <Route path="/" element={App} />
            <Route path="/Quote" element={<CreateQuote/>} />
            <Route path="/Results" element={<ShowResults/>} /> 
            {/* <Route path="/new"  element={   <Signup/>} />  */}
        </Routes>
          
      </div>
     </Router>
  );
}