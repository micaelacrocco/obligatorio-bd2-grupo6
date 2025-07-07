import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  Box,
  Button,
  Typography,
  styled,
  GlobalStyles,
  Paper
} from "@mui/material";
import {
  CheckCircle as CheckIcon
} from "@mui/icons-material";
import {
  PieChart,
  Pie,
  Cell,
  Tooltip,
  Legend,
  ResponsiveContainer
} from "recharts";

const COLORS = ["#8884d8", "#82ca9d", "#ffc658", "#ff7f7f"];

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
  width: '700px',
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

const InfoSection = styled(Box)({
    backgroundColor: 'white',
    border: '1px solid #e9ecef',
    borderRadius: '8px',
    padding: '20px',
    margin: '10px',
    maxWidth: '700px',
    width: '700px',
});

const ElectionResult = () => {
  const navigate = useNavigate();
  const [results, setResults] = useState([]);
  const circuitId = 1; // cambiar según corresponda

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  useEffect(() => {
    const token = localStorage.getItem("token");
    fetch(`http://localhost:8080/circuits/${circuitId}/results`, {
      headers: {
        "Authorization": `Bearer ${token}`
      }
    })
      .then(res => res.json())
      .then(data => setResults(data))
      .catch(err => console.error("Error al obtener resultados:", err));
  }, []);

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

          <Box elevation={3} sx={{ padding: 4, width: 700 }}>
            <Typography variant="h5" gutterBottom sx={{ fontWeight: 'bold' }}>
              Resultado de la Elección
            </Typography>

            <ResponsiveContainer width="100%" height={300}>
              <PieChart>
                <Pie
                  data={results}
                  dataKey="percentage"
                  nameKey="list"
                  cx="50%"
                  cy="50%"
                  outerRadius={100}
                  label={({ name, percent }) => `${name} ${(percent * 100).toFixed(0)}%`}
                >
                  {results.map((_, index) => (
                    <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                  ))}
                </Pie>
                <Tooltip formatter={(value) => `${value.toFixed(2)}%`} />
                <Legend />
              </PieChart>
            </ResponsiveContainer>

            <Box sx={{ mt: 3 }}>
              {results.map((result, idx) => (
                <Typography key={idx}>
                  <strong>{result.list}</strong> ({result.party_name}): {result.vote_count} votos - {result.percentage.toFixed(2)}%
                </Typography>
              ))}
            </Box>
          </Box>
        </MainContent>
      </Box>
    </>
  );
};

export default ElectionResult;
