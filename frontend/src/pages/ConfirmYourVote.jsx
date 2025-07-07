import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import {
    Box,
    Button,
    Typography,
    styled,
    GlobalStyles,
    TextField 
} from "@mui/material";
import {
    Info as InfoIcon,
    Warning as WarningIcon
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
    backgroundColor: "#fff3cd",
    border: "1px solid #ffeeba",
    borderLeft: "4px solid #f0ad4e",
    padding: "15px 20px",
    borderRadius: "4px",
    display: "flex",
    alignItems: "center",
    gap: "10px",
    marginBottom: "25px",
    maxWidth: "700px",
    width: "100%",
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
    margin: '10px',
    maxWidth: '700px',
    width: '700px',
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

const ConfirmYourVote = () => {
    const navigate = useNavigate();
    const [voteData, setVoteData] = useState(null);
    const [circuitNumber, setCircuitNumber] = useState(""); 

    useEffect(() => {
        const storedVote = localStorage.getItem("voteData");
        if (storedVote) {
            setVoteData(JSON.parse(storedVote));
        }
    }, []);

    if (!voteData) return <p>No hay voto para mostrar.</p>;

    const selectedListNumber = voteData.listasSeleccionadas.length > 0
        ? voteData.listasSeleccionadas[0]
        : null;

    const handleConfirmVote = async () => {
        const storedVote = localStorage.getItem("voteData");
        if (!storedVote) {
            alert("No hay voto para enviar");
            return;
        }
        const voteData = JSON.parse(storedVote);

        if (voteData.votoEnBlanco || voteData.votoAnulado) {
            navigate("/electoral-system");
            return;
        }

        const circuitId = parseInt(circuitNumber); 
        if (isNaN(circuitId)) {
            alert("Por favor ingrese un número de circuito válido.");
            return;
        }

        const payload = {
            vote_date: new Date().toISOString().split("T")[0],
            list_number: voteData.listasSeleccionadas[0],
            circuit_id: circuitId,
        };

        const token = localStorage.getItem("token");

        try {
            const response = await fetch("http://localhost:8080/list-votes", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    ...(token && { Authorization: `Bearer ${token}` }),
                },
                body: JSON.stringify(payload),
            });

            if (!response.ok) {
                throw new Error("Error al enviar el voto");
            }

            navigate("/electoral-system");
        } catch (error) {
            console.error("Hubo un error al registrar el voto:", error);
            alert("Ocurrió un error al registrar el voto. Intenta nuevamente.");
        }
    };

    const handleModifyVote = () => {
        navigate("/ongoing-voting");
    };

    const handleLogout = () => {
        localStorage.removeItem("token");
        navigate("/login");
    };

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
                        <WarningIcon sx={{ color: '#f0ad4e', fontSize: '24px' }} />
                        <Typography sx={{ color: '#856404', fontWeight: 500 }}>
                            Una vez confirmado, no podrá modificar su voto
                        </Typography>
                    </Alert>

                    <InfoSection>
                        <Box sx={{ display: 'flex', alignItems: 'center', gap: 1, mb: 2 }}>
                            <InfoIcon sx={{ color: '#22559c' }} />
                            <Typography variant="h6" sx={{ fontWeight: 'bold', color: '#495057' }}>
                                Resumen de su voto:
                            </Typography>
                        </Box>

                        <InfoItem>
                            <strong>Voto Anulado:</strong> {voteData.votoAnulado ? "Sí" : "No"}
                        </InfoItem>

                        {voteData.votoAnulado ? (
                            <>
                                <InfoItem>
                                    <strong>Voto Observado:</strong> {voteData.votoObservado ? "Sí" : "No"}
                                </InfoItem>
                                <InfoItem>
                                    <strong>Voto en Blanco:</strong> No
                                </InfoItem>
                            </>
                        ) : (
                            <>
                                <InfoItem>
                                    <strong>Voto en Blanco:</strong> {voteData.votoEnBlanco ? "Sí" : "No"}
                                </InfoItem>

                                {!voteData.votoEnBlanco && (
                                    <>
                                        <InfoItem>
                                            <strong>Lista seleccionada:</strong> {selectedListNumber}
                                        </InfoItem>
                                        <InfoItem>
                                            <strong>Nombre del partido:</strong> {voteData.partyName || "Desconocido"}
                                        </InfoItem>
                                    </>
                                )}

                                <InfoItem>
                                    <strong>Voto Observado:</strong> {voteData.votoObservado ? "Sí" : "No"}
                                </InfoItem>
                            </>
                        )}

                        {/* 🔧 NUEVO: Campo para número de circuito */}
                        <InfoItem sx={{ flexDirection: "column", alignItems: "start", mt: 2 }}>
                            <Typography fontWeight="bold">Número de circuito:</Typography>
                            <TextField
                                variant="outlined"
                                size="small"
                                value={circuitNumber}
                                onChange={(e) => setCircuitNumber(e.target.value)}
                                placeholder="Ej. 101"
                                sx={{ mt: 1, width: '200px' }}
                            />
                        </InfoItem>
                    </InfoSection>

                    <ButtonContainer>
                        <StyledButton
                            variant="contained"
                            onClick={handleConfirmVote}
                            sx={{
                                backgroundColor: '#3d56a6',
                                '&:hover': {
                                    backgroundColor: '#2c3e50',
                                },
                            }}
                        >
                            CONFIRMAR VOTO
                        </StyledButton>
                        <StyledButton
                            variant="contained"
                            onClick={handleModifyVote}
                            sx={{
                                backgroundColor: '#6c757d',
                                '&:hover': {
                                    backgroundColor: '#5a6268',
                                },
                            }}
                        >
                            MODIFICAR VOTO
                        </StyledButton>
                    </ButtonContainer>
                </MainContent>
            </Box>
        </>
    );
};

export default ConfirmYourVote;