# 🧬 Digital Monster API · Tamagotchi moderno em Go!

![Go Version](https://img.shields.io/badge/Go-1.22-blue.svg)
![SQLite](https://img.shields.io/badge/SQLite-lightblue.svg)
![Gin](https://img.shields.io/badge/Framework-Gin-green.svg)
![Clean Architecture](https://img.shields.io/badge/Design-Clean--Architecture-yellow.svg)
![Status](https://img.shields.io/badge/status-ativo-success)

> Uma API 100% em Golang que simula um Tamagotchi moderno, seguindo padrões de **Clean Architecture**, **DDD**, **Event-Driven Design** e com testes idiomáticos em Go.

---

## 📐 Arquitetura

```
internal/
├── entity/       → Regras de negócio (Monster, métodos e eventos)
├── usecase/      → Casos de uso (Feed, Play, Sleep)
├── repository/   → Interface + implementação do repositório
├── handler/      → HTTP Handlers
├── routes.go     → Registro das rotas
cmd/
└── main.go       → Setup inicial, CORS, database, server
```

- 🚀 **Separação total de camadas**: handler, domínio, persistência, orquestração.
- 🧠 **Eventos de domínio** (`MonsterDiedEvent`) com publisher interno.
- 🧪 **Testes** em todas as camadas: entidade, usecase, events.

---

## 🐲 Entidade: Monster

```go
type Monster struct {
  ID        uuid.UUID
  Name      string
  Hunger    int // 0 - 100
  Energy    int // 0 - 100
  Health    int // 0 - 100
  Happiness int // 0 - 100
  IsAlive   bool
  ...
}
```

### Métodos

- `Feed()`
- `Play()`
- `Sleep()`
- `Die()`
- Eventos acumulados via `AddEvent()` + `PullEvents()`

---

## ⚙️ Endpoints REST

| Método | Rota                         | Descrição               |
|--------|------------------------------|-------------------------|
| POST   | `/monsters`                  | Cria um novo monstro   |
| GET    | `/monsters/:id`              | Retorna status do monstro |
| PATCH  | `/monsters/feed/:id`         | Alimenta o monstro     |
| PATCH  | `/monsters/play/:id`         | Brinca com o monstro   |
| PATCH  | `/monsters/sleep/:id`        | Monstro descansa       |

---

## 🧪 Testes

- 🧱 Entidades: Validação de regras (hunger, energy, happy, death)
- 🔁 Table-Driven Tests (idiomático em Go)
- 🔜 Em breve: Testes de Handlers via `httptest`

### Rodar os testes:

```bash
go test ./...
```

---

## 💾 Banco de dados

- Banco: **SQLite**
- ORM: [GORM](https://gorm.io/)
- AutoMigrate aplicado ao iniciar a aplicação

---

## 🛡️ Segurança

- CORS habilitado (em dev: aberto)
- Headers preparados para futura autenticação (JWT)

---

## 📦 Como rodar

```bash
go run ./cmd/main.go
```

Banco será gerado automaticamente (`./monster.db`).

---

## 🎯 Objetivo

> Aprender e praticar desenvolvimento real de APIs com Go, aplicando:
>
> - Clean Architecture
> - Domain-Driven Design
> - Domain Events
> - Testes idiomáticos
> - Boas práticas REST

---

## 💡 Autor

Feito com 🧠 e ☕ por **Victor**, para estudos e evolução profissional em Golang.

---

## 🔮 Futuras melhorias

- 🧑‍🤝‍🧑 Autenticação e relação Usuário-Monstro
- 🦾 Evolução de fases e evolução por idade
- 🕰️ WebSocket ou Server Sent Events para feedback real-time
- 🧠 EventStore e Worker separado para rotina de decay

---

> 💬 *"Digital Monster API — o mundo é seu código."*
