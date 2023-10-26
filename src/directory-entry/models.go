package directoryEntry

/*
Uma entrada de diretório é um registro de 32 bytes que contém informações sobre um arquivo ou diretório.
Existe o Root Directory Entry que é uma parte reservada logo após a região da FAT que contém as informações
dos arquivos e diretórios que estão no diretório raiz.
*/
type DirectoryEntry struct {
	FileName            [8]byte
	FileExtension       [3]byte
	FileAttributes      byte
	Reserved            [2]byte
	CreationTime        [2]byte
	CreationDate        [2]byte
	LastAccessDate      [2]byte
	Ignored             [2]byte
	LastWriteTime       [2]byte
	LastWriteDate       [2]byte
	FirstLogicalCluster [2]byte
	FileSize            [4]byte
}

type DirectoryEntryMainInfos struct {
	FileName    string
	FileSize    uint32
	FirstSector uint16
}
