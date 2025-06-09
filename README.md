# Manicure Scheduler - Backend

## Descrição

Este é o backend do projeto **Manicure Scheduler**, uma aplicação voltada à organização e agendamento de serviços de manicure. A API expõe endpoints REST para que clientes, profissionais e recepcionistas possam interagir com o sistema de agendamentos de forma segura, fluida e escalável.

Este projeto segue boas práticas de arquitetura e organização de código, com separação clara entre `core` (regras de negócio) e `infra` (implementações técnicas), utilizando uma stack moderna com **Golang**, **Fiber** e **GORM**.

---

## Tecnologias Utilizadas

- **Go** — Linguagem principal.
- **Fiber** — Framework web rápido e minimalista para criação da API REST.
- **GORM** — ORM para conexão e manipulação de dados no PostgreSQL.
- **PostgreSQL** — Banco de dados relacional.
- **Swagger (swaggo)** — Geração automática da documentação da API.
- **Docker** — Containerização da aplicação.
- **GoDotEnv** — Carregamento de variáveis de ambiente.
- **Migrations (gormigrate)** — Gerenciamento de versionamento do banco.

---

## Estrutura Inicial do Projeto

```plaintext
manicure-scheduler/
├── build/
│
├── cmd/
│   └── main.go                        # Entrypoint principal
│
├── internal/
│   ├── core/
│   │   ├── application/              # Casos de uso (comandos e queries)
│   │   └── domain/                   # Entidades e interfaces (DDD)
│   │
│   └── infra/
│       ├── dataprovider/
│       │   └── sql/
│       │       └── pg/
│       │           ├── gorm/         # Conexão GORM e repositórios
│       │           └── migrations/   # Scripts e versionamento do banco
│       │
│       └── webserver/
│           └── http/
│               ├── handlers/         # Controladores HTTP
│               └── router.go
│
├── pkg/
│   └── config/
│       └── env.go                    # Leitura de variáveis de ambiente
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
