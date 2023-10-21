package handleDriver

import (
	"fmt"
	"os"
)

func CreateFileAllocationTable(f *os.File, bootSector BootSectorMainInfos) {
	sizeOfSector := bootSector.BytesPerSector
	//sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	//sizeOfFatInBytes := sizeOfFat * 8

	fmt.Print("Size of sector in bytes: ", sizeOfSector, "\n")
	fmt.Print("Size of FAT in bytes: ", bootSector.SectorsPerFat, "\n")
	// fmt.Println("Size of FAT in bytes:", sizeOfFatInBytes)
	// fmt.Println("Size of FAT in bits:", sizeOfFat)
}
