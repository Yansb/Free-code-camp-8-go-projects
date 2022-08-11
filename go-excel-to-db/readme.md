## Desafio grupo de estudos GO

## Objetivo
Testar a fazer a aplicação em GO e em Node para comparar as facilidades 
e dificuldades de cada linguagem e no final fazermos um estudo de caso.

## Como rodar

primeiro de tudo precisamos subir a imagem postgres e criar dentro dessa imagem
um banco de dados com o nome "go"

```bash
 docker run --name postgres -p 5432:5432  -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d postgres:14-alpine
 ```
depois que o container estiver rodando e tivermos criado o banco se rodarmos o

```bash
  go run main.go
 ```

O programa deve criar os models no banco de dados, ler o arquivo excel e começar a escrever seus dados no banco
