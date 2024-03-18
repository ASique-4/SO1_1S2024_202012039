CREATE DATABASE IF NOT EXISTS PROYECTO1;

USE PROYECTO1;

CREATE TABLE MemoriaRAM (
    id INT PRIMARY KEY AUTO_INCREMENT,
    total BIGINT,
    enUso BIGINT,
    porcentaje BIGINT,
    libre BIGINT,
    fechaHora DATETIME
);

CREATE TABLE MemoriaCPU (
    id INT AUTO_INCREMENT,
    total BIGINT,
    enUso BIGINT,
    porcentaje BIGINT,
    libre BIGINT,
    fechaHora DATETIME,
    PRIMARY KEY (id)
);

CREATE TABLE ProcesoPadre (
    pid INT PRIMARY KEY,
    name VARCHAR(255),
    ram INT,
    state INT,
    user INT
);

CREATE TABLE ProcesoHijo (
    pid INT PRIMARY KEY,
    name VARCHAR(255),
    pidPadre INT,
    state INT,
    FOREIGN KEY (pidPadre) REFERENCES ProcesoPadre(pid)
);

