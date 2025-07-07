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
    Info as InfoIcon,
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

const FloatingAlert = styled(Box)({
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
    margin: '20px',
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
    marginTop: '30px',
    flexWrap: 'wrap',
    justifyContent: 'center',
});

const StyledButton = styled(Button)({
    padding: '12px 24px',
    fontSize: '1rem',
    fontWeight: 'bold',
    textTransform: 'none',
    minWidth: '180px',
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
                    Administracion Electoral
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

                </MainContent>
            </Box>
        </>
    );
};

export default ElectionResult;
