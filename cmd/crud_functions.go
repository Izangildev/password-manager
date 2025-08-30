package cmd

import (
	"fmt"
	"password-manager/utils"
	"strings"
)

func existService(service string) bool {
	_, exists := utils.PASSWORDS[service]
	return exists
}

// Añadir credencial
func add(service, username, password string) {
	if existService(service) {
		fmt.Println("Ya existen credenciales para", service)
		return
	}

	utils.PASSWORDS[service] = utils.Credential{
		Username: username,
		Password: password,
	}
	fmt.Println("Credenciales de", service, "creadas con éxito!")
	utils.SavePasswords()
}

func list() {
	fmt.Printf("%-15s %-20s %-s\n", "SERVICE", "USERNAME", "PASSWORD")
	fmt.Println("-------------------------------------------------")
	for service, cred := range utils.PASSWORDS {
		masked := strings.Repeat("*", len(cred.Password))
		fmt.Printf("%-15s %-20s %s\n", service, cred.Username, masked)
	}
}

func get(service string) {
	cred, exists := utils.PASSWORDS[service]
	if exists {
		fmt.Println("Usuario:", cred.Username)
		fmt.Println("Contraseña:", cred.Password)
	} else {
		fmt.Println("No hay credenciales guardadas para", service)
	}
}

func remove(service string) {
	_, exists := utils.PASSWORDS[service]
	if exists {
		delete(utils.PASSWORDS, service)
		utils.SavePasswords()
		fmt.Println("Credenciales de", service, "eliminadas.")
	} else {
		fmt.Println("No hay credenciales guardadas para", service)
	}
}
