# Go Expert: Desafio Client Server API

Sistema para armazenamento do histórico de cotação do dólar e fornecimento da cotação do dólar em tempo real.

O sistema é composto por dois componentes, denominados aqui como Cliente e Servidor.

## Servidor

O Servidor é o componente responsável por monitorar as variações na cotação do dólar através da integração com um parceiro externo, persistir estas alterações de forma histórica e fornecer uma API para retornar a cotação atual do dólar.

## Cliente

O Cliente é o componente responsável por solicitar ao Servidor a cotação atual do dólar e manter um arquivo local "./cotacao.txt"

## Requisitos

Este repositório requer que o Go esteja instalado previamente, para que seja possível a execução ou compilação do sistema.

A versão correta do Go que precisa estar instalada pode ser encontrada em [./go.mod](go.mod)

## Como executar o sistema

Devido as dependências entre os componentes, é necessário iniciar primeiro o Servidor e apenas após o servidor iniciado é que iniciaremos o Cliente.

No diretório root desde repositório, execute os seguintes comandos:

```bash
go run ./cmd/server/server.go "./config/server.json"
go run ./cmd/client/client.go "./config/client.json"
```

## Sugestão para verificação dos requisitos

Todos os parametros configuráveis são encontrados nos arquivos `./config/server.json` e `./config/client.json`.

As lógicas de timeout podem ser facilmente testadas apenas alterando os valores configurados para os limites de leitura e escrita, sugestão de valor para tests é de "1ns".
