import { useEffect, useState } from "react";
import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ children }) => {
    const [token, setToken] = useState(localStorage.getItem("token"));

    useEffect(() => {
        const handleStorageChange = () => {
            setToken(localStorage.getItem("token"));
        };

        window.addEventListener("storage", handleStorageChange);

        return () => {
            window.removeEventListener("storage", handleStorageChange);
        };
    }, []);

    if (!token) {
        return <Navigate to="/login" replace />;
    }

    return children;
};

export default ProtectedRoute;
