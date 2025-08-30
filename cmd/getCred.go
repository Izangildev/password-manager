/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// getCredCmd represents the getCred command
var getCredCmd = &cobra.Command{
	Use:   "getCred",
	Short: "Muestra los detalles de una credencial específica",
	Long: `Recibe el nombre de un servicio y muestra la credencial asociada.
Incluye el usuario y la contraseña (desenmascarada).
Ejemplo: password-manager get gmail`,

	Run: func(cmd *cobra.Command, args []string) {
		service := os.Args[2]
		get(service)
	},
}

func init() {
	rootCmd.AddCommand(getCredCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCredCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCredCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
