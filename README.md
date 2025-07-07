# Sistema de Votación Electrónica - Obligatorio BD2

## 📋 Descripción del Proyecto

Sistema de votación electrónica desarrollado para la Corte Electoral de Uruguay como MVP (Minimum Viable Product). El sistema permite gestionar el proceso electoral.

## 🎯 Objetivos

- Desarrollar un sistema electoral seguro que garantice el **voto secreto**
- Permitir votación electrónica en múltiples circuitos

## 👥 Equipo de Desarrollo

**Grupo 6 - Bases de Datos II**
- Arispe Oviedo Leonardo Roy
- Crocco Micaela
- Lavecchia Jose Ignacio

**Universidad Católica del Uruguay - Facultad de Ingeniería y Tecnologías**

## 🏗️ Arquitectura del Sistema

### Tecnologías Utilizadas

- **Base de Datos:** MySQL
- **Backend:** GO
- **Frontend:** React
- **Contenedores:** Docker & Docker Compose
- **Arquitectura:** Cliente-Servidor

### Características Principales

- ✅ **Voto Secreto**: Sistema que registra que se votó sin trackear por quién
- ✅ **Multiusuario**: Soporte para votantes y administradores
- ✅ **Control de Circuitos**: Gestión de establecimientos, zonas y mesas
- ✅ **Reportes Estadísticos**: Resultados por circuito
- ✅ **Seguridad**: Prevención de votos múltiples y control de acceso

## 🚀 Instalación y Ejecución

### Prerrequisitos

- Docker y Docker Compose instalados
- MySQL 8.0 o superior

### Pasos de Instalación

1. **Clonar el repositorio**
   ```bash
   git clone https://github.com/micaelacrocco/obligatorio-bd2-grupo6.git
   cd obligatorio-bd2-grupo6
   ```

2. **Configurar variables de entorno**
   ```bash
   cp .env.example .env
   # Editar .env con las credenciales de base de datos
   ```

3. **Ejecutar con Docker Compose**
   ```bash
   docker-compose up -d
   ```

4. **Inicializar la base de datos**
   ```bash
   # Ejecutar scripts de creación
   docker-compose exec db mysql -u root -p < scripts/create_schema.sql
   docker-compose exec db mysql -u root -p < scripts/insert_test_data.sql
   ```

### Ejecución Local (Desarrollo)

```bash
# Backend
cd backend
go run main.go

# Frontend
cd frontend
npm start
```

## 🗂️ Estructura del Proyecto

```
obligatorio-bd2-grupo6/
├── backend/
│   ├── configuration/
│   ├── db/
│   └── domains/
│   │   ├── interfaces/
│   │   ├── repositories/
│   │   └── usecases/
│   ├── dtos/
│   ├── handlers/
│   ├── middlewares/
│   ├── models/
│   └── utils/
│   └── main.go
├── frontend/
│   ├── public/
│   └── src/
│       ├── assets/
│       ├── components/
│       ├── layouts/
│       ├── pages/
│       ├── routes/
│       └── App.js
├── docker-compose.yml
└── README.md
```

Este proyecto es desarrollado con fines académicos para la Universidad Católica del Uruguay.
