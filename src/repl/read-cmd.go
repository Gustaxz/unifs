package repl

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	rootDirectoryEntry "github.com/gustaxz/unifs/src/directory-entry/root"
	handleErrors "github.com/gustaxz/unifs/src/errors"
	"github.com/gustaxz/unifs/src/files"
	"github.com/gustaxz/unifs/src/unifs"
	"github.com/gustaxz/unifs/src/utils"
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
		case "copy-to":
			color.Yellow("Copiando arquivo do sistema unifs...")
			if len(tokens) < 4 {
				color.Red(
					"São necessarios 4 argumentos para copiar um arquivo!\nExemplo: copy-to arquivo-alvo.txt NOME_ARQUIVO EXTENSAO_ARQUIVO\nOs útimos dois argumentos são referentes ao nome e extensão dentro do unifs!")
				break
			}

			targetPath := tokens[1]
			fileName := tokens[2]
			fileExt := tokens[3]

			fOrigin := files.File{
				Name: [8]byte(utils.StringToBytes(fileName, 8)),
				Ext:  [3]byte(utils.StringToBytes(fileExt, 3)),
				Data: []byte{},
			}

			err := files.CopyTo(targetPath, fOrigin, f, bootSector)
			err = handleCommandsErrors(driverPath, err)
			if err != nil {
				color.Red(err.Error())
			} else {
				color.Green("Arquivo copiado com sucesso!")
			}
		case "read-file":
			color.Yellow("Lendo arquivo do sistema unifs...")
			if len(tokens) < 3 {
				color.Red(
					"São necessarios 3 argumentos para copiar um arquivo!\nExemplo: copy-to NOME_ARQUIVO EXTENSAO_ARQUIVO\nOs útimos dois argumentos são referentes ao nome e extensão dentro do unifs!")
				break
			}

			fileName := tokens[1]
			fileExt := tokens[2]

			file := files.File{
				Name: [8]byte(utils.StringToBytes(fileName, 8)),
				Ext:  [3]byte(utils.StringToBytes(fileExt, 3)),
				Data: []byte{},
			}

			data, _, err := files.ReadFile(file, f, bootSector)
			err = handleCommandsErrors(driverPath, err)
			if err != nil {
				color.Red(err.Error())
			} else {
				color.Green("Arquivo lido com sucesso!")
				fmt.Println(string(data))
			}
		case "hexdump":
			color.Yellow("Lendo arquivo do sistema unifs...")
			if len(tokens) < 3 {
				color.Red(
					"São necessarios 3 argumentos para copiar um arquivo!\nExemplo: copy-to NOME_ARQUIVO EXTENSAO_ARQUIVO\nOs útimos dois argumentos são referentes ao nome e extensão dentro do unifs!")
				break
			}

			fileName := tokens[1]
			fileExt := tokens[2]

			file := files.File{
				Name: [8]byte(utils.StringToBytes(fileName, 8)),
				Ext:  [3]byte(utils.StringToBytes(fileExt, 3)),
				Data: []byte{},
			}

			data, _, err := files.ReadFile(file, f, bootSector)
			err = handleCommandsErrors(driverPath, err)
			if err != nil {
				color.Red(err.Error())
			} else {
				color.Green("Arquivo lido com sucesso!")
				fmt.Println(hex.Dump(data))
			}
		case "infos":
			unifs.DriverInfos(bootSector)
		default:
			color.Yellow("Comando não reconhecido!")
		}
	}
}
