// import style
import './style/style.css'

import { Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import Daftar from './pages/Daftar';
import Hasil from './pages/Hasil';

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

        <Route
          path={"/hasil"}
          exact
          element={<Hasil/>}
        />

      </Routes>
    </div>

  </>);
}

export default App;
