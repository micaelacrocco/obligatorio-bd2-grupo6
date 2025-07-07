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

const Alert = styled(Box)({
    width: '600px',
    backgroundColor: '#d4edda',
    border: '1px solid #c3e6cb',
    borderLeft: '4px solid #28a745',
    padding: '15px 20px',
    borderRadius: '4px',
    display: 'flex',
    alignItems: 'center',
    gap: '10px',
    marginBottom: '20px',
});

const MainContent = styled(Box)({
    flex: 1,
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    justifyContent: 'center',
    textAlign: 'center',
});

const ElectionResult = () => {
    const navigate = useNavigate();

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
                    <Alert>
                        <CheckIcon sx={{ color: '#28a745', fontSize: '24px' }} />
                        <Typography sx={{ color: '#155724', fontWeight: 500 }}>
                            Mesa cerrada - Resultados finales disponibles
                        </Typography>
                    </Alert>
                    <Box
                        sx={{
                            display: 'flex',
                            flexDirection: 'column',
                            alignItems: 'start',
                            width: '600px',
                        }}
                    >
                        <Typography
                            component="h1"
                            variant="h5"
                            align="start"
                            gutterBottom
                            sx={{ fontWeight: 'bold', color: '#040111' }}
                        >
                            Resultado de la Elección
                        </Typography>
                        <Box>
                            {/* Aquí el contenido de resultados */}
                        </Box>
                    </Box>
                </MainContent>
            </Box>
        </>
    );
};

export default ElectionResult;
