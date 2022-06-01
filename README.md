# Pismo Payment Serrvice

O Payment Service possibilita que usuários criem contas e registrem transações.

## Informações técnicas

### Estrutura

O projeto foi desenvolvido com base no Clean Architecture, levando em consideração o [`Standard Go Project Layout`](https://github.com/golang/go/wiki/Modules) e o artigo [`Package Oriented Design`](https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html) do William Kennedy, e os commits seguiram o padrão [`Conventional Commits`](https://www.conventionalcommits.org/en/v1.0.0/).

O projeto está representado na seguinte estrutura de pastas:

```text
payment
├── cmd
│   ├── grpcapp
│   └── restapp
├── configs
├── internal
│   ├── app
│   │   ├── domain
│   │   │   └── operations
│   │   └── usecases
│   │       └── dto
│   └── platform
│       ├── handlers
│       └── repository
│           └── datastruct
└── vendor
```

- `cmd:` pasta com as principais aplicações para o projeto (não tive tempo para implementar os mesmos serviços via gRPC).
- `cmd/restapp:` a pasta da aplicação, onde encontraremos o arquivo `main.go`.
- `config:` arquivos de configuração de variáveis de ambiente.
- `internal/app:` representa o core da aplicação sem dependência alguma de pacotes externos.
- `internal/domain:` entidades ricas.
- `internal/usecases:` a porta de entrada do 'mundo externo", seja via chamadas REST, gRPC, ou qualquer outra de acesso.
- `internal/dto:` objetos anémicos utilizados para transferência de dados entre as camadas.
- `platform:` pastas utilizada para agrupar todo o 'mundo externo". Seja implementações de acesso a banco de dados, cache, filas, ect.
- `platform/handlers:` handlers de chamadas REST, separados por arquivo. É possível identificar as possíveis chamadas pelos respectivos nomes de arquivo. Foi utilizado o Gin para agilizar a implementação.
- `platform/repository:` configuração de conexão com o banco de dados e suas operações.
- `platform/repository/datastruct:` structs anêmicas mapeadas com tags `db` para facilitar a interação com o banco. NÃO FOI USADO.
- `vendor:` dependências da aplicação

### Executando

Para "buildar" e executar a aplicação de forma fácil, existe o arquivo `docker-compose.yml` na raiz do projeto. O Docker compose "builda" uma imagem Docker `FROM scratch` utilizando **multi-stage builds** e subirá um PostgreSQL e um clinente para facilitar a consulta visual no banco.

O arquivo `initdb.sql`, também na raiz do projeto, será lido no momento da subida do container do PostgreSQL, via volumes configurado no docker compose.

Também foi disponibilizado um `Makefile` com os seguintes targuets:

- `watch` inicia o reflex server para observar mudanças em arquivos '.go' e re-run o serviço (Mac ou Linux). Utilizado em momento de desenvolvivemto para não precisar reestartar manualmento a cada modificação
- `setup` prepara o ambiente para iniciar o desenvolvimento
- `run` executa o docker-compose com a flag --build
- `test` executa todos os testes de unidades

Para iniciar o serviço localmente sem o Docker, será necessário criar um arquivo `.env` com variáveis de ambiente necessárias. Tal arquivo pode ser criado executado o comando abaixo no terminal:

```bash
cat <<EOF >>.env
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=admin
DB_NAME=payment
EOF
```

### Acessando

O serviço é iniciado por padrão na porta **3000**, podendo ser consumido através dos comandos `cURL` a seguir:

#### `GET: /health`

Para checar a saúde do serviço e do banco de dados.

```bash
curl -v -X GET \
  'localhost:3000/health'
```

#### `POST: /accounts`

Criação de uma conta. Possíveis códigos http: `201 Created`, `400 BadRequest`, `422 UnprocessableEntity`, `500 InternalServerError`

```bash
curl -v -X POST \
  'localhost:3000/v1/accounts' \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "document_number":"12345678901"
    }'
```

#### `GET: /accounts/:id`

Consulta de informações de uma conta. Possíveis códigos http: `200 OK`, `400 BadRequest`, `404 NotFound`, `500 InternalServerError`

```bash
curl -v -X GET \
  'localhost:3000/v1/accounts/1' \
  --header 'Accept: application/json'
```

#### `POST: /transactions`

Criação de uma transação. Possíveis códigos http: `201 Created`, `400 BadRequest`, `422 UnprocessableEntity`, `500 InternalServerError`

```bash
curl -v -X POST \
  'localhost:3000/v1/transactions' \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "account_id": 1,
      "operation_type_id": 1,
      "amount": -29.99
    }'
```

### TODOs

- [ ] adicionar logs na camada de caso de uso
- [ ] tratar possíveis problemas de precisão decimal
- [ ] finalizar implementação dos testes
- [ ] implementar JWT middleware
