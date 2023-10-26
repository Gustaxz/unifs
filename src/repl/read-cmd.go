package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustaxz/unifs/src/unifs"
)

func ReadCommands(driverPath string) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("unifs >> ")
		scanner.Scan()
		text := scanner.Text()

		switch text {
		case "exit":
			fmt.Println("Saindo do sistema unifs...")
			os.Exit(0)
		case "clear":
			fmt.Print("\033[H\033[2J")
		case "format":
			fmt.Println("Formatando o sistema de arquivos...")
			err := unifs.FormatDrive(driverPath, 2*1024*1024)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Sistema de arquivos formatado com sucesso!")
			}
		default:
			fmt.Println("Comando não reconhecido!")
		}
	}
}
