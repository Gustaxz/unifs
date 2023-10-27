package unifs

import (
	"fmt"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

func DriverInfos(bootSector *bootSector.BootSectorMainInfos) {
	driverSize := int64(bootSector.TotalSectors) * int64(bootSector.BytesPerSector)

	fmt.Println("Informações do driver ", bootSector.VolumeLabel)
	fmt.Println("\nSistema que formatou o drive: ", bootSector.OemName)
	fmt.Printf("Tamanho do setor: %d bytes\n", bootSector.BytesPerSector)
	fmt.Printf("Número de entradas no Root Directory: %d\n", bootSector.RootEntries)
	fmt.Printf("Número total de setores: %d\n", bootSector.TotalSectors)
	fmt.Printf("Número de setores por FAT: %d\n", bootSector.SectorsPerFat)
	fmt.Printf("Tamanho do driver: %d bytes (%d MB)\n", driverSize, driverSize/(1024*1024))
	// TODO: Espaço utilizado por arquivos e espaço realmente disponível

}
