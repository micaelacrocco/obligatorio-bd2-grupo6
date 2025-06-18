-- Crear y seleccionar la base de datos
CREATE DATABASE IF NOT EXISTS elecciones;
USE elecciones;

-- Tabla base de ciudadanos
CREATE TABLE CIUDADANOS (
    ci INT PRIMARY KEY,
    nombre VARCHAR(100),
    apellido VARCHAR(100),
    fecha_nacimiento DATE,
    credencial VARCHAR(50)
);

-- Tabla de votos por persona
CREATE TABLE VOTOS_PERSONAS (
    id INT PRIMARY KEY AUTO_INCREMENT,
    fecha DATE,
    observado BOOLEAN,
    tipo_voto VARCHAR(50),
    ci_ciudadano INT,
    FOREIGN KEY (ci_ciudadano) REFERENCES CIUDADANOS(ci)
);

-- Partidos y sus listas
CREATE TABLE PARTIDOS (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(100)
);

CREATE TABLE LISTAS (
    numero INT PRIMARY KEY,
    id_partido INT,
    FOREIGN KEY (id_partido) REFERENCES PARTIDOS(id)
);

-- Votos emitidos a listas
CREATE TABLE VOTOS_LISTAS (
    id INT PRIMARY KEY AUTO_INCREMENT,
    fecha DATE,
    numero_lista INT,
    FOREIGN KEY (numero_lista) REFERENCES LISTAS(numero)
);

-- Ciudadanos que integran listas
CREATE TABLE INTEGRAN (
    ci_ciudadano INT,
    numero_lista INT,
    fecha_inicio DATE,
    fecha_fin DATE,
    tipo_candidatura VARCHAR(50),
    PRIMARY KEY (ci_ciudadano, numero_lista),
    FOREIGN KEY (ci_ciudadano) REFERENCES CIUDADANOS(ci),
    FOREIGN KEY (numero_lista) REFERENCES LISTAS(numero)
);

-- Ciudadanos que son líderes de partidos
CREATE TABLE LIDER (
    ci_ciudadano INT,
    id_partido INT,
    anio_electoral INT,
    cargo VARCHAR(100),
    PRIMARY KEY (ci_ciudadano, id_partido, anio_electoral),
    FOREIGN KEY (ci_ciudadano) REFERENCES CIUDADANOS(ci),
    FOREIGN KEY (id_partido) REFERENCES PARTIDOS(id)
);

-- Departamentos y comisarías
CREATE TABLE DEPARTAMENTOS (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(100)
);

CREATE TABLE COMISARIAS (
    id INT PRIMARY KEY AUTO_INCREMENT,
    numero INT,
    direccion VARCHAR(200),
    id_departamento INT,
    FOREIGN KEY (id_departamento) REFERENCES DEPARTAMENTOS(id)
);

-- Zonas y establecimientos
CREATE TABLE ZONAS (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(100),
    direccion VARCHAR(200),
    id_departamento INT,
    FOREIGN KEY (id_departamento) REFERENCES DEPARTAMENTOS(id)
);

CREATE TABLE ESTABLECIMIENTOS (
    id INT PRIMARY KEY AUTO_INCREMENT,
    nombre VARCHAR(100),
    tipo VARCHAR(50),
    direccion VARCHAR(200),
    id_zona INT,
    FOREIGN KEY (id_zona) REFERENCES ZONAS(id)
);

-- Asignación de agentes a comisarías y establecimientos
CREATE TABLE AGENTES_POLICIALES (
    ci_ciudadano INT,
    id_comisarias INT,
    id_establecimiento INT,
    PRIMARY KEY (ci_ciudadano, id_comisarias, id_establecimiento),
    FOREIGN KEY (ci_ciudadano) REFERENCES CIUDADANOS(ci),
    FOREIGN KEY (id_comisarias) REFERENCES COMISARIAS(id),
    FOREIGN KEY (id_establecimiento) REFERENCES ESTABLECIMIENTOS(id)
);

-- Circuitos y mesas de votación
CREATE TABLE CIRCUITOS (
    id INT PRIMARY KEY AUTO_INCREMENT,
    pueblo_ciudad_paraje VARCHAR(100),
    accesible BOOLEAN,
    inicio_credencial INT,
    fin_credencial INT,
    id_establecimiento INT,
    FOREIGN KEY (id_establecimiento) REFERENCES ESTABLECIMIENTOS(id)
);

CREATE TABLE MESAS (
    id INT PRIMARY KEY AUTO_INCREMENT,
    id_circuito INT,
    FOREIGN KEY (id_circuito) REFERENCES CIRCUITOS(id)
);

-- Ciudadanos que conforman mesas
CREATE TABLE CONFORMAN (
    id_mesa INT,
    ci_ciudadano INT,
    fecha_integracion DATE,
    trabajo VARCHAR(100),
    PRIMARY KEY (id_mesa, ci_ciudadano),
    FOREIGN KEY (id_mesa) REFERENCES MESAS(id),
    FOREIGN KEY (ci_ciudadano) REFERENCES CIUDADANOS(ci)
);
