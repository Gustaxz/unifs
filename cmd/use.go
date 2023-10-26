/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/fatih/color"
	"github.com/gustaxz/unifs/src/repl"
	"github.com/gustaxz/unifs/src/unifs"
	"github.com/spf13/cobra"
)

var drivePath string

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Coloque o nome/caminho do arquivo para usar.",
	Long:  `Use o comando 'use' mais o nome do arquivo para usar o sistema unifs nele. Caso não exista, criará um formato FAT16 no arquivo.`,
	Run: func(cmd *cobra.Command, args []string) {
		if drivePath == "" {
			color.Red("Você precisa especificar o caminho do arquivo para usar o sistema unifs!")
			return
		}

		f, b, err := unifs.OpenDrive(drivePath)

		repl.ReadCommands(f, b, drivePath, err)
	},
}

func init() {
	rootCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	useCmd.Flags().StringVarP(&drivePath, "path", "p", "", "Caminho do arquivo para usar o sistema unifs")
}
