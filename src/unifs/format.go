package unifs

import (
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	"github.com/gustaxz/unifs/src/utils"
)

func FormatDrive(drivePath string, size int64) error {
	if _, err := os.Stat(drivePath); err == nil {
		os.Remove(drivePath)
	}

	f, err := CreateEmptyDriver(drivePath, size)
	if err != nil {
		return err
	}
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

	err = bootSector.CreateBootSector(bootSectorData, f)
	if err != nil {
		return err
	}

	return nil
}
