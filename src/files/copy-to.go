package files

import (
	"encoding/binary"
	"log"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

func CopyTo(targetPath string, fOrigin File, f *os.File, bootSector *bootSector.BootSectorMainInfos) error {
	fTatget, err := os.OpenFile(targetPath, os.O_WRONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			fTatget, err = os.Create(targetPath)
			if err != nil {
				return err
			}
			log.Println("Arquivo criado com sucesso!")
		} else {
			return err
		}
	}
	defer fTatget.Close()

	content, err := ReadFile(fOrigin, f, bootSector)
	if err != nil {
		return err
	}

	err = binary.Write(fTatget, binary.LittleEndian, content)
	if err != nil {
		return err
	}

	log.Println("Arquivo copiado com sucesso!")

	return nil
}
