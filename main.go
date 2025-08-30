/*
Copyright © 2025 IZANGILDEV HERE IZANGF2004@GMAIL.COM
*/
package main

import (
	"password-manager/cmd"
	"password-manager/utils"
)

func main() {
	//Implementar contraseña maestra

	utils.LoadPasswords(utils.PASSWORD_FILE)
	cmd.Execute()
}
