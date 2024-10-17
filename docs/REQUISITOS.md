# Requisitos

Os requisitos para cumprir este desafio são:

Construir dois sistemas em Go:

- client.go
  - deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
  - receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON).
    - o timeout máximo de 300ms para receber o resultado do server.go.
  - salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
- server.go
  - deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL 
    - o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms
  - deverá retornar no formato JSON o resultado para o cliente.
  - deverá registrar no banco de dados SQLite cada cotação recebida
    - o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
  - a porta a ser utilizada pelo servidor HTTP será a 8080
  - o endpoint necessário gerado será: /cotacao
- ambos
  - retornar erro nos logs caso o tempo de execução seja insuficiente.
