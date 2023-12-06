# UNIFS

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)
[![forthebadge](https://forthebadge.com/images/badges/license-mit.svg)](https://forthebadge.com)

Unifs é uma versão simplificada do FAT16, usando a linguagem Go. É apenas um projeto de aprendizado para sistemas de arquivos

<img src="./docs/use.png" width="60%">

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

-   `exit` : sai do repl

<img src="./docs/exit.png" width="50%">

-   `clear` : limpa o terminal

-   `format` : formata o arquivo passado na flag `-p` para o sistema de arquivos unifs

<img src="./docs/format.png" width="50%">

-   `delete-driver` : apaga o arquivo passado na flag `-p`

<img src="./docs/delete_driver.png" width="50%">

-   `list-root` : lista os arquivos e diretórios da raiz do sistema de arquivos

<img src="./docs/list_root.png" width="50%">

-   `copy-from` : copia um arquivo do seu computador para o sistema de arquivos, seguindo o seguinte padrão

```bash

$ copy-from /path/to/file

```

<img src="./docs/copy_from.png" width="50%">

-   `copy-to` : copia um arquivo do sistema de arquivos para o seu computador, seguindo o seguinte padrão

```bash

$ copy-to /path/unifs /path/to/destination

```

<img src="./docs/copy_to.png" width="50%">
