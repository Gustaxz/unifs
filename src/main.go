package main

import (
	"fmt"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	handleDriver "github.com/gustaxz/unifs/src/handle-driver"
	saveFile "github.com/gustaxz/unifs/src/save-file"
	"github.com/gustaxz/unifs/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func formatDrive() {
	os.Remove("mydriver")

	f, err := handleDriver.CreateEmptyDriver()
	check(err)
	defer f.Close()

	// Tirando o Boot Sector, o Root Directory e a FAT, temos 4076 setores ou cerca de 2.039 KB para armazenar dados

	bootSectorData := bootSector.BootSector{
		JumpBoot:          [3]byte{0xEB, 0x3C, 0x90},
		OemName:           [8]byte(utils.StringToBytes("UNIFS.0", 8)),
		BytesPerSector:    [2]byte{0x02, 0x00}, // 0x200 em hexadecimal corresponde a 512 bytes em decimal
		SectorsPerCluster: [1]byte{0x01},
		ReservedSectors:   [2]byte{0x00, 0x00},
		NumberOfFats:      [1]byte{0x01},
		RootEntries:       [2]byte{0x00, 0x10}, // 0x10 em hexadecimal corresponde a 16 em decimal
		TotalSectors:      [2]byte{0x10, 0x00}, // 0x1000 em hexadecimal corresponde a 4096 em decimal. 4096 * 512 = 2.097.152 bytes ou 2MB de espaço total
		Media:             [1]byte{0xF8},
		MediaDescriptor:   [2]byte{0x00, 0x00},
		SectorsPerFat:     [2]byte{0x00, 0x10},
		/*
			A tabela FAT guarda o endereço de cada setor do disco. Tendo 4096 setores, precisamos de 16 bits para endereçar cada um deles.
			Como 16 bits equivalem a 2 bytes, multiplicando 2 por 4096, temos 8192 bytes para alocar uma tabela FAT.
			Como cada setor tem 512 bytes, dividimos 8192 por 512 e temos 16 (0x0010) setores para alocar a tabela FAT. */
		SectorsPerTrack:   [2]byte{0x00, 0x00},
		NumberOfHeads:     [2]byte{0x00, 0x00},
		HiddenSectors:     [4]byte{0x00, 0x00, 0x00, 0x00},
		LargeTotalSectors: [4]byte{0x00, 0x00, 0x00, 0x00},
		Ignored:           [2]byte{0x00, 0x00},
		BootSignature:     [1]byte{0x29},
		VolumeId:          [4]byte{0x00, 0x00, 0x00, 0x00},
		VolumeLabel:       [11]byte(utils.StringToBytes("UNIFSYS", 11)),
		FileSystemType:    [8]byte(utils.StringToBytes("FAT16", 8)),
	}

	err = bootSector.CreateBootSector(&bootSectorData, f)
	check(err)
}

func initDrive() (*os.File, *bootSector.BootSectorMainInfos) {
	f, err := os.Open("mydriver")
	check(err)

	bootSector, err := bootSector.ReadBootSector(f)
	check(err)
	fmt.Println(bootSector)
	f.Close()

	f, err = os.OpenFile("mydriver", os.O_RDWR, 0644)
	check(err)

	return f, bootSector
}

