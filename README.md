
---

# Store API

A **Store API** é uma aplicação em GoLang que fornece operações de CRUD para a entidade "Produto", além de incluir autenticação JWT, autorização e logging.

## Sumário

- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Configuração](#configuração)
- [Executando a API](#executando-a-api)
- [Endpoints](#endpoints)
- [Testando a API](#testando-a-api)
- [Contribuindo](#contribuindo)
- [Licença](#licença)

## Pré-requisitos

- Go 1.20 ou superior
- Git
- Postman ou cURL para testar a API

## Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/store-api.git
   cd store-api
   ```

2. Inicie o módulo Go e instale as dependências:

   ```bash
   go mod tidy
   ```

## Configuração

1. **Variáveis de Ambiente**:

   Certifique-se de definir a variável de ambiente `JWT_SECRET` que será usada para assinar os tokens JWT.

   ```bash
   export JWT_SECRET="sua-chave-secreta-aqui"
   ```

## Executando a API

Para iniciar o servidor, execute:

```bash
go run cmd/main.go
```

A API será executada em `http://localhost:8080`.

## Endpoints

### Autenticação

- **Gerar Token JWT**:

  Faça uma requisição `POST` para `/login` (Endpoint fictício para gerar token):

  ```bash
  curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "user", "password": "password"}'
  ```

  Retorna um token JWT que deve ser usado nos próximos requests.

### Produtos

- **Criar Produto**:

  ```bash
  POST /products
  ```

  - Corpo da requisição:

    ```json
    {
      "name": "Produto X",
      "description": "Descrição do produto X",
      "price": 100.0
    }
    ```

  - Exemplo:

    ```bash
    curl -X POST http://localhost:8080/products -H "Authorization: Bearer <seu-token>" -H "Content-Type: application/json" -d '{"name":"Produto X", "description":"Descrição", "price":100.0}'
    ```

- **Listar Produtos**:

  ```bash
  GET /products
  ```

  - Exemplo:

    ```bash
    curl -X GET http://localhost:8080/products -H "Authorization: Bearer <seu-token>"
    ```

- **Obter Produto por ID**:

  ```bash
  GET /products/{id}
  ```

  - Exemplo:

    ```bash
    curl -X GET http://localhost:8080/products/1 -H "Authorization: Bearer <seu-token>"
    ```

- **Atualizar Produto**:

  ```bash
  PUT /products/{id}
  ```

  - Corpo da requisição:

    ```json
    {
      "name": "Produto Y",
      "description": "Descrição do produto Y",
      "price": 150.0
    }
    ```

  - Exemplo:

    ```bash
    curl -X PUT http://localhost:8080/products/1 -H "Authorization: Bearer <seu-token>" -H "Content-Type: application/json" -d '{"name":"Produto Y", "description":"Nova descrição", "price":150.0}'
    ```

- **Excluir Produto**:

  ```bash
  DELETE /products/{id}
  ```

  - Exemplo:

    ```bash
    curl -X DELETE http://localhost:8080/products/1 -H "Authorization: Bearer <seu-token>"
    ```

## Testando a API

1. **Testes Manuais**:
   - Utilize o [Postman](https://www.postman.com/) ou cURL para realizar chamadas aos endpoints da API conforme descrito acima.

2. **Testes Automatizados**:
   - Você pode utilizar ferramentas de testes Go como `testing` para escrever e executar testes unitários. Integre esses testes com a pipeline CI/CD para garantir a qualidade do código.

## Contribuindo

Se você quiser contribuir para este projeto:

1. Faça um fork do projeto.
2. Crie uma nova branch (`git checkout -b feature/nova-feature`).
3. Commit suas mudanças (`git commit -am 'Add nova feature'`).
4. Faça um push para a branch (`git push origin feature/nova-feature`).
5. Crie um Pull Request.

## Licença

Este projeto está licenciado sob a licença MIT - consulte o arquivo [LICENSE](LICENSE) para mais detalhes.

---
