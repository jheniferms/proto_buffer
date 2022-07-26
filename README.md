# Proto Buffer

Projeto para estudo do proto buffer com foco no uso com a linguagem go

### Precisa fazer a instalação abaixo

1. go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
2. go get google.golang.org/protobuf/encoding/protojson
2. Instalar o [Chocolatey](https://chocolatey.org/install)

### Iniciando o projeto

1. No terminal execute o comando `go mod init` caso o arquivo go.mod não exista
2. Crie o arquivo com a extensão .proto na pasta proto
3. Execute o comando `make` no terminal


### Decode e Encode via bash

Executar na pasta protoc os comandos abaixo:

Exibi os valores pelas flags: 
`cat simple.bin | protoc --decode_raw`

Exibi os valores pela mensagem:
`cat simple.bin | protoc --decode=Simple simple.proto`

Copia a estrutura dos valores exibidos pela mensagem em um arquivo:

`cat simple.bin | protoc --decode=Simple simple.proto > simple.txt`

Cria um arquivo binario:
`cat simple.txt | protoc --encode=Simple simple.proto > simple.pb`