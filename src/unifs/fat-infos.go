package unifs

import (
	"fmt"
	"os"
	"strings"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	FAT "github.com/gustaxz/unifs/src/fat"
)

func FATInfos(f *os.File, bootSector *bootSector.BootSectorMainInfos) error {
	adresses, err := FAT.ReadFATAdresses(f, bootSector)
	if err != nil {
		return err
	}

	if len(adresses) == 0 {
		return fmt.Errorf("não há endereços na FAT")
	}

	fmt.Println("Endereços da FAT em hexadecimal:\n")
	for i, adress := range adresses {
		formatAdress := fmt.Sprintf("%x", adress)
		fmt.Printf("%d: ", i)
		fmt.Println(strings.ToUpper(formatAdress))
	}

	fmt.Println("")

	return nil

}
