# Aplicación Multicontainer con Docker Compose

Esta es una aplicación compuesta por 3 contenedores:

- **Api** (node)
- **Frontend** (react)
- **Base de Datos** (mongoDB)

El objetivo de esta aplicación es permitir el ingreso de fotografías tomadas desde el navegador web y almacenarlas en una base de datos no relacional. Además, se debe poder visualizar las fotos ingresadas. Para lograr esto, se utiliza una base de datos MongoDB y se guarda la información de forma persistente.

## VIdeo

El video se encuentra en este enlace: [Tarea2](https://drive.google.com/file/d/1heuq_7OUdn_m0wpqvpP7uQaV-FaV_bll/view?usp=sharinghttps://youtu.be/)
*Cuidado con el volumen al inicio*

## Requisitos de la base de datos

La base de datos debe almacenar los siguientes datos:

1. Fotografía en formato base64.
2. Fecha de carga de la fotografía.

## Docker Compose

Toda la aplicación está contenerizada utilizando Docker Compose. Esto permite gestionar y orquestar los contenedores de manera sencilla y eficiente.

## Uso

Para ejecutar la aplicación, se debe clonar el repositorio y ejecutar el siguiente comando en la raíz del proyecto:

```bash
docker-compose up
```
