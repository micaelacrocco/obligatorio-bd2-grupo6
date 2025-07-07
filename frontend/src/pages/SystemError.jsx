import React from "react";
import { useNavigate } from "react-router-dom";
import {
    Box,
    Button,
    Typography,
    styled,
    GlobalStyles
} from "@mui/material";
import {
    Cancel as CancelIcon,
    Info as InfoIcon,
    CalendarToday as CalendarIcon,
    AccessTime as TimeIcon,
    LocationOn as LocationIcon,
    CheckCircle as CheckIcon
} from "@mui/icons-material";

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
                fontFamily: 'Arial, sans-serif',
            },
            '#root': {
                width: '100%',
                minHeight: '100%',
                display: 'flex',
                flexDirection: 'column',
            }
        }}
    />
);

const Header = styled(Box)({
    backgroundColor: '#3d56a6',
    color: 'white',
    padding: '15px 20px',
    fontSize: '1.1rem',
    fontWeight: 'bold',
});

const FloatingErrorAlert = styled(Box)({
    position: 'fixed',
    top: '75px',
    right: '20px',
    zIndex: 9999,
    backgroundColor: '#f8d7da',
    border: '1px solid #f5c6cb',
    borderLeft: '4px solid #dc3545',
    padding: '15px 20px',
    borderRadius: '4px',
    display: 'flex',
    alignItems: 'center',
    gap: '10px',
    boxShadow: '0 2px 8px rgba(0, 0, 0, 0.15)',
});

const MainContent = styled(Box)({
    flex: 1,
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    textAlign: 'center',
});

const InfoSection = styled(Box)({
    backgroundColor: 'white',
    border: '1px solid #e9ecef',
    borderRadius: '8px',
    padding: '20px',
    maxWidth: '600px',
    width: '100%',
});

const InfoItem = styled(Box)({
    display: 'flex',
    alignItems: 'center',
    gap: '12px',
    marginBottom: '12px',
    '&:last-child': {
        marginBottom: 0,
    },
});

const ButtonContainer = styled(Box)({
    display: 'flex',
    gap: '15px',
    marginTop: '25px',
    flexWrap: 'wrap',
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

const SystemError = () => {
    const navigate = useNavigate();

    const handleGoHome = () => {
        navigate("/");
    };

    const handleLogout = () => {
        localStorage.removeItem("token");
        navigate("/login");
    }

    return (
        <>
            {globalStyles}
            <Box sx={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
                <Header sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
                    Sistema Electoral
                    <Button
                        variant="outlined"
                        onClick={handleLogout}
                        sx={{
                            alignSelf: 'flex-end',
                            color: 'white',
                            borderColor: 'white',
                            height: '24px',
                            '&:hover': {
                                backgroundColor: 'rgba(255,255,255,0.1)',
                                borderColor: 'white',
                            },
                        }}
                    >
                        Salir
                    </Button>
                </Header>

                <MainContent>
                    <FloatingErrorAlert>
                        <CancelIcon sx={{ color: '#dc3545', fontSize: '24px' }} />
                        <Typography sx={{ color: '#721c24', fontWeight: 500 }}>
                            Ya ha emitido su voto en esta elección
                        </Typography>
                    </FloatingErrorAlert>

                    <CancelIcon sx={{ color: '#dc3545', fontSize: '80px', mb: 2 }} />

                    <Typography
                        variant="h4"
                        sx={{ fontWeight: 'bold', color: '#495057', mb: 2 }}
                    >
                        Voto Duplicado Detectado
                    </Typography>

                    <Typography
                        variant="h6"
                        sx={{ color: '#6c757d', mb: 1 }}
                    >
                        Su voto ya fue registrado el día 11/05/2025 a las 14:30
                    </Typography>

                    <Typography
                        variant="body1"
                        sx={{ color: '#6c757d', marginBottom: '35px' }}
                    >
                        Por seguridad, no es posible votar nuevamente.
                    </Typography>

                    <InfoSection>
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1, mb: 2 }}>
                            <InfoIcon sx={{ color: '#22559c' }} />
                            <Typography variant="h6" sx={{ fontWeight: 'bold', color: '#495057' }}>
                                Información de su voto:
                            </Typography>
                        </Box>

                        <InfoItem>
                            <CalendarIcon sx={{ color: '#22559c' }} />
                            <Typography>Fecha: 11/05/2025</Typography>
                        </InfoItem>

                        <InfoItem>
                            <TimeIcon sx={{ color: '#22559c' }} />
                            <Typography>Hora: 14:30</Typography>
                        </InfoItem>

                        <InfoItem>
                            <LocationIcon sx={{ color: '#22559c' }} />
                            <Typography>Circuito: Mesa 1234</Typography>
                        </InfoItem>

                        <InfoItem>
                            <CheckIcon sx={{ color: '#22559c' }} />
                            <Typography>Estado: Registrado correctamente</Typography>
                        </InfoItem>
                    </InfoSection>

                    <ButtonContainer>
                        <StyledButton
                            variant="contained"
                            onClick={handleGoHome}
                            sx={{
                                mt: 2,
                                mb: 0,
                                py: 1.5,
                                fontSize: '1rem',
                                fontWeight: 'bold',
                            }}
                        >
                            VOLVER AL INICIO
                        </StyledButton>
                    </ButtonContainer>
                </MainContent>
            </Box>
        </>
    );
};

export default SystemError;
