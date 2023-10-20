package main

import (
	"bytes"
	"encoding/binary"
	"os"
	"reflect"
)

func StringToBytes(s string, size int) []byte {
	var buf bytes.Buffer

	for i := 0; i < size; i++ {
		if i < len(s) {
			buf.WriteByte(s[i])
		} else {
			buf.WriteByte(0x20)
		}
	}

	return buf.Bytes()
}

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	values := reflect.ValueOf(p)
	for i := 0; i < values.NumField(); i++ {
		err := binary.Write(&buf, binary.LittleEndian, values.Field(i).Interface())
		if err != nil {
			panic("Error encoding ID:" + err.Error())
		}
	}

	return buf.Bytes()
}

type FileContent struct {
	Name  [8]byte
	Age   [2]byte
	Hobby [8]byte
}

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

func main() {
	filePath := "mydriver"

	os.Remove(filePath)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sizeInBytes := 2 * 1024 * 1024
	if err := file.Truncate(int64(sizeInBytes)); err != nil {
		panic(err)
	}

	data := BootSector{
		JumpBoot:          [3]byte{0xEB, 0x3C, 0x90},
		OemName:           [8]byte(StringToBytes("UNIFS.0", 8)),
		BytesPerSector:    [2]byte{0x20, 0x00}, // 0x200 em hexadecimal corresponde a 512 bytes em decimal
		SectorsPerCluster: [1]byte{0x01},
		ReservedSectors:   [2]byte{0x00, 0x00},
		NumberOfFats:      [1]byte{0x01},
		RootEntries:       [2]byte{0x14, 0x00}, // 0x14 em hexadecimal corresponde a 20 em decimal
		TotalSectors:      [2]byte{0x10, 0x00}, // 0x1000 em hexadecimal corresponde a 4096 em decimal. 4096 * 512 = 2.097.152 bytes ou 2MB de espaço total
		Media:             [1]byte{0xF8},
		MediaDescriptor:   [2]byte{0x00, 0x00},
		SectorsPerFat:     [2]byte{0x00, 0x0C},
		/*  0x01 em hexadecimal corresponde a 1 em decimal.
		A tabela FAT guarda o endereço de cada setor do disco. Tendo 4096 setores, precisamos de 12 bits para endereçar cada um deles.
		Como 12 bits equivalem a 1,5 bytes, multiplicando 1,5 por 4096, temos 6144 bytes para alocar uma tabela FAT.
		Como cada setor tem 512 bytes, dividimos 6144 por 512 e temos 12 (0x000C) setores para alocar a tabela FAT. */
		SectorsPerTrack:   [2]byte{0x00, 0x00},
		NumberOfHeads:     [2]byte{0x00, 0x00},
		HiddenSectors:     [4]byte{0x00, 0x00, 0x00, 0x00},
		LargeTotalSectors: [4]byte{0x00, 0x00, 0x00, 0x00},
		Ignored:           [2]byte{0x00, 0x00},
		BootSignature:     [1]byte{0x29},
		VolumeId:          [4]byte{0x00, 0x00, 0x00, 0x00},
		VolumeLabel:       [11]byte(StringToBytes("UNIFSYS", 11)),
		FileSystemType:    [8]byte(StringToBytes("FAT12", 8)),
	}

	err = binary.Write(file, binary.LittleEndian, EncodeToBytes(data))

	if err != nil {
		panic(err)
	}

}
