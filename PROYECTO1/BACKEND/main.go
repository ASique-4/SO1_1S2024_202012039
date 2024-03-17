package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"time"
)

type RamData struct {
	Total      uint64 `json:"totalRam"`
	EnUso      uint64 `json:"memoriaEnUso"`
	Porcentaje uint64 `json:"porcentaje"`
	Libre      uint64 `json:"libre"`
}

var db *sql.DB

func saveRamData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Ejecutar el comando cat
	out, err := exec.Command("cat", "/proc/ram_so1_1s2024").Output()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Decodificar el JSON
	var ram RamData
	err = json.Unmarshal(out, &ram)
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
}

func main() {
	// Conexión a la base de datos
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/ram", saveRamData)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
