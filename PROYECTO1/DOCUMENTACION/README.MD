# PROYECTO 1

| Nombre                       | Carnet    |
| ---------------------------- | --------- |
| Angel Francisco Sique Santos | 202012039 |

## MÓDULOS

## Módulo del Kernel Linux para Monitoreo de CPU (cpu_so1_1s2024)

Este código fuente implementa un módulo para el kernel de Linux que expone información detallada sobre la CPU y los procesos del sistema en formato JSON. La información se hace accesible a través de un archivo ubicado en el sistema de archivos `/proc`.

**Funciones:**

* `escribir_a_proc`: Esta función calcula y escribe información relevante sobre la CPU y los procesos en formato JSON al archivo ubicado en el sistema de archivos proc.
* Calcula el tiempo total de CPU, el tiempo libre y el porcentaje de uso.
* Recorre todos los procesos del sistema y recopila información como ID de proceso, nombre, estado, uso de memoria y procesos secundarios.
* `abrir_aproc`:  Esta función se activa cuando alguien intenta abrir el archivo asociado al módulo en el sistema. .
* `modulo_init`: Se ejecuta durante la carga del módulo. Crea el archivo proc denominado `cpu_so1_1s2024` y registra las funciones para manejar la apertura y lectura del archivo.
* `modulo_cleanup`: Se ejecuta durante la descarga del módulo. Elimina el archivo proc asociado `cpu_so1_1s2024`.

## Módulo del Kernel Linux para Monitoreo de RAM (ram_so1_1s2024)

De manera similar al módulo anterior, este código fuente implementa un módulo para el kernel de Linux que expone información sobre la memoria RAM del sistema en formato JSON. La información se hace accesible a través de un archivo ubicado en el sistema de archivos `/proc`.

**Funciones:**

* `write_to_proc`: Esta función recupera información sobre la memoria del sistema utilizando la función `si_meminfo` y la escribe en formato JSON al archivo proc asociado.
* `open_proc`:  Esta función se activa cuando alguien intenta abrir el archivo asociado al módulo en el sistema. .
* `ram_module_init`: Se ejecuta durante la carga del módulo. Crea el archivo proc llamado `ram_so1_1s2024` y registra las funciones para manejar la apertura y lectura del archivo.
* `ram_module_exit`: Se ejecuta durante la descarga del módulo. Elimina el archivo proc asociado `ram_so1_1s2024`.

## BASE DE DATOS

Se ha creado una base de datos llamada PROYECTO1 con el objetivo de almacenar información sobre el monitoreo de recursos del sistema. Se utilizan tipos de datos de gran tamaño, como BIGINT, debido a que Linux puede proporcionar números muy grandes para algunos valores.

La base de datos contiene las siguientes tablas:

### MemoriaRAM

Esta tabla almacena información relacionada con la memoria RAM del sistema.

| Campo       | Tipo     | Descripción               |
|-------------|----------|---------------------------|
| id          | INT      | Clave primaria autoincremental |
| total       | BIGINT   | Total de memoria RAM en bytes |
| enUso       | BIGINT   | Memoria RAM en uso en bytes |
| porcentaje  | BIGINT   | Porcentaje de memoria en uso |
| libre       | BIGINT   | Memoria RAM libre en bytes |
| fechaHora   | DATETIME | Fecha y hora de registro    |

### MemoriaCPU

Esta tabla guarda información sobre la CPU del sistema.

| Campo       | Tipo     | Descripción               |
|-------------|----------|---------------------------|
| id          | INT      | Clave primaria autoincremental |
| total       | BIGINT   | Total de capacidad de la CPU en bytes |
| enUso       | BIGINT   | Uso actual de la CPU en bytes |
| porcentaje  | BIGINT   | Porcentaje de uso de la CPU |
| libre       | BIGINT   | Capacidad de la CPU libre en bytes |
| fechaHora   | DATETIME | Fecha y hora de registro    |

### ProcesoPadre

Esta tabla registra información sobre los procesos principales del sistema.

| Campo       | Tipo     | Descripción               |
|-------------|----------|---------------------------|
| pid         | INT      | Identificador del proceso |
| name        | VARCHAR(255) | Nombre del proceso    |
| ram         | INT      | Uso de memoria del proceso en bytes |
| state       | INT      | Estado del proceso         |
| user        | INT      | Usuario asociado al proceso |

### ProcesoHijo

Esta tabla almacena información sobre los procesos secundarios del sistema, haciendo referencia al proceso padre al que pertenecen.

| Campo       | Tipo     | Descripción               |
|-------------|----------|---------------------------|
| pid         | INT      | Identificador del proceso |
| name        | VARCHAR(255) | Nombre del proceso    |
| pidPadre    | INT      | Identificador del proceso padre |
| state       | INT      | Estado del proceso         |

## SIMULADOR DE ESTADOS

### StartProcess

Esta función se utiliza para iniciar un nuevo proceso en el sistema. Toma la solicitud HTTP del cliente y ejecuta el comando especificado, que en este caso es "./proyecto1". Si el proceso se inicia correctamente, responde al cliente con el PID (Identificador de Proceso) del proceso hijo recién creado.

### StopProcess

