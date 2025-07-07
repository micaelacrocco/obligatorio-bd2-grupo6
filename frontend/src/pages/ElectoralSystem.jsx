import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  Box,
  Button,
  Typography,
  styled,
  GlobalStyles,
  CircularProgress,
} from "@mui/material";
import { Info as InfoIcon } from "@mui/icons-material";

const globalStyles = (
  <GlobalStyles
    styles={{
      "*": {
        margin: 0,
        padding: 0,
        boxSizing: "border-box",
      },
      body: {
        backgroundColor: "#eff2ff",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        flexDirection: "column",
        height: "100vh",
        width: "100vw",
        overflow: "hidden",
        fontFamily: "Arial, sans-serif",
      },
      "#root": {
        width: "100%",
        minHeight: "100%",
        display: "flex",
        flexDirection: "column",
      },
    }}
  />
);

const Header = styled(Box)({
  backgroundColor: "#3d56a6",
  color: "white",
  padding: "15px 20px",
  fontSize: "1.1rem",
  fontWeight: "bold",
});

const MainContent = styled(Box)({
  flex: 1,
  display: "flex",
  flexDirection: "column",
  alignItems: "center",
  justifyContent: "center",
  textAlign: "center",
});

const InfoSection = styled(Box)({
  backgroundColor: "white",
  border: "1px solid #e9ecef",
  borderRadius: "8px",
  padding: "20px",
  margin: "20px",
  maxWidth: "600px",
  width: "100%",
});

const InfoItem = styled(Box)({
  display: "flex",
  alignItems: "center",
  gap: "12px",
  marginBottom: "12px",
  "&:last-child": {
    marginBottom: 0,
  },
});

const ButtonContainer = styled(Box)({
  display: "flex",
  gap: "15px",
  marginTop: "30px",
  flexWrap: "wrap",
  justifyContent: "center",
});

const StyledButton = styled(Button)({
  padding: "12px 24px",
  fontSize: "1rem",
  fontWeight: "bold",
  textTransform: "none",
  minWidth: "180px",
});

const formatDate = (dateStr) => {
  if (!dateStr) return "";
  const d = new Date(dateStr);
  return d.toLocaleDateString();
};

const ElectoralSystem = () => {
  const navigate = useNavigate();
  const [citizen, setCitizen] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const handleVoteNow = () => {
    navigate("/ongoing-voting");
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    localStorage.removeItem("ci");
    navigate("/login");
  };

  useEffect(() => {
    const fetchCitizen = async () => {
      const ci = localStorage.getItem("ci");

      if (!ci) {
        setError("Sesión inválida: CI no encontrado.");
        setLoading(false);
        return;
      }

      try {
        const res = await fetch(`http://localhost:8080/citizens/${ci}`);

        if (!res.ok) {
          throw new Error("No se pudo obtener la información del votante");
        }

        const data = await res.json();
        setCitizen(data);
      } catch (err) {
        setError(err.message || "Error al obtener datos del ciudadano");
      } finally {
        setLoading(false);
      }
    };

    fetchCitizen();
  }, []);

  return (
    <>
      {globalStyles}
      <Box sx={{ minHeight: "100vh", display: "flex", flexDirection: "column" }}>
        <Header
          sx={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}
        >
          Sistema Electoral
          <Button
            variant="outlined"
            onClick={handleLogout}
            sx={{
              alignSelf: "flex-end",
              color: "white",
              borderColor: "white",
              height: "24px",
              "&:hover": {
                backgroundColor: "rgba(255,255,255,0.1)",
                borderColor: "white",
              },
            }}
          >
            Salir
          </Button>
        </Header>

        <MainContent>
          {loading ? (
            <CircularProgress />
          ) : error ? (
            <Typography sx={{ color: "red", fontWeight: "bold" }}>{error}</Typography>
          ) : (
            <InfoSection>
              <Box sx={{ display: "flex", alignItems: "center", gap: 1, mb: 2 }}>
                <InfoIcon sx={{ color: "#22559c" }} />
                <Typography variant="h6" sx={{ fontWeight: "bold", color: "#495057" }}>
                  Información del Votante:
                </Typography>
              </Box>

              <InfoItem>
                <Typography>
                  <strong>CI:</strong> {citizen.id}
                </Typography>
              </InfoItem>
              <InfoItem>
                <Typography>
                  <strong>Nombre:</strong> {citizen.first_name} {citizen.last_name}
                </Typography>
              </InfoItem>
              <InfoItem>
                <Typography>
                  <strong>Credencial:</strong> {citizen.credential}
                </Typography>
              </InfoItem>
              <InfoItem>
                <Typography>
                  <strong>Fecha de nacimiento:</strong> {formatDate(citizen.birth_date)}
                </Typography>
              </InfoItem>
            </InfoSection>
          )}

          <ButtonContainer>
            <StyledButton
              variant="contained"
              onClick={handleVoteNow}
              sx={{
                backgroundColor: "#3d56a6",
                "&:hover": {
                  backgroundColor: "#2c3e50",
                },
              }}
              disabled={loading || error}
            >
              VOTAR AHORA
            </StyledButton>
            <StyledButton
              variant="contained"
              onClick={handleLogout}
              sx={{
                backgroundColor: "#6c757d",
                "&:hover": {
                  backgroundColor: "#5a6268",
                },
              }}
            >
              SALIR
            </StyledButton>
          </ButtonContainer>
        </MainContent>
      </Box>
    </>
  );
};

export default ElectoralSystem;
