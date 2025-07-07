import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import {
    Box,
    CardContent,
    TextField,
    Button,
    Typography,
    Alert,
    CircularProgress,
    styled,
    GlobalStyles
} from "@mui/material";
import { ReactComponent as LogoCorteElectoral } from "../assets/logo-corte-electoral.svg";

const globalStyles = (
    <GlobalStyles
        styles={{
            '*': { margin: 0, padding: 0, boxSizing: 'border-box' },
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

const StyledCard = styled(Box)({
    backgroundColor: 'white',
    border: '1px solid #e9ecef',
    borderRadius: '8px',
    padding: '20px',
    margin: '20px',
    maxWidth: '400px',
    width: '100%',
});

const LogoContainer = styled(Box)({
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

// ... (importaciones existentes)
const Register = () => {
    const [formData, setFormData] = useState({
        ci: "",
        credencial: "",
        password: "",
    });
    const [errors, setErrors] = useState({});
    const [isSubmitting, setIsSubmitting] = useState(false);
    const navigate = useNavigate();

    const validate = () => {
        const newErrors = {};
        if (!/^\d{7,8}$/.test(formData.ci)) {
            newErrors.ci = "La cédula debe tener 7 u 8 dígitos.";
        }
        if (!/^[A-Za-z]{3}\d{5}$/.test(formData.credencial)) {
            newErrors.credencial = "Formato ABC12345 (3 letras + 5 números).";
        }
        if (!formData.password || formData.password.length < 6) {
            newErrors.password = "La contraseña debe tener al menos 6 caracteres.";
        }
        setErrors(newErrors);
        return Object.keys(newErrors).length === 0;
    };

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: name === "credencial" ? value.toUpperCase() : value,
        }));
        if (errors[name]) {
            setErrors((prev) => ({ ...prev, [name]: "" }));
        }
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        if (!validate()) return;

        setIsSubmitting(true);
        try {
            const res = await fetch(`${import.meta.env.REACT_APP_API_URL}/auth/register`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(formData),
            });
            if (!res.ok) throw new Error("No se pudo registrar. Verifique los datos.");
            navigate("/login");
        } catch (err) {
            setErrors({ submit: err.message });
        } finally {
            setIsSubmitting(false);
        }
    };

    return (
        <>
            {globalStyles}
            <Box sx={{ width: '100%', height: '100%', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
                <LogoContainer>
                    <LogoCorteElectoral />
                </LogoContainer>

                <StyledCard>
                    <CardContent sx={{ p: 0, '&:last-child': { pb: 0 } }}>
                        <Typography component="h1" variant="h5" align="start" gutterBottom sx={{ fontWeight: 'bold', color: '#040111' }}>
                            Registro de Usuario
                        </Typography>

                        <Typography variant="body2" align="start" sx={{ fontWeight: 500, color: '#636365', mb: 2 }}>
                            Complete sus datos para crear una cuenta
                        </Typography>

                        <Box component="form" onSubmit={handleSubmit}>
                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                id="ci"
                                label="Cédula de Identidad"
                                name="ci"
                                value={formData.ci}
                                onChange={handleChange}
                                error={!!errors.ci}
                                helperText={errors.ci}
                                disabled={isSubmitting}
                                variant="outlined"
                                sx={{ mb: 1.5 }}
                            />

                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                id="credencial"
                                label="Credencial Cívica"
                                name="credencial"
                                value={formData.credencial}
                                onChange={handleChange}
                                error={!!errors.credencial}
                                helperText={errors.credencial}
                                disabled={isSubmitting}
                                variant="outlined"
                                sx={{ mb: 1.5 }}
                            />

                            <TextField
                                margin="normal"
                                required
                                fullWidth
                                type="password"
                                id="password"
                                label="Contraseña"
                                name="password"
                                value={formData.password}
                                onChange={handleChange}
                                error={!!errors.password}
                                helperText={errors.password}
                                disabled={isSubmitting}
                                variant="outlined"
                                sx={{ mb: 1.5 }}
                            />

                            {errors.submit && (
                                <Alert severity="error" sx={{ mb: 2 }}>
                                    {errors.submit}
                                </Alert>
                            )}

                            <StyledButton
                                type="submit"
                                fullWidth
                                variant="contained"
                                onClick={handleSubmit}
                                disabled={isSubmitting}
                                sx={{ mt: 2, mb: 0, py: 1.5, fontSize: '1rem', fontWeight: 'bold' }}
                                startIcon={isSubmitting ? <CircularProgress size={20} /> : null}
                            >
                                {isSubmitting ? "Registrando..." : "Registrarse"}
                            </StyledButton>
                        </Box>
                    </CardContent>
                </StyledCard>

                <Typography variant="body2" align="center" sx={{ mt: 1 }}>
                    ¿Ya tenés cuenta?{" "}
                    <Link to="/login" style={{ color: '#3d56a6', fontWeight: 'bold', textDecoration: 'none' }}>
                        Ingresá acá
                    </Link>
                </Typography>
            </Box>
        </>
    );
};

export default Register;
