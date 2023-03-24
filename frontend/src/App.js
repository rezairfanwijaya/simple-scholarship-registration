// import style
import './style/style.css'

import { Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import Daftar from './pages/Daftar';

function App() {
  return (<>
    <div className="App font-inter">
      <Routes>
        <Route
          path={"/"}
          exact
          element={<LandingPage/>}
        />

        <Route
          path={"/daftar"}
          exact
          element={<Daftar/>}
        />

      </Routes>
    </div>

  </>);
}

export default App;
