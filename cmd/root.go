/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "unifs",
	Short: "Uma implementação FAT16 simplificada!",
	Long:  `Use o comando 'use' mais o nome do arquivo para usar o sistema unifs nele. Caso não exista, criará um formato FAT16 no arquivo.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
