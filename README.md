# gin-api-rest

API RESTful desenvolvida em Go utilizando o framework [Gin](https://gin-gonic.com/) e o ORM [GORM](https://gorm.io/), com operações CRUD para gerenciamento de alunos.

Este projeto foi construído a partir do treinamento [Go: desenvolva uma API REST com simplicidade usando o Gin](https://cursos.alura.com.br/course/go-gin-api-rest-simplicidade) da Alura.

## Funcionalidades

- Cadastro de alunos
- Consulta de todos os alunos
- Busca de aluno por ID ou CPF
- Atualização de dados do aluno
- Remoção de aluno

## Tecnologias Utilizadas

- [Go](https://golang.org/)
- [Gin](https://gin-gonic.com/)
- [GORM](https://gorm.io/)
- [Docker](https://www.docker.com/) (opcional)
- [Swagger (gin-swagger)](https://github.com/swaggo/gin-swagger) (documentação automática da API)

## Instalação de dependências

Antes de rodar o projeto, instale as dependências:

```sh
go mod tidy
go get gopkg.in/validator.v2
go get -t github.com/andersongni/gin-api-rest
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/swag/example/celler/httputil
swag init --parseDependency --parseInternal --parseDepth 1 #Gera as documentacoes (pasta docs)
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

## Documentação automática com Swagger

O projeto utiliza o **gin-swagger** para gerar e exibir automaticamente a documentação da API REST. O Swagger permite que você visualize, teste e entenda todos os endpoints da sua API de forma interativa, facilitando o desenvolvimento e a integração com outros sistemas.

Para gerar a documentação, utilize o comando:

```sh
swag init
```

Depois, acesse a rota `/swagger/index.html` no seu servidor para visualizar a documentação interativa.

## Como rodar

1. Clone o repositório:
   ```sh
   git clone https://github.com/seu-usuario/gin-api-rest.git
   cd gin-api-rest
   ```
2. Configure o banco de dados no arquivo `.env` ou diretamente no código.
3. Rode a aplicação (recomendado para evitar alertas do Windows):
   ```sh
   go build -o main.exe; .\main.exe
   ```
   > **Nota:** Usar `go build` seguido de `main.exe` evita que o Windows mostre alertas de segurança ao rodar o binário.

   Ou, se preferir, use Docker Compose:
   ```sh
   docker-compose up
   ```

## Exemplos de uso

- **GET** `/alunos` — Lista todos os alunos
- **GET** `/alunos/:id` — Busca aluno por ID
- **GET** `/alunos/cpf/:cpf` — Busca aluno por CPF
- **POST** `/alunos` — Cria novo aluno
- **PUT** `/alunos/:id` — Atualiza aluno
- **DELETE** `/alunos/:id` — Remove aluno

## Contribuição

Pull requests são bem-vindos! Sinta-se à vontade para abrir issues e sugerir melhorias.

---

Feito com Go, Gin e GORM 🚀
