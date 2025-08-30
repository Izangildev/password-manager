package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Comprueba si el fichero existe
func existPasswordsFile(passwordsFile string) bool {
	_, err := os.Stat(passwordsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println("Error al comprobar el archivo:", err)
		return false
	}
	return true
}

// Carga el JSON al mapa global PASSWORDS
func LoadPasswords(passwordsFile string) {
	if !existPasswordsFile(passwordsFile) {
		return
	}

	data, err := os.ReadFile(passwordsFile)
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}

	if len(data) == 0 {
		return
	}

	err = json.Unmarshal(data, &PASSWORDS)
	if err != nil {
		fmt.Println("Error al parsear JSON:", err)
	}
}

// Guarda el mapa global PASSWORDS en el JSON
func SavePasswords() {
	data, err := json.MarshalIndent(PASSWORDS, "", "  ")
	if err != nil {
		fmt.Println("Error al transformar en JSON:", err)
		return
	}

	err = os.WriteFile(PASSWORD_FILE, data, 0644)
	if err != nil {
		fmt.Println("Error al escribir el fichero:", err)
		return
	}
}
