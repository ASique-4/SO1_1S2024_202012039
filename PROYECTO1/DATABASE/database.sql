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