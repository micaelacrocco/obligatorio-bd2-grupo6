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

const Login = () => {
  const [formData, setFormData] = useState({ ci: "", credential: "", password: "" });
  const [errors, setErrors] = useState({});
  const [isSubmitting, setIsSubmitting] = useState(false);
  const navigate = useNavigate();

  const validate = () => {
    const newErrors = {};
    if (!/^\d{7,8}$/.test(formData.ci)) {
      newErrors.ci = "La cédula debe tener 7 u 8 dígitos.";
    }
    if (!/^[A-Za-z]{3}\d{6}$/.test(formData.credential)) {
      newErrors.credential = "Formato ABC123456 (3 letras + 6 números).";
    }
    if (!formData.password) {
      newErrors.password = "La contraseña es obligatoria.";
    }
    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleChange = (e) => {
    const { name, value } = e.target;

    setFormData((prev) => ({
      ...prev,
      [name]:
        name === "credential"
          ? value.toUpperCase()
          : name === "ci"
          ? value.replace(/\D/g, "") // solo dígitos para ci
          : value,
    }));

    if (errors[name]) {
      setErrors((prev) => ({ ...prev, [name]: "" }));
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!validate()) return;

    setIsSubmitting(true);

    const payload = {
      ci: parseInt(formData.ci, 10),
      credential: formData.credential.trim().toUpperCase(),
      password: formData.password,
    };

    console.log("Payload enviado al backend:", payload);

    try {
      const res = await fetch("http://localhost:8080/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(payload),
      });

      if (!res.ok) {
        const errorRes = await res.json();
        throw new Error(errorRes.error || "Credenciales incorrectas");
      }

      const { token } = await res.json();
      localStorage.setItem("token", token);
      localStorage.setItem("ci", payload.ci);

      navigate("/electoral-system");
    } catch (err) {
      setErrors({ submit: err.message });
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <>
      {globalStyles}
      <Box
        sx={{
          width: '100%',
          height: '100%',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          justifyContent: 'center',
        }}
      >
        <LogoContainer>
          <LogoCorteElectoral />
        </LogoContainer>

        <StyledCard>
          <CardContent sx={{ p: 0, '&:last-child': { pb: 0 } }}>
            <Typography
              component="h1"
              variant="h5"
              align="start"
              gutterBottom
              sx={{ fontWeight: 'bold', color: '#040111' }}
            >
              Acceso al Sistema de Votación
            </Typography>

            <Typography
              variant="body2"
              align="start"
              sx={{ fontWeight: 500, color: '#636365', mb: 2 }}
            >
              Ingrese sus credenciales
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
                inputProps={{ maxLength: 8 }}
              />

              <TextField
                margin="normal"
                required
                fullWidth
                id="credential"
                label="Credencial Cívica"
                name="credential"
                value={formData.credential}
                onChange={handleChange}
                error={!!errors.credential}
                helperText={errors.credential}
                disabled={isSubmitting}
                variant="outlined"
                inputProps={{ maxLength: 9 }} // acá 9 chars (3 letras + 6 números)
              />

              <TextField
                margin="normal"
                required
                fullWidth
                id="password"
                label="Contraseña"
                name="password"
                type="password"
                value={formData.password}
                onChange={handleChange}
                error={!!errors.password}
                helperText={errors.password}
                disabled={isSubmitting}
                variant="outlined"
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
                disabled={isSubmitting}
                sx={{
                  mt: 2,
                  mb: 1,
                  py: 1.5,
                  fontSize: '1rem',
                  fontWeight: 'bold',
                }}
                startIcon={isSubmitting ? <CircularProgress size={20} /> : null}
              >
                {isSubmitting ? "Ingresando..." : "Ingresar al Sistema"}
              </StyledButton>
            </Box>
          </CardContent>
        </StyledCard>

        <Typography
          variant="body2"
          align="center"
          sx={{ mt: 1 }}
        >
          ¿No tenés cuenta?{" "}
          <Link to="/register" style={{ color: '#3d56a6', fontWeight: 'bold', textDecoration: 'none' }}>
            Registrate acá
          </Link>
        </Typography>
      </Box>
    </>
  );
};

export default Login;
