/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// deleteCredCmd represents the deleteCred command
var deleteCredCmd = &cobra.Command{
	Use:   "deleteCred",
	Short: "Elimina una credencial",
	Long: `Recibe el nombre de un servicio y elimina la credencial correspondiente.
Esta operación es irreversible, así que úsala con cuidado.
Ejemplo: password-manager delete gmail`,
	Run: func(cmd *cobra.Command, args []string) {
		service := os.Args[2]
		remove(service)
	},
}

func init() {
	rootCmd.AddCommand(deleteCredCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCredCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCredCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
