package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// Datos de la RAM
type RamData struct {
	Total      uint64 `json:"totalRam"`
	EnUso      uint64 `json:"memoriaEnUso"`
	Porcentaje uint64 `json:"porcentaje"`
	Libre      uint64 `json:"libre"`
	FechaHora  string `json:"fechaHora"`
}

// Conexión a la base de datos
var db *sql.DB

// Datos de la CPU
type CpuData struct {
	Total      uint64 `json:"cpu_total"`
	EnUso      uint64 `json:"cpu_uso"`
	Porcentaje uint64 `json:"cpu_porcentaje"`
	Libre      uint64 `json:"cpu_libre"`
	FechaHora  string `json:"fechaHora"`
}

func saveCpuData(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("mpstat", "1", "1").Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	lines := strings.Split(string(out), "\n")
	fields := strings.Fields(lines[3])

	idleStr := strings.Replace(fields[11], ",", ".", -1)
	idle, err := strconv.ParseFloat(idleStr, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cpu := CpuData{
		Total:      100,
		EnUso:      uint64(100 - idle),
		Porcentaje: uint64(100 - idle),
		Libre:      uint64(idle),
		FechaHora:  time.Now().Format("2006-01-02 15:04:05"),
	}

	_, err = db.Exec("INSERT INTO MemoriaCPU (total, enUso, porcentaje, libre, fechaHora) VALUES (?, ?, ?, ?, ?)", cpu.Total, cpu.EnUso, cpu.Porcentaje, cpu.Libre, cpu.FechaHora)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func getCpuData(w http.ResponseWriter, r *http.Request) {
	// Primero, guarda los datos de la CPU
	saveCpuData(w, r)

	// Luego, obtén los datos de la CPU
	row := db.QueryRow("SELECT total, enUso, porcentaje, libre, fechaHora FROM MemoriaCPU ORDER BY fechaHora DESC LIMIT 1")

	var dato CpuData
	err := row.Scan(&dato.Total, &dato.EnUso, &dato.Porcentaje, &dato.Libre, &dato.FechaHora)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(dato)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Función para manejar las solicitudes HTTP POST
func saveRamData(w http.ResponseWriter, r *http.Request) {
	// Leer los datos de /proc/ram_so1_1s2024
	file, err := os.Open("/proc/ram_so1_1s2024")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var ram RamData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ram)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Guardar los datos en la base de datos
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err = db.Exec("INSERT INTO MemoriaRAM (total, enUso, porcentaje, libre, fechaHora) VALUES (?, ?, ?, ?, ?)", ram.Total, ram.EnUso, ram.Porcentaje, ram.Libre, now)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder al cliente con un código de estado 201 (Created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Datos de RAM guardados"}`))
}

// Función para manejar las solicitudes HTTP GET a /ram
func getRamData(w http.ResponseWriter, r *http.Request) {
	// Leer los datos de /proc/ram_so1_1s2024
	file, err := os.Open("/proc/ram_so1_1s2024")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var ram RamData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&ram)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Guardar los datos en la base de datos
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err = db.Exec("INSERT INTO MemoriaRAM (total, enUso, porcentaje, libre, fechaHora) VALUES (?, ?, ?, ?, ?)", ram.Total, ram.EnUso, ram.Porcentaje, ram.Libre, now)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Consultar el último dato ingresado en la tabla MemoriaRAM
	row := db.QueryRow("SELECT total, enUso, porcentaje, libre, fechaHora FROM MemoriaRAM ORDER BY fechaHora DESC LIMIT 1")

	var dato RamData
	err = row.Scan(&dato.Total, &dato.EnUso, &dato.Porcentaje, &dato.Libre, &dato.FechaHora)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir los datos a JSON
	jsonData, err := json.Marshal(dato)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Enviar los datos al cliente
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getAllCpuData(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT total, enUso, porcentaje, libre, fechaHora FROM MemoriaCPU ORDER BY fechaHora")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var datos []CpuData
	for rows.Next() {
		var dato CpuData
		err := rows.Scan(&dato.Total, &dato.EnUso, &dato.Porcentaje, &dato.Libre, &dato.FechaHora)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		datos = append(datos, dato)
	}

	jsonData, err := json.Marshal(datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getAllRamData(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT total, enUso, porcentaje, libre, fechaHora FROM MemoriaRAM ORDER BY fechaHora")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var datos []RamData
	for rows.Next() {
		var dato RamData
		err := rows.Scan(&dato.Total, &dato.EnUso, &dato.Porcentaje, &dato.Libre, &dato.FechaHora)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		datos = append(datos, dato)
	}

	jsonData, err := json.Marshal(datos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Función para manejar las solicitudes HTTP GET a /procesos
func handleGetProcesses(w http.ResponseWriter, r *http.Request) {
	// Abrir el archivo /proc/cpu_so1_1s2024
	file, err := os.Open("/proc/cpu_so1_1s2024")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Decodificar el archivo JSON
	var data map[string]interface{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Extraer la lista de procesos
	processesInterface, ok := data["processes"]
	if !ok {
		http.Error(w, "No se encontraron datos de procesos", http.StatusInternalServerError)
		return
	}

	// Convertir la lista de procesos a JSON
	processesJSON, err := json.Marshal(processesInterface)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Process struct {
		Pid      int
		Name     string
		Ram      int
		State    int
		User     int
		PidPadre int
	}

	var processes []Process

	// Convertir processesInterface a []Process y agregar a processes
	for _, processInterface := range processesInterface.([]interface{}) {
		processMap := processInterface.(map[string]interface{})
		var process Process
		if pid, ok := processMap["Pid"].(float64); ok {
			process.Pid = int(pid)
		}
		if name, ok := processMap["Name"].(string); ok {
			process.Name = name
		}
		if ram, ok := processMap["Ram"].(float64); ok {
			process.Ram = int(ram)
		}
		if state, ok := processMap["State"].(float64); ok {
			process.State = int(state)
		}
		if user, ok := processMap["User"].(float64); ok {
			process.User = int(user)
		}
		if pidPadre, ok := processMap["PidPadre"].(float64); ok {
			process.PidPadre = int(pidPadre)
		}
		processes = append(processes, process)
	}

	// Iterar sobre los procesos y insertarlos en la base de datos
	for _, process := range processes {
		var count int
		if process.PidPadre == 0 {
			// Verificar si el proceso ya existe en ProcesoPadre
			err = db.QueryRow("SELECT COUNT(*) FROM ProcesoPadre WHERE pid = ?", process.Pid).Scan(&count)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if count == 0 {
				// Insertar en ProcesoPadre
				_, err = db.Exec("INSERT INTO ProcesoPadre (pid, name, ram, state, user) VALUES (?, ?, ?, ?, ?)",
					process.Pid, process.Name, process.Ram, process.State, process.User)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
				// Actualizar en ProcesoPadre
				_, err = db.Exec("UPDATE ProcesoPadre SET name = ?, ram = ?, state = ?, user = ? WHERE pid = ?",
					process.Name, process.Ram, process.State, process.User, process.Pid)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		} else {
			// Verificar si el proceso ya existe en ProcesoHijo
			err = db.QueryRow("SELECT COUNT(*) FROM ProcesoHijo WHERE pid = ?", process.Pid).Scan(&count)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if count == 0 {
				// Insertar en ProcesoHijo
				_, err = db.Exec("INSERT INTO ProcesoHijo (pid, name, pidPadre, state) VALUES (?, ?, ?, ?)",
					process.Pid, process.Name, process.PidPadre, process.State)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
				// Actualizar en ProcesoHijo
				_, err = db.Exec("UPDATE ProcesoHijo SET name = ?, pidPadre = ?, state = ? WHERE pid = ?",
					process.Name, process.PidPadre, process.State, process.Pid)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		}
	}

	// Escribir la respuesta HTTP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(processesJSON)
}

var process *exec.Cmd

func StartProcess(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("bash", "-c", "./proyecto1")
	err := cmd.Start()
	if err != nil {
		http.Error(w, "Error al iniciar el proceso", http.StatusInternalServerError)
		return
	}

	// Obtener el PID del proceso hijo
	childPID := cmd.Process.Pid

	// Almacenar el cmd para futuras operaciones si es necesario
	process = cmd

	// Responder al cliente con el PID del proceso hijo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"pid": ` + strconv.Itoa(childPID) + `}`))
}

func StopProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGSTOP al proceso con el PID proporcionado stop kill -SIGSTOP PID
	cmd := exec.Command("kill", "-SIGSTOP", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, "Error al detener el proceso con PID "+strconv.Itoa(pid), http.StatusInternalServerError)
		return
	}

	// Responder al cliente con un mensaje de éxito
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Proceso con PID ` + strconv.Itoa(pid) + ` detenido"}`))
}

func ResumeProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGCONT al proceso con el PID proporcionado kill -SIGCONT PID
	cmd := exec.Command("kill", "-SIGCONT", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, "Error al reanudar el proceso con PID "+strconv.Itoa(pid), http.StatusInternalServerError)
		return
	}

	// Responder al cliente con un mensaje de éxito
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Proceso con PID ` + strconv.Itoa(pid) + ` reanudado"}`))
}

func KillProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGCONT al proceso con el PID proporcionado KILL -9 PID o KILL SIGTERM PID
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, "Error al matar el proceso con PID "+strconv.Itoa(pid), http.StatusInternalServerError)
		return
	}

	// Responder al cliente con un mensaje de éxito
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Proceso con PID ` + strconv.Itoa(pid) + ` terminado"}`))
}

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	// Obtener las credenciales de la base de datos desde las variables de entorno
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Configurar la conexión a la base de datos MySQL
	dbURI := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	db, err = sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Intentar hacer ping a la base de datos para verificar la conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	} else {
		log.Println("Conexión a la base de datos exitosa")
	}

	// Crear el enrutador HTTP
	router := mux.NewRouter()

	// Manejar la ruta POST para guardar datos de RAM
	router.HandleFunc("/ram", saveRamData).Methods("POST")

	// Manejar la ruta GET para obtener datos de RAM
	router.HandleFunc("/ram", getRamData).Methods("GET")

	// Manejar la ruta POST para guardar datos de CPU
	router.HandleFunc("/cpu", saveCpuData).Methods("POST")

	// Manejar la ruta GET para obtener datos de CPU
	router.HandleFunc("/cpu", getCpuData).Methods("GET")

	// Manejar la ruta GET para obtener todos los datos de CPU
	router.HandleFunc("/cpu/all", getAllCpuData).Methods("GET")

	// Manejar la ruta GET para obtener todos los datos de RAM
	router.HandleFunc("/ram/all", getAllRamData).Methods("GET")

	// Manejar la ruta GET para obtener los datos de los procesos
	router.HandleFunc("/procesos", handleGetProcesses).Methods("GET")

	// Manejar la ruta POST para iniciar un proceso
	router.HandleFunc("/procesos/iniciar", StartProcess).Methods("POST")

	// Manejar la ruta POST para detener un proceso
	router.HandleFunc("/procesos/detener", StopProcess).Methods("POST")

	// Manejar la ruta POST para reanudar un proceso
	router.HandleFunc("/procesos/reanudar", ResumeProcess).Methods("POST")

	// Manejar la ruta POST para matar un proceso
	router.HandleFunc("/procesos/matar", KillProcess).Methods("POST")

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},           // Permitir todas las origenes
		AllowedMethods: []string{"GET", "POST"}, // Permitir métodos GET y POST
		AllowedHeaders: []string{"*"},           // Permitir todas las cabeceras
	})

	// Iniciar el servidor HTTP con CORS habilitado
	log.Println("Servidor iniciado en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
}
