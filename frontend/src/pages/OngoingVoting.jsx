import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import {
    Box,
    Button,
    Typography,
    styled,
    GlobalStyles,
    Switch,
    FormControlLabel,
    Container,
} from "@mui/material";
import { Warning as WarningIcon } from "@mui/icons-material";

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
                fontFamily: "Arial, sans-serif",
                minHeight: "100vh",
            },
            "#root": {
                width: "100%",
                minHeight: "100vh",
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
    position: "sticky",
    top: 0,
    zIndex: 1000,
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

const MainContent = styled(Container)({
    flex: 1,
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    padding: "30px 20px",
    maxWidth: "800px",
});

const ButtonContainer = styled(Box)({
    display: "flex",
    gap: "15px",
    marginTop: "30px",
    flexWrap: "wrap",
    justifyContent: "center",
    paddingBottom: "30px",
});

const StyledButton = styled(Button)({
    padding: "12px 24px",
    fontSize: "1rem",
    fontWeight: "bold",
    textTransform: "none",
    minWidth: "180px",
});

const ListContainer = styled(Box)({
    width: "100%",
    maxWidth: "700px",
    marginTop: "10px",
});

const ListItem = styled(Box)(({ selected, disabled }) => ({
    padding: "16px",
    border: "2px solid",
    borderColor: selected ? "#28a745" : "#ccc",
    backgroundColor: selected ? "#e6ffed" : "#fff",
    borderRadius: "8px",
    marginBottom: "10px",
    boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
    cursor: disabled ? "not-allowed" : "pointer",
    opacity: disabled ? 0.6 : 1,
    transition: "all 0.2s ease-in-out",
    "&:hover": {
        borderColor: disabled ? "#ccc" : "#28a745",
        backgroundColor: disabled ? "#fff" : "#f1fdf5",
    },
    "&:last-child": {
        marginBottom: 0,
    },
}));

const SwitchContainer = styled(Box)({
    marginTop: "16px",
    width: "100%",
    maxWidth: "700px",
    display: "flex",
    flexDirection: "column",
    gap: "10px",
    alignItems: "start",
});

const OngoingVoting = () => {
    const navigate = useNavigate();

    const [votoObservado, setVotoObservado] = useState(false);
    const [votoEnBlanco, setVotoEnBlanco] = useState(false);
    const [votoAnulado, setVotoAnulado] = useState(false);
    const [listas, setListas] = useState([]);
    const [parties, setParties] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const [selectedList, setSelectedList] = useState(null);

    const partyIdToName = parties.reduce((acc, party) => {
        acc[party.id] = party.name;
        return acc;
    }, {});

useEffect(() => {
    const fetchData = async () => {
        setLoading(true);
        setError(null);

        const token = localStorage.getItem("token");
        if (!token) {
            setError("No se encontró token de autenticación. Por favor, inicia sesión.");
            setLoading(false);
            return;
        }

        try {
            const listasRes = await fetch("http://localhost:8080/party-lists", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (!listasRes.ok) throw new Error("Error al obtener listas");
            const listasData = await listasRes.json();

            const partiesRes = await fetch("http://localhost:8080/political-parties", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });
            if (!partiesRes.ok) throw new Error("Error al obtener partidos");
            const partiesData = await partiesRes.json();

            setListas(listasData);
            setParties(partiesData);
        } catch (err) {
            setError(err.message);
        } finally {
            setLoading(false);
        }
    };

    fetchData();
}, []);

    const handleResetVote = () => {
        setVotoObservado(false);
        setVotoEnBlanco(false);
        setVotoAnulado(false);
        setSelectedList(null);
    };

    const handleLogout = () => {
        localStorage.removeItem("token");
        navigate("/login");
    };

    const handleVotoObservadoChange = (event) => {
        setVotoObservado(event.target.checked);
    };

    const handleVotoEnBlancoChange = (event) => {
        const checked = event.target.checked;
        setVotoEnBlanco(checked);
        if (checked) {
            setSelectedList(null);
            setVotoAnulado(false);
        }
    };

    const handleVotoAnuladoChange = (event) => {
        const checked = event.target.checked;
        setVotoAnulado(checked);
        if (checked) {
            setVotoEnBlanco(false);
            setSelectedList(null);
        }
    };

    const handleSelectList = (list_number) => {
        if (votoEnBlanco || votoAnulado) return;
        setSelectedList(list_number);
    };

    const handleConfirmVote = () => {
        if (!selectedList && !votoEnBlanco && !votoAnulado) {
            alert("Por favor, seleccioná una lista o marcá voto en blanco o voto anulado antes de confirmar tu voto.");
            return;
        }

        const selectedListObj = listas.find((l) => l.list_number === selectedList);

        const voteData = {
            listasSeleccionadas: votoEnBlanco || votoAnulado ? [] : [selectedList],
            votoObservado,
            votoEnBlanco,
            votoAnulado,
            partyId: selectedListObj ? selectedListObj.party_id : null,
            partyName: selectedListObj ? (partyIdToName[selectedListObj.party_id] || "Desconocido") : null,
        };

        localStorage.setItem("voteData", JSON.stringify(voteData));
        navigate("/confirm-your-vote");
    };

    return (
        <>
            {globalStyles}
            <Box sx={{ minHeight: "100vh", display: "flex", flexDirection: "column" }}>
                <Header sx={{ display: "flex", justifyContent: "space-between", alignItems: "center" }}>
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
                    <Alert>
                        <WarningIcon sx={{ color: "#f0ad4e", fontSize: "24px" }} />
                        <Typography sx={{ color: "#856404", fontWeight: 500 }}>
                            Su voto será <strong>SECRETO</strong>. Seleccione a su candidato cuidadosamente.
                        </Typography>
                    </Alert>

                    <Box sx={{ width: "100%", maxWidth: "700px", textAlign: "left" }}>
                        <Typography
                            component="h1"
                            variant="h5"
                            gutterBottom
                            sx={{
                                fontWeight: "bold",
                                color: "#040111",
                                marginBottom: "10px",
                                fontSize: "1.5rem",
                            }}
                        >
                            Listas Disponibles
                        </Typography>

                        <Typography
                            variant="body2"
                            sx={{
                                color: "#666",
                                marginBottom: "20px",
                                fontSize: "0.95rem",
                            }}
                        >
                            Selecciona una lista haciendo clic sobre ella
                        </Typography>
                    </Box>

                    <ListContainer>
                        {loading && (
                            <Box sx={{ textAlign: "center", padding: "40px" }}>
                                <Typography variant="h6" color="text.secondary">
                                    Cargando listas...
                                </Typography>
                            </Box>
                        )}

                        {error && (
                            <Box sx={{ textAlign: "center", padding: "40px" }}>
                                <Typography color="error" variant="h6">
                                    {error}
                                </Typography>
                            </Box>
                        )}

                        {!loading && !error && listas.length === 0 && (
                            <Box sx={{ textAlign: "center", padding: "40px" }}>
                                <Typography variant="h6" color="text.secondary">
                                    No hay listas disponibles.
                                </Typography>
                            </Box>
                        )}

                        {!loading &&
                            !error &&
                            listas.map((lista) => (
                                <ListItem
                                    key={lista.list_number}
                                    selected={selectedList === lista.list_number}
                                    disabled={votoEnBlanco || votoAnulado}
                                    onClick={() => handleSelectList(lista.list_number)}
                                >
                                    <Typography
                                        variant="body1"
                                        sx={{
                                            fontWeight: "bold",
                                            marginBottom: "4px",
                                        }}
                                    >
                                        Lista #{lista.list_number}
                                    </Typography>
                                    <Typography variant="body2" color="text.secondary">
                                        Partido: {partyIdToName[lista.party_id] || "Desconocido"}
                                    </Typography>
                                </ListItem>
                            ))}
                    </ListContainer>

                    <SwitchContainer>
                        <FormControlLabel
                            control={
                                <Switch
                                    checked={votoObservado}
                                    onChange={handleVotoObservadoChange}
                                    color="primary"
                                />
                            }
                            label="Voto Observado"
                        />

                        <FormControlLabel
                            control={
                                <Switch
                                    checked={votoEnBlanco}
                                    onChange={handleVotoEnBlancoChange}
                                    color="primary"
                                />
                            }
                            label="Voto en Blanco"
                        />

                        <FormControlLabel
                            control={
                                <Switch
                                    checked={votoAnulado}
                                    onChange={handleVotoAnuladoChange}
                                    color="error"
                                />
                            }
                            label="Voto Anulado"
                        />
                    </SwitchContainer>

                    <ButtonContainer>
                        <StyledButton
                            variant="contained"
                            onClick={handleConfirmVote}
                            sx={{
                                backgroundColor: "#3d56a6",
                                "&:hover": {
                                    backgroundColor: "#2c3e50",
                                },
                            }}
                        >
                            CONFIRMAR VOTO
                        </StyledButton>
                        <StyledButton
                            variant="contained"
                            onClick={handleResetVote}
                            sx={{
                                backgroundColor: "#6c757d",
                                "&:hover": {
                                    backgroundColor: "#5a6268",
                                },
                            }}
                        >
                            RESETEAR
                        </StyledButton>
                    </ButtonContainer>
                </MainContent>
            </Box>
        </>
    );
};

export default OngoingVoting;
