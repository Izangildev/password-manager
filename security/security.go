package security

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Devuelve la contraseña maestra hasheada desde el JSON
func GetHashedPassword() string {
	data, err := os.ReadFile("hashedMasterPassword.json")
	if err != nil {
		fmt.Println("Error al leer el fichero:", err)
		return ""
	}

	var result map[string]string
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error al parsear JSON:", err)
		return ""
	}

	return result["masterPassword"]
}

// Pide al usuario la contraseña maestra por stdin
func GetMasterPassword() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduce la contraseña maestra: ")
	pass, _ := reader.ReadString('\n')
	pass = strings.TrimSpace(pass)
	return pass
}

// Compara contraseña introducida con la hasheada
func CheckMasterPassword(hashedPassword, masterPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(masterPassword))
	return err == nil
}
