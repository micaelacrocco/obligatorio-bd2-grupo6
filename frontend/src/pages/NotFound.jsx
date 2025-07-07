import React from "react";
import { useNavigate } from "react-router-dom";
import {
  Box,
  Button,
  Typography,
  styled,
  GlobalStyles
} from "@mui/material";
import { ReactComponent as LogoCorteElectoral } from "../assets/logo-corte-electoral.svg";

const globalStyles = (
  <GlobalStyles
    styles={{
      '*': {
        margin: 0,
        padding: 0,
        boxSizing: 'border-box',
      },
      body: {
        backgroundColor: '#eff2ff',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        flexDirection: 'column',
        height: '100vh',
        width: '100vw',
        overflow: 'hidden',
      },
      '#root': {
        width: '100%',
        height: '100%',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        flexDirection: 'column',
      }
    }}
  />
);

const StyledCard = styled(Box)({
  backgroundColor: '#fff',
  borderRadius: '8px',
  boxShadow: '0 2px 10px rgba(0, 0, 0, 0.1)',
  padding: '20px',
  width: '100%',
  maxWidth: '400px',
  margin: 0,
  textAlign: 'center',
});

const LogoContainer = styled(Box)({
  position: 'absolute',
  top: '30px',
  left: '30px',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
});

const StyledButton = styled(Button)({
  backgroundColor: '#3d56a6',
  color: 'white',
  '&:hover:not(:disabled)': {
    backgroundColor: '#2c3e50',
  },
  '&:disabled': {
    opacity: 0.6,
    cursor: 'not-allowed',
  },
});


const ErrorNumber = styled(Typography)({
  fontSize: '8rem',
  fontWeight: 'bold',
  color: '#040111',
  marginBottom: '20px',
  lineHeight: 1,
  opacity: 0.1,
});

const NotFound = () => {
  const navigate = useNavigate();

  const handleGoHome = () => {
    navigate("/");
  };

  return (
    <>
      {globalStyles}
      <LogoContainer>
        <LogoCorteElectoral/>
      </LogoContainer>

      <Box
        sx={{
          width: '100%',
          height: '100%',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
          position: 'relative',
        }}
      >
        <ErrorNumber>404</ErrorNumber>
        
        <Box sx={{ textAlign: 'center', maxWidth: '400px', px: 2 }}>
          <Typography
            component="h1"
            variant="h4"
            align="center"
            gutterBottom
            sx={{ fontWeight: 'bold', color: '#040111', mb: 1.5 }}
          >
            Página no encontrada
          </Typography>
          
          <Typography
            variant="body1"
            align="center"
            sx={{ fontWeight: 500, color: '#636365', mb: 2 }}
          >
            Lo sentimos, la página que está buscando no existe o ha sido movida.
          </Typography>

          <StyledButton
            variant="contained"
            onClick={handleGoHome}
            sx={{
              mt: 2,
              mb: 0,
              py: 1.5,
              px: 4,
              fontSize: '1rem',
              fontWeight: 'bold',
            }}
          >
            Volver al Inicio
          </StyledButton>
        </Box>
      </Box>
    </>
  );
};

export default NotFound;