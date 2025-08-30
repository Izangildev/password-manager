/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// addCredCmd represents the addCred command
var addCredCmd = &cobra.Command{
	Use:   "addCred",
	Short: "Añade una credencial",
	Long: `Añade una nueva credencial al gestor de contraseñas.
Recibe tres parámetros obligatorios:
1. service  → nombre del servicio (ej: gmail)
2. username → usuario o email asociado al servicio
3. password → contraseña a guardar`,

	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		service := os.Args[2]
		username := os.Args[3]
		password := os.Args[4]

		add(service, username, password)
	},
}

func init() {
	rootCmd.AddCommand(addCredCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCredCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCredCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
