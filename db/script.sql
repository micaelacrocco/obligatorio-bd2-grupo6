-- Create and select the database
CREATE DATABASE IF NOT EXISTS elections;
USE elections;

-- Citizens table
CREATE TABLE CITIZENS
(
    id         INT PRIMARY KEY,
    first_name VARCHAR(100),
    last_name  VARCHAR(100),
    birth_date DATE,
    credential VARCHAR(50)
);

-- Votes by person
CREATE TABLE PERSON_VOTES
(
    id          INT PRIMARY KEY AUTO_INCREMENT,
    vote_date   DATE,
    is_observed BOOLEAN,
    vote_type   VARCHAR(50),
    citizen_id  INT,
    FOREIGN KEY (citizen_id) REFERENCES CITIZENS (id)
);

-- Parties and their lists
CREATE TABLE PARTIES
(
    id   INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100)
);

CREATE TABLE PARTY_LISTS
(
    list_number INT PRIMARY KEY,
    party_id    INT,
    FOREIGN KEY (party_id) REFERENCES PARTIES (id)
);

-- Votes cast to party lists
CREATE TABLE LIST_VOTES
(
    id          INT PRIMARY KEY AUTO_INCREMENT,
    vote_date   DATE,
    list_number INT,
    FOREIGN KEY (list_number) REFERENCES PARTY_LISTS (list_number)
);

-- Citizens who are candidates in a list
CREATE TABLE CANDIDATES
(
    citizen_id     INT,
    list_number    INT,
    start_date     DATE,
    end_date       DATE,
    candidacy_type VARCHAR(50),
    PRIMARY KEY (citizen_id, list_number),
    FOREIGN KEY (citizen_id) REFERENCES CITIZENS (id),
    FOREIGN KEY (list_number) REFERENCES PARTY_LISTS (list_number)
);

-- Party leaders
CREATE TABLE LEADERS
(
    citizen_id    INT,
    party_id      INT,
    election_year INT,
    role          VARCHAR(100),
    PRIMARY KEY (citizen_id, party_id, election_year),
    FOREIGN KEY (citizen_id) REFERENCES CITIZENS (id),
    FOREIGN KEY (party_id) REFERENCES PARTIES (id)
);

-- Departments and police stations
CREATE TABLE DEPARTMENTS
(
    id   INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100)
);

CREATE TABLE POLICE_STATIONS
(
    id             INT PRIMARY KEY AUTO_INCREMENT,
    station_number INT,
    address        VARCHAR(200),
    department_id  INT,
    FOREIGN KEY (department_id) REFERENCES DEPARTMENTS (id)
);

-- Zones and polling places
CREATE TABLE ZONES
(
    id            INT PRIMARY KEY AUTO_INCREMENT,
    name          VARCHAR(100),
    address       VARCHAR(200),
    department_id INT,
    FOREIGN KEY (department_id) REFERENCES DEPARTMENTS (id)
);

CREATE TABLE POLLING_PLACES
(
    id      INT PRIMARY KEY AUTO_INCREMENT,
    name    VARCHAR(100),
    type    VARCHAR(50),
    address VARCHAR(200),
    zone_id INT,
    FOREIGN KEY (zone_id) REFERENCES ZONES (id)
);

-- Police agents assigned to polling places
CREATE TABLE POLICE_AGENTS
(
    citizen_id        INT,
    police_station_id INT,
    polling_place_id  INT,
    PRIMARY KEY (citizen_id, police_station_id, polling_place_id),
    FOREIGN KEY (citizen_id) REFERENCES CITIZENS (id),
    FOREIGN KEY (police_station_id) REFERENCES POLICE_STATIONS (id),
    FOREIGN KEY (polling_place_id) REFERENCES POLLING_PLACES (id)
);

-- Circuits and polling tables
CREATE TABLE CIRCUITS
(
    id               INT PRIMARY KEY AUTO_INCREMENT,
    location         VARCHAR(100),
    is_accessible    BOOLEAN,
    credential_start INT,
    credential_end   INT,
    polling_place_id INT,
    FOREIGN KEY (polling_place_id) REFERENCES POLLING_PLACES (id)
);

CREATE TABLE TABLES
(
    id         INT PRIMARY KEY AUTO_INCREMENT,
    circuit_id INT,
    FOREIGN KEY (circuit_id) REFERENCES CIRCUITS (id)
);

-- Members assigned to polling tables
CREATE TABLE TABLE_MEMBERS
(
    table_id         INT,
    citizen_id       INT,
    integration_date DATE,
    duty             VARCHAR(100),
    PRIMARY KEY (table_id, citizen_id),
    FOREIGN KEY (table_id) REFERENCES TABLES (id),
    FOREIGN KEY (citizen_id) REFERENCES CITIZENS (id)
);
