package unifs

import (
	"fmt"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	rootDirectoryEntry "github.com/gustaxz/unifs/src/directory-entry/root"
	FAT "github.com/gustaxz/unifs/src/fat"
)

func DriverInfos(f *os.File, bootSector *bootSector.BootSectorMainInfos) error {
	driverSize := int64(bootSector.TotalSectors) * int64(bootSector.BytesPerSector)
	sizeOfAllFiles := 0

	fmt.Println("Informações do driver ", bootSector.VolumeLabel)
	fmt.Println("\nSistema que formatou o drive: ", bootSector.OemName)
	fmt.Printf("Tamanho do setor: %d bytes\n", bootSector.BytesPerSector)
	fmt.Printf("Número de entradas no Root Directory: %d\n", bootSector.RootEntries)
	fmt.Printf("Número total de setores: %d\n", bootSector.TotalSectors)
	fmt.Printf("Número de setores por FAT: %d\n", bootSector.SectorsPerFat)
	fmt.Printf("Tamanho do driver: %d bytes (%d KB)\n", driverSize, driverSize/(1024))

	rootEntries, err := rootDirectoryEntry.List(f, bootSector)
	if err != nil {
		return err
	}

	for _, entry := range rootEntries {
		sizeOfAllFiles += int(entry.FileSize)
	}
	fmt.Printf("Espaço utilizado por arquivos: %d bytes (%d KB)\n", sizeOfAllFiles, sizeOfAllFiles/(1024))

	utilyzedRootEntries := len(rootEntries)
	fmt.Printf("Entradas disponíveis no diretório raiz: %d \n", int(bootSector.RootEntries)-utilyzedRootEntries)

	emptyFatEntries, err := FAT.ListOfEmptyAdressesFAT(f, bootSector)
	if err != nil {
		return err
	}
	emptyFatEntriesSize := len(emptyFatEntries)
	availableSpace := emptyFatEntriesSize * int(bootSector.BytesPerSector)
	fmt.Printf("Espaço disponível: %d bytes (%d KB) \n", availableSpace, availableSpace/(1024))

	return nil
}
