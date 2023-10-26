package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	rootDirectoryEntry "github.com/gustaxz/unifs/src/directory-entry/root"
	handleErrors "github.com/gustaxz/unifs/src/errors"
	"github.com/gustaxz/unifs/src/files"
	"github.com/gustaxz/unifs/src/unifs"
)

func exit() {
	color.Yellow("Saindo do sistema unifs...")
	os.Exit(0)
}

func handleCommandsErrors(driverPath string, err error) error {
	switch err {
	case nil:
		return nil
	case handleErrors.ErrFileNotFormatted:
		color.Yellow("O arquivo não está formatado como UNIFS!")
		fmt.Print("Deseja formatar o arquivo? (y/n) ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()

		if text == "y" {
			err := unifs.FormatDrive(driverPath, 2*1024*1024)
			if err != nil {
				return err
			} else {
				color.Green("Sistema de arquivos formatado com sucesso!")
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
		color.Red(err.Error())
		exit()
	}

	c := color.New(color.FgGreen).Add(color.Bold)
	c.Println("\nSistema unifs inicializado com sucesso!\n")

	for {
		c := color.New(color.FgHiMagenta).Add(color.Bold)
		c.Print("@unifs: ")
		scanner.Scan()
		textScanned := scanner.Text()

		tokens := strings.Split(textScanned, " ")
		text := tokens[0]

		switch text {
		case "exit":
			exit()
		case "clear":
			fmt.Print("\033[H\033[2J")
		case "format":
			color.Yellow("Formatando o sistema de arquivos...")
			err := unifs.FormatDrive(driverPath, 2*1024*1024)
			if err != nil {
				color.Red(err.Error())
			} else {
				color.Green("Sistema de arquivos formatado com sucesso!")
			}
		case "delete-driver":
			color.Yellow("Deletando o driver...")
			f.Close()
			err := unifs.DeleteDriver(driverPath)
			if err != nil {
				color.Red(err.Error())
			} else {
				color.Green("Driver deletado com sucesso!")
				exit()
			}
		case "list-root":
			color.Yellow("Listando arquivos da raiz...")

			entrys, err := rootDirectoryEntry.List(f, bootSector)
			err = handleCommandsErrors(driverPath, err)
			if err != nil {
				color.Red(err.Error())
			} else {
				if len(entrys) == 0 {
					color.Red("Não há arquivos na raiz!")
				}

				for _, entry := range entrys {
					fmt.Println("Nome:", entry.FileName)
					fmt.Println("Tamanho:", entry.FileSize)

				}
			}
		case "copy-from":
			color.Yellow("Copiando arquivo para o sistema unifs...")
			if len(tokens) < 2 {
				color.Red("Você precisa especificar o caminho do arquivo para copiar!")
				break
			}

			originPath := tokens[1]
			err := files.CopyFrom(originPath, f, bootSector)
			err = handleCommandsErrors(driverPath, err)
			if err != nil {
				color.Red(err.Error())
			} else {
				color.Green("Arquivo copiado com sucesso!")
			}
		default:
			color.Yellow("Comando não reconhecido!")
		}
	}
}
