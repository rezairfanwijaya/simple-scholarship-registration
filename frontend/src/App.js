// import style
import './style/style.css'

import { Routes, Route } from 'react-router-dom';
import LandingPage from './pages/LandingPage';
import Daftar from './pages/Daftar';
import Hasil from './pages/Hasil';
import Login from './pages/Login';
import PrivateRoute from './utils/auth';
import Dashboard from './pages/Dashboard';
import Index from './components/dashboard/Index';

function App() {
  return (<>
    <div className="App font-inter">
      <Routes>
        <Route
          path={"/"}
          exact
          element={<LandingPage />}
        />

        <Route
          path={"/daftar"}
          exact
          element={<Daftar />}
        />

        <Route
          path={"/hasil"}
          exact
          element={<Hasil />}
        />

        <Route
          path={"/login"}
          exact
          element={<Login />}
        />

        <Route
          path={"/dashboard"}
          element={
            <PrivateRoute>
              <Dashboard />
            </PrivateRoute>
          }
        >

          <Route path={""} element={<Index />} />
          {/* <Route path={"alumni"} element={<Alumni />} /> */}
        </Route>

      </Routes>
    </div>

  </>);
}

export default App;
