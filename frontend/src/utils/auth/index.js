import { Navigate } from "react-router-dom";

function PrivateRoute({ children }) {
    // ambil token di local storage
    let authUser = localStorage.getItem("token")

    // jika tidak punya token maka navigate ke halaman login
    if (!authUser) {
        return <Navigate to="/login" />
    }

    return children
}

export default PrivateRoute