La función StopProcess se utiliza para detener un proceso en ejecución. Toma el PID del proceso como parámetro de la solicitud HTTP y envía una señal SIGSTOP al proceso correspondiente, lo que suspende su ejecución. Luego, responde al cliente con un mensaje de éxito si el proceso se detiene correctamente.

### ResumeProcess

ResumeProcess se utiliza para reanudar la ejecución de un proceso que ha sido previamente detenido. Al igual que StopProcess, toma el PID del proceso como parámetro y envía una señal SIGCONT al proceso correspondiente para reanudar su ejecución. Luego, responde al cliente con un mensaje de éxito si el proceso se reanuda correctamente.

### KillProcess

Esta función se utiliza para terminar un proceso de manera abrupta. Toma el PID del proceso como parámetro y envía una señal de terminación al proceso, lo que causa su finalización inmediata. Dependiendo de la implementación, se puede enviar una señal SIGKILL (-9) o SIGTERM al proceso. Luego, responde al cliente con un mensaje de éxito si el proceso se termina correctamente.

Estas funciones permiten a los usuarios controlar la ejecución y el estado de los procesos en el sistema de forma remota a través de solicitudes HTTP.

Estas funciones se utilizan para controlar procesos, donde StartProcess inicia un nuevo proceso, StopProcess detiene un proceso en ejecución, ResumeProcess reanuda un proceso detenido y KillProcess termina un proceso en ejecución.

## FRONTEND (REACT)

### Frontend con Tailwind CSS y Paleta de Colores

El frontend de la aplicación se desarrolló utilizando Tailwind CSS y una paleta de colores definida en el documento CSS usando variables.

```css
:root {
  --primary-100: #C7B299;
  --primary-200: #a9957d;
  --primary-300: #675640;
  --accent-100: #A68B5B;
  --accent-200: #483507;
  --text-100: #4D4D4D;
  --text-200: #797979;
  --bg-100: #F5EFE7;
  --bg-200: #ebe5dd;
  --bg-300: #c2bdb5;
}
```

### Páginas de la Aplicación

Se crearon varias páginas utilizando React y React Router:

* **Home**

* **MonitoreoEnTiempoReal**

* **MonitoreoHistorico**

* **Arbol**

* **CambioEstados**

### Componente `App`

El componente `App` utiliza `react-router-dom` para definir las rutas de la aplicación y asociarlas con los componentes correspondientes.

```jsx
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './components/Home';
import Navbar from './components/Navbar';
import MonitoreoEnTiempoReal from './components/MonitoreoEnTiempoReal';
import MonitoreoHistorico from './components/MonitoreoHistorico';
import Arbol from './components/Arbol';
import CambioEstados from './components/CambioEstados';

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/home' element={<Home />} />
        <Route path='/MonitoreoEnTiempoReal' element={<MonitoreoEnTiempoReal />} />
        <Route path='/MonitoreoHistorico' element={<MonitoreoHistorico />} />
        <Route path='/ArbolProcesos' element={<Arbol />} />
        <Route path='/CambioEstados' element={<CambioEstados />} />
      </Routes>
    </Router>
  );
}

export default App;
```

Esto facilita la navegación entre las páginas de la aplicación, manteniendo una interfaz de usuario coherente y fácil de usar.

## BACKEND (GOLANG)

Se crea un servidor HTTP en Go que recopila datos de la RAM y la CPU de un sistema Linux y los almacena en una base de datos MySQL. Aquí está el resumen de las principales funcionalidades:

1. **Paquetes Importados**: Importa los paquetes necesarios, incluyendo `database/sql`, `encoding/json`, `net/http`, entre otros.

2. **Estructuras de Datos**: Define las estructuras `RamData` y `CpuData` para almacenar los datos de RAM y CPU respectivamente.

3. **Conexión a la Base de Datos**: Establece la conexión a la base de datos MySQL utilizando las credenciales proporcionadas en un archivo `.env`.

4. **Manejo de Solicitudes HTTP**:
   * `saveRamData`: Maneja la solicitud POST para guardar datos de RAM en la base de datos.
   * `getRamData`: Maneja la solicitud GET para obtener datos de RAM de la base de datos.
   * `saveCpuData`: Maneja la solicitud POST para guardar datos de CPU en la base de datos.
   * `getCpuData`: Maneja la solicitud GET para obtener datos de CPU de la base de datos.
   * `getAllCpuData`: Maneja la solicitud GET para obtener todos los datos de CPU de la base de datos.
   * `getAllRamData`: Maneja la solicitud GET para obtener todos los datos de RAM de la base de datos.
   * `handleGetProcesses`: Maneja la solicitud GET para obtener datos de procesos del sistema y los almacena en la base de datos.

5. **Operaciones con Procesos**:
   * `StartProcess`: Inicia un nuevo proceso y devuelve su ID.
   * `StopProcess`: Detiene un proceso dado su ID.
   * `ResumeProcess`: Reanuda un proceso dado su ID.
   * `KillProcess`: Mata un proceso dado su ID.

6. **Enrutamiento y Configuración CORS**: Configura las rutas de API y habilita el Cross-Origin Resource Sharing (CORS) para permitir solicitudes desde cualquier origen.

7. **Inicio del Servidor**: Inicia un servidor HTTP en el puerto 8080 y maneja las solicitudes utilizando el enrutador definido.