func main() {

	// formatDrive()

	f, bootSector := initDrive()
	defer f.Close()

	// saveFile.SaveFile(saveFile.File{
	// 	Name: [8]byte(utils.StringToBytes("TESTE", 8)),
	// 	Ext:  [3]byte(utils.StringToBytes("TXT", 3)),
	// 	Data: []byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam quis metus non dui scelerisque hendrerit ac in arcu. Nunc id malesuada eros. Praesent cursus nisi in elit tristique, id scelerisque ipsum congue. Cras vitae fringilla ligula, id bibendum purus. Aenean lacinia fringilla orci nec dictum. Maecenas sit amet justo sit amet massa mattis semper ut sed leo. Phasellus tincidunt velit sit amet odio bibendum euismod. Integer at sem vitae purus varius feugiat. Curabitur eu arcu vel odio dictum feugiat ac in est. Aliquam erat volutpat. Vestibulum auctor nisl vel efficitur tincidunt. Sed consequat eget turpis at sollicitudin. Donec vel vulputate sapien. Aenean elementum justo sit amet urna vulputate, vel bibendum arcu malesuada.

	// 	Vivamus feugiat scelerisque arcu, id consequat massa luctus eget. Vivamus convallis elit ut felis facilisis, eu mattis velit interdum. Morbi viverra hendrerit magna a accumsan. Praesent et tellus massa. Sed ullamcorper feugiat ligula eu gravida. Maecenas eget viverra lectus, id tristique velit. Proin in purus sed ex dignissim luctus. Sed a bibendum velit. Phasellus tempus arcu eget ex elementum, vel viverra massa vestibulum. Curabitur vulputate odio sit amet elit varius, nec dapibus tortor interdum.

	// 	Nunc euismod id odio eget euismod. Fusce a nunc in massa dapibus facilisis. Sed in enim nisl. Sed eget enim ac sem condimentum scelerisque a et sapien. Vivamus tristique nec metus eu cursus. Nulla facilisi. In hac habitasse platea dictumst. Suspendisse ut neque id arcu elementum eleifend in nec dolor. Suspendisse eget lacinia erat, nec luctus libero. Nulla facilisi. Quisque vel quam eget est rhoncus iaculis. Etiam nec odio vel mi fringilla interdum. Curabitur auctor hendrerit mauris ac lacinia. Ut quis tincidunt purus.

	// 	Fusce vitae massa nec purus tincidunt vehicula. Curabitur tincidunt vel justo vel tincidunt. In id sollicitudin neque. Nullam sit amet mattis augue. Phasellus id risus eget nulla scelerisque tempus. Proin nec erat vel ipsum facilisis rhoncus. Sed malesuada nec odio ut sollicitudin. Fusce ac orci lacinia, accumsan tortor at, dictum mi. Aliquam ullamcorper orci id massa efficitur volutpat. Etiam et sapien at nisi tincidunt auctor id eu est. Vivamus eu risus justo. Vestibulum et erat vel arcu facilisis dignissim. Sed lacinia volutpat bibendum.
	// 	`),
	// }, f, bootSector)
	// saveFile.SaveFile(saveFile.File{
	// 	Name: [8]byte(utils.StringToBytes("FELLIPE", 8)),
	// 	Ext:  [3]byte(utils.StringToBytes("TXT", 3)),
	// 	Data: []byte(`Larem ipsum dalar sit amet, cansectetur adipiscing elit. Nullam quis metus nan dui scelerisque hendrerit ac in arcu. Nunc id malesuada eras. Praesent cursus nisi in elit tristique, id scelerisque ipsum cangue. Cras vitae fringilla ligula, id bibendum purus. Aenean lacinia fringilla arci nec dictum. Maecenas sit amet justa sit amet massa mattis semper ut sed lea. Phasellus tincidunt velit sit amet adia bibendum euismad. Integer at sem vitae purus varius feugiat. Curabitur eu arcu vel adia dictum feugiat ac in est. Aliquam erat valutpat. Vestibulum auctar nisl vel efficitur tincidunt. Sed cansequat eget turpis at sallicitudin. Danec vel vulputate sapien. Aenean elementum justa sit amet urna vulputate, vel bibendum arcu malesuada.

	// 	Vivamus feugiat scelerisque arcu, id cansequat massa luctus eget. Vivamus canvallis elit ut felis facilisis, eu mattis velit interdum. Marbi viverra hendrerit magna a accumsan. Praesent et tellus massa. Sed ullamcarper feugiat ligula eu gravida. Maecenas eget viverra lectus, id tristique velit. Prain in purus sed ex dignissim luctus. Sed a bibendum velit. Phasellus tempus arcu eget ex elementum, vel viverra massa vestibulum. Curabitur vulputate adia sit amet elit varius, nec dapibus tartar interdum.

	// 	Nunc euismad id adia eget euismad. Fusce a nunc in massa dapibus facilisis. Sed in enim nisl. Sed eget enim ac sem candimentum scelerisque a et sapien. Vivamus tristique nec metus eu cursus. Nulla facilisi. In hac habitasse platea dictumst. Suspendisse ut neque id arcu elementum eleifend in nec dalar. Suspendisse eget lacinia erat, nec luctus libera. Nulla facilisi. Quisque vel quam eget est rhancus iaculis. Etiam nec adia vel mi fringilla interdum. Curabitur auctar hendrerit mauris ac lacinia. Ut quis tincidunt purus.

	// 	Fusce vitae massa nec purus tincidunt vehicula. Curabitur tincidunt vel justa vel tincidunt. In id sallicitudin neque. Nullam sit amet mattis augue. Phasellus id risus eget nulla scelerisque tempus. Prain nec erat vel ipsum facilisis rhancus. Sed malesuada nec adia ut sallicitudin. Fusce ac arci lacinia, accumsan tartar at, dictum mi. Aliquam ullamcarper arci id massa efficitur valutpat. Etiam et sapien at nisi tincidunt auctar id eu est. Vivamus eu risus justa. Vestibulum et erat vel arcu facilisis dignissim. Sed lacinia valutpat bibendum.
	// 	`),
	// }, f, bootSector)

	saveFile.ReadFile(saveFile.File{
		Name: [8]byte(utils.StringToBytes("FELLIPE", 8)),
		Ext:  [3]byte(utils.StringToBytes("TXT", 3)),
	}, f, bootSector)

}
