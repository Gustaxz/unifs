package repl

import (
	"bufio"
	"fmt"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	rootDirectoryEntry "github.com/gustaxz/unifs/src/directory-entry/root"
	handleErrors "github.com/gustaxz/unifs/src/errors"
	"github.com/gustaxz/unifs/src/unifs"
)

func exit() {
	fmt.Println("Saindo do sistema unifs...")
	os.Exit(0)
}

func handleCommandsErrors(driverPath string, err error) error {
	switch err {
	case nil:
		return nil
	case handleErrors.ErrFileNotFormatted:
		fmt.Println("O arquivo não está formatado como UNIFS!")
		fmt.Print("Deseja formatar o arquivo? (y/n) ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()

		if text == "y" {
			err := unifs.FormatDrive(driverPath, 2*1024*1024)
			if err != nil {
				return err
			} else {
				fmt.Println("Sistema de arquivos formatado com sucesso!")
			}
		} else {
			exit()
		}
	default:
		return err
	}

	return nil
}

func ReadCommands(f *os.File, bootSector *bootSector.BootSectorMainInfos, driverPath string, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	err = handleCommandsErrors(driverPath, err)
	if err != nil {
		fmt.Println(err)
		exit()
	}

	fmt.Println("Sistema unifs usado com sucesso!")

	for {
		fmt.Print("unifs >> ")
		scanner.Scan()
		text := scanner.Text()

		switch text {
		case "exit":
			exit()
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
		case "delete-driver":
			fmt.Println("Deletando o driver...")
			f.Close()
			err := unifs.DeleteDriver(driverPath)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Driver deletado com sucesso!")
				exit()
			}
		case "list-root":
			fmt.Println("Listando arquivos da raiz...")

			entrys, err := rootDirectoryEntry.List(f, bootSector)
			err = handleCommandsErrors(driverPath, err)
			if err != nil {
				fmt.Println(err)
			} else {
				if len(entrys) == 0 {
					fmt.Println("Não há arquivos na raiz!")
				}

				for _, entry := range entrys {
					fmt.Println("Nome:", entry.FileName)
					fmt.Println("Tamanho:", entry.FileSize)

				}
			}
		default:
			fmt.Println("Comando não reconhecido!")
		}
	}
}
