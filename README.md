# UNIFS

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/license-mit.svg)](https://forthebadge.com)

Unifs é uma versão simplificada do FAT16, usando a linguagem Go. É apenas um projeto de aprendizado para sistemas de arquivos

## Instalação

Para gerar o executável

```bash

$ git clone github.com/gustaxz/unifs
$ cd unifs
$ go build

```

Caso queira instalar o executável e adicionais nas variáveis de ambiente do go

```bash
$ git clone github.com/gustaxz/unifs
$ cd unifs
$ go build
$ go install

```

## Uso

O unifs funciona emulando um arquivo do seu computador como se fosse um driver físico. Para isso, passe o caminho do arquivo que será usado como driver na flag `-p`

```bash
$ unifs use -p /path/to/file

```

Após isso, você entrará em um repl, onde poderá usar os comandos do unifs

### Comandos

-   `help` : mostra os comandos disponíveis

-   `exit` : sai do repl

-   `clear` : limpa o terminal

-   `format` : formata o arquivo passado na flag `-p` para o sistema de arquivos unifs

-   `delete-driver` : apaga o arquivo passado na flag `-p`

-   `list-root` : lista os arquivos e diretórios da raiz do sistema de arquivos

-   `copy-from` : copia um arquivo do seu computador para o sistema de arquivos, seguindo o seguinte padrão

```bash

$ copy-from /path/to/file

```

-   `copy-to` : copia um arquivo do sistema de arquivos para o seu computador, seguindo o seguinte padrão

```bash

$ copy-to /path/unifs /path/to/destination

```

-   `read-file` : lê um arquivo do sistema de arquivos, seguindo o seguinte padrão

```bash

$ read-file /path/unifs

```

-   `delete-file` : apaga um arquivo do sistema de arquivos, seguindo o seguinte padrão

```bash

$ delete-file /path/unifs

```

-   `hexdump` : mostra o conteúdo de um arquivo em hexadecimal, seguindo o seguinte padrão

```bash

$ hexdump /path/unifs

```

-   `infos` : mostra informações sobre o sistema de arquivos

-   `fat` : mostra a tabela FAT do sistema de arquivos

### Licença

Esse projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.
