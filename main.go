/*
Copyright © 2025 IZANGILDEV HERE IZANGF2004@GMAIL.COM
*/
package main

import (
	"fmt"
	"os"
	"password-manager/cmd"
	"password-manager/security"
	"password-manager/utils"
)

func main() {
	//Implementar contraseña maestra

	// Master password --> password123
	hashed := security.GetHashedPassword()
	master := security.GetMasterPassword()
	if !security.CheckMasterPassword(hashed, master) {
		fmt.Println("Contraseña incorrecta... Saliendo del programa.")
		os.Exit(0)
	}
	fmt.Println("Contraseña correcta!")

	utils.LoadPasswords(utils.PASSWORD_FILE)
	cmd.Execute()
}
