import { Outlet } from "react-router-dom";

const ProtectedLayout = () => {
    return (
        <div style={{ display: "flex", height: "100vh", overflowX: "hidden" }}>
            <div style={{ flex: 1, background: "#fff", overflowY: "auto" }}>
                <Outlet />
            </div>
        </div>
    );
};

export default ProtectedLayout;
