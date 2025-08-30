package utils

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var PASSWORDS = make(map[string]Credential)

const PASSWORD_FILE = "C:/Users/izang/Desktop/PROYECTOS/go-password-manager/passwords.json"
