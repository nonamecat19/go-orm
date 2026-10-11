# go-orm-test-app

A full-stack **shopping list** application built as the demo and integration testbed for [**go-orm**](../README.md) — a reflection-based, adapter-driven ORM and query builder for Go, with an embeddable web admin UI ([Studio](../README.md#studio-web-admin-ui)).

Stack: Go + [Fiber](https://gofiber.io/) API · React + Vite + Tailwind frontend · PostgreSQL · nginx · Docker Compose.

---

## Table of contents

- [What this demonstrates](#what-this-demonstrates)
- [Layout](#layout)
- [Running it](#running-it)
- [Configuration](#configuration)
- [API](#api)
- [Studio](#studio)
- [Related repositories](#related-repositories)

## What this demonstrates

The backend consumes go-orm as a published dependency (see `backend/go.mod`), not as a workspace module — so it exercises the library exactly as an external consumer would.

| go-orm feature | Documentation | Where it is used here |
|---|---|---|
| Client construction from `ORMConfig` + `AdapterPostgres` | [Quick start](../README.md#quick-start) | `backend/database/database.go` |
| Entities as plain structs with `db:` tags and an `Info()` table name | [Defining entities](../README.md#defining-entities) | `backend/entities/` |
| One-to-many relations via `relation:` tags and preloading | [Relations and preloading](../README.md#relations-and-preloading) | `backend/entities/list.go`, `backend/services/` |
| Fluent query builder — read, insert, update, delete | [Query builder API](../README.md#query-builder-api) | `backend/services/` |
| Mounting the admin UI into a Fiber app | [Studio](../README.md#studio-web-admin-ui) | `backend/main.go` |

The two entities are `List` (`lists`) and `Item` (`items`), joined by `items.list_id`:

```go
type List struct {
	Model
	Name  string `db:"name" json:"name"`
	Items []Item `db:"items" relation:"list_id" json:"items,omitempty"`
}

type Item struct {
	Model
	Name   string `db:"name" json:"name"`
	Bought bool   `db:"bought" json:"bought"`
	ListId *int64 `db:"list_id" json:"listId,omitempty"`
	List   *List  `db:"list" relation:"list_id" json:"list,omitempty"`
}
```

## Layout

```
backend/           Go + Fiber REST API
  database/        go-orm client setup and goose migration runner
  entities/        go-orm entity structs (List, Item)
  services/        Query builder calls
  handlers/        HTTP handlers
  migrations/      SQL migrations (goose)
frontend/          React + Vite + Tailwind + shadcn/ui SPA
nginx.conf         Serves the built frontend, proxies /api to the backend
docker-compose.yml postgres + backend + frontend build + nginx
```

## Running it

Requires Docker and Docker Compose.

Create a `.env` file at the repository root (it is gitignored — see [Configuration](#configuration) for the variables), then:

```bash
docker compose up --build
```

Then open:

- **http://localhost** — the shopping list UI (nginx)
- **http://localhost:8080/api/health** — API health check
- **http://localhost/api/admin** — go-orm Studio

Migrations run automatically on backend startup. Postgres is exposed on host port `25432`.

### Without Docker

```bash
# backend
cd backend && go run .

# frontend
cd frontend && npm install && npm run dev
```

The backend expects a reachable Postgres and the `DB_*` variables below.

## Configuration

`.env` at the repository root:

| Variable | Used by | Purpose |
|---|---|---|
| `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB` | postgres | Database bootstrap |
| `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` | backend | go-orm `ORMConfig` |
| `VITE_API_URL` | frontend build | API base URL baked into the bundle |

Inside Compose, `DB_HOST` is `postgres` and `DB_PORT` is `5432`.

## API

All routes are under `/api`.

| Method | Route | Description |
|---|---|---|
| `GET` | `/health` | Health check |
| `GET` | `/lists` | List all shopping lists |
| `POST` | `/lists` | Create a list |
| `GET` | `/lists/:id` | Get a list with its items |
| `DELETE` | `/lists/:id` | Delete a list |
| `POST` | `/lists/:listId/items` | Attach an item to a list |
| `DELETE` | `/lists/:listId/items/:itemId` | Detach an item from a list |
| `GET` | `/items` | List all items |
| `POST` | `/items` | Create an item |
| `PATCH` | `/items/:id` | Update an item |
| `DELETE` | `/items/:id` | Delete an item |

## Studio

`backend/main.go` mounts go-orm's admin UI at `/api/admin`:

```go
tables := []coreEntities.Entity{
	entities.Item{},
	entities.List{},
}
studio.AddStudioFiberGroup(app, tables, database.DbClient, "/api/admin")
```

> **Warning:** Studio ships with no authentication and interpolates its `sort`/`dir` query parameters directly into `ORDER BY`. It is mounted unprotected here because this is a local demo — never expose it publicly. See [Known issues and limitations](../README.md#known-issues-and-limitations) in the library README.

## Related project

- [**go-orm**](../README.md) — the ORM library this app is built on.
