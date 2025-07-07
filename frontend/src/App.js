import React from "react";
import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import ProtectedRoute from "./routes/ProtectedRoute";
import ConfirmYourVote from "./pages/ConfirmYourVote";
import ElectionAdministration from "./pages/ElectionAdministration";
import ElectionResult from "./pages/ElectionResult";
import ElectoralSystem from "./pages/ElectoralSystem";
import Login from "./pages/Login";
import NotFound from "./pages/NotFound";
import OngoingVoting from "./pages/OngoingVoting";
import SystemError from "./pages/SystemError";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/login" replace />} />
        <Route path="/login" element={<Login />} />
        <Route
          path="/electoral-system"
          element={
            <ProtectedRoute>
              <ElectoralSystem />
            </ProtectedRoute>
          }
        />
        <Route
          path="/ongoing-voting"
          element={
            <ProtectedRoute>
              <OngoingVoting />
            </ProtectedRoute>
          }
        />
        <Route
          path="/confirm-your-vote"
          element={
            <ProtectedRoute>
              <ConfirmYourVote />
            </ProtectedRoute>
          }
        />
        <Route
          path="/election-administration"
          element={
            <ProtectedRoute>
              <ElectionAdministration />
            </ProtectedRoute>
          }
        />
        <Route
          path="/election-results"
          element={
            <ProtectedRoute>
              <ElectionResult />
            </ProtectedRoute>
          }
        />
        <Route path="/system-error" element={<SystemError />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
