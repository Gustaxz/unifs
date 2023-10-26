package bootSector

import (
	"encoding/binary"
	"log"
	"os"

	"github.com/gustaxz/unifs/utils"
)

// Seguindo a especificação da Microsoft (tirando os blocos Ignored), são essas as informações que precisam estar no boot sector da FAT
type BootSector struct {
	JumpBoot          [3]byte  // Instruction para pular o boot
	OemName           [8]byte  // Nome do sistema operacional que formatou o disco
	BytesPerSector    [2]byte  // Quantidade de bytes por setor
	SectorsPerCluster [1]byte  // Quantidade de setores por cluster
	ReservedSectors   [2]byte  // Quantidade de setores reservados
	NumberOfFats      [1]byte  // Quantidade de tabelas FAT
	RootEntries       [2]byte  // Quantidade de máxima entradas no diretório raiz
	TotalSectors      [2]byte  // Quantidade total de setores (No FAT12 usa-se tamanho de 512 bytes)
	Media             [1]byte  // Tipo de mídia
	MediaDescriptor   [2]byte  // Descritor de mídia (No FAT12 usa-se 0xF8)
	SectorsPerFat     [2]byte  // Quantidade de setores por FAT
	SectorsPerTrack   [2]byte  // Quantidade de setores por trilha
	NumberOfHeads     [2]byte  // Quantidade de cabeças
	HiddenSectors     [4]byte  // Quantidade de setores ocultos
	LargeTotalSectors [4]byte  // Quantidade total de setores (caso TotalSectors seja maior que 65535)
	Ignored           [2]byte  // Ignorado
	BootSignature     [1]byte  // Assinatura do boot
	VolumeId          [4]byte  // ID do volume
	VolumeLabel       [11]byte // Rótulo do volume, um nome mesmo
	FileSystemType    [8]byte  // Tipo do sistema de arquivos
}

func CreateBootSector(data BootSector, f *os.File) error {

	err := binary.Write(f, binary.LittleEndian, utils.EncodeToBytes(data))
	if err != nil {
		return err
	}

	log.Println("Boot sector created successfully!")
	return nil

}
