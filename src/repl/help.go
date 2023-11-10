package repl

import (
	"fmt"
	"reflect"
)

type CommandHelp struct {
	Exit         string
	Clear        string
	Format       string
	DeleteDriver string
	ListRoot     string
	CopyFrom     string
	CopyTo       string
	ReadFile     string
	DeleteFile   string
	Hexdump      string
	Infos        string
	FAT          string
}

// TODO - Tentar usar aquelas strings tipo pra json 'cmd: exit', algo assim, pra deixar mais bonito
func (c *CommandHelp) Init() {
	c.Exit = "exit - Sair do sistema unifs"
	c.Clear = "clear - Limpar a tela"
	c.Format = "format - Formatar o sistema de arquivos em uso"
	c.DeleteDriver = "delete-driver - Deletar o driver em uso"
	c.ListRoot = "list-root - Listar os arquivos e diretórios da raiz"
	c.CopyFrom = "copy-from - Copiar um arquivo do sistema de arquivos para o sistema de arquivos do computador. Exemplo: copy-from /arquivo.txt"
	c.CopyTo = "copy-to - Copiar um arquivo do sistema de arquivos do computador para o sistema de arquivos. Exemplo: copy-to ARQUIVO TXT /arquivo.txt. ARQUIVO TXT é o arquivo dentro do unifs, e /arquivo.txt é o arquivo no computador."
	c.ReadFile = "read-file - Ler um arquivo do sistema de arquivos. Exemplo: read-file ARQUIVO TXT. ARQUIVO TXT é o arquivo dentro do unifs."
	c.DeleteFile = "delete-file - Deletar um arquivo do sistema de arquivos. Exemplo: delete-file ARQUIVO TXT. ARQUIVO TXT é o arquivo dentro do unifs."
	c.Hexdump = "hexdump - Mostrar o conteúdo de um arquivo em hexadecimal. Exemplo: hexdump ARQUIVO TXT. ARQUIVO TXT é o arquivo dentro do unifs."
	c.Infos = "infos - Mostrar informações do sistema de arquivos."
	c.FAT = "fat - Mostrar os endereços da FAT ocupados em hexadecimal."
}

func (c *CommandHelp) Print() {

	fmt.Println()
	val := reflect.ValueOf(c).Elem()
	valNumFields := val.NumField()
	for i := 0; i < valNumFields; i++ {
		field := val.Field(i)

		fmt.Println(field)
		fmt.Println()

	}
}
