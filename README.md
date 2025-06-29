# Manicure Scheduler - Backend

## 📌 Descrição

Este é o backend da POC do projeto **Manicure Scheduler**, uma aplicação voltada à organização e agendamento de serviços de manicure. A API expõe endpoints REST para que clientes, profissionais e recepcionistas possam interagir com o sistema de agendamentos de forma segura e escalável.

Este projeto adota boas práticas de **arquitetura hexagonal (ports & adapters)**, com separação entre regras de negócio (`core`) e implementações técnicas (`infra`). Utiliza uma stack moderna com **Go**, **Chi**, **GORM** e **PostgreSQL**.

---

## 🚀 Tecnologias Utilizadas

- **Go** — Linguagem principal do projeto.
- **Chi** — Router HTTP minimalista e idiomático para Go.
- **GORM** — ORM para acesso e manipulação do banco de dados.
- **PostgreSQL** — Banco de dados relacional.
- **GoDotEnv** — Leitura de variáveis de ambiente a partir de `.env`.
- **Docker** — (Preparado para containerização via `deployments/docker`).
- **Testify** — Framework para testes unitários.
- **uuid** — Geração de UUIDs.
- **GORM Migrations** — Migrations SQL gerenciadas manualmente.

---

## 📁 Estrutura do Projeto

```plaintext
manicure-scheduler/
├── build
│   └── ci                        # Scripts de integração contínua (ex: GitHub Actions)
├── cmd
│   └── main.go                   # Entry point da aplicação
├── config
│   └── di
│       └── appointment_module.go # Injeção de dependências por módulo
├── deployments
│   └── docker                    # Configurações para execução com Docker
├── go.mod / go.sum              # Gerenciador de dependências do Go
├── internal
│   ├── core
│   │   ├── app
│   │   │   └── appointment
│   │   │       └── usecase
│   │   │           └── commands
│   │   │               └── create
│   │   │                   ├── appointment_create_input.go
│   │   │                   ├── appointment_create_usecase.go
│   │   │                   └── appointment_create_usecase_impl.go
│   │   └── domain
│   │       └── appointment
│   │           ├── appointment.go
│   │           └── appointment_repository.go
│   └── infra
│       ├── db
│       │   └── sql
│       │       └── pg
│       │           ├── gorm
│       │           │   ├── appointments
│       │           │   │   ├── appointment_entity.go
│       │           │   │   └── appointment_repository.go
│       │           │   ├── gorm_config.go
│       │           │   └── README.md
│       │           └── migrations
│       │               └── V1__create_table_appointments.sql
│       └── web
│           └── http
│               ├── handler
│               │   └── appointment
│               │       ├── appointment_handler.go
│               │       ├── request
│               │       │   └── create_appointment_dto.go
│               │       └── response
│               │           └── appointment_presenter.go
│               ├── middleware
│               │   └── recovery.go
│               ├── router
│               │   └── router.go
│               └── server.go
├── pkg
│   ├── config
│   │   └── env.go                # Leitura de variáveis de ambiente
│   └── logger                   # Logger padrão da aplicação
└── README.md
