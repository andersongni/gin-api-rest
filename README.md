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

## Como rodar

1. Clone o repositório:
   ```sh
   git clone https://github.com/seu-usuario/gin-api-rest.git
   cd gin-api-rest
   ```
2. Configure o banco de dados no arquivo `.env` ou diretamente no código.
3. Instale as dependências:
   ```sh
   go mod tidy
   ```
4. Rode a aplicação:
   ```sh
   go run main.go
   ```
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
