# go-orm

A reflection-based, adapter-driven ORM and query builder for Go, plus **Studio** — an embeddable web admin UI for browsing and editing your tables.

Repository: `github.com/nonamecat19/go-orm` · License: MIT

Looking for a working example? See the [demo application](go-orm-test-app/README.md) — a full-stack shopping list app built on this library.

Supported databases: **PostgreSQL**, **MySQL**, **SQLite**, **Microsoft SQL Server**.

---

## Table of contents

- [Overview](#overview)
- [Repository layout](#repository-layout)
- [Installation](#installation)
- [Quick start](#quick-start)
- [Defining entities](#defining-entities)
- [Query builder API](#query-builder-api)
  - [Read](#read)
  - [Insert](#insert)
  - [Update](#update)
  - [Delete](#delete)
  - [Raw SQL](#raw-sql)
  - [Debugging](#debugging)
- [Relations and preloading](#relations-and-preloading)
- [Architecture](#architecture)
  - [Module graph](#module-graph)
  - [The `DbAdapter` interface](#the-dbadapter-interface)
  - [Query generation pipeline](#query-generation-pipeline)
  - [Placeholder normalization](#placeholder-normalization)
  - [Adapter differences](#adapter-differences)
- [Studio (web admin UI)](#studio-web-admin-ui)
- [Demo application](#demo-application)
- [Development](#development)
  - [Local databases](#local-databases)
  - [Running tests](#running-tests)
  - [CI](#ci)
- [Known issues and limitations](#known-issues-and-limitations)

---

## Overview

go-orm is built around three ideas:

1. **Entities are plain structs.** Table and column mapping comes from an embedded `entities.Model` struct plus `db:"..."` struct tags. Table name comes from an `Info() string` method.
2. **A fluent query builder.** `CreateQueryBuilder(client).Where(...).OrderBy(...).Limit(...).FindMany(&out)` — chainable methods accumulate state, terminal methods (`FindMany`, `InsertOne`, `UpdateMany`, `DeleteMany`, …) build and execute SQL.
3. **Dialects live behind an adapter interface.** Every SQL fragment (`WHERE`, `LIMIT`, `SET`, `INSERT`, placeholder syntax) is produced by a `DbAdapter`. A shared `adapter-base` package holds the ANSI-ish default implementation; each dialect adapter delegates to it and overrides only what differs.

Result binding is done with `reflect` and `database/sql`'s `Rows.Scan`, writing directly into pointers to your struct fields.

## Repository layout

This is a **Go workspace** (`go.work`) containing eight modules:

| Path | Module | Purpose |
|---|---|---|
| `core/` | `.../core` | Shared contracts: `DbAdapter` interface, `ORMConfig`, `entities.Model`/`Entity`, `JoinClause`, and reflection/string utilities. No third-party DB deps. |
| `adapter-base/` | `.../adapter-base` | Default SQL fragment generators shared by all dialects (free functions, not methods). |
| `adapter-postgres/` | `.../adapter-postgres` | `AdapterPostgres` — `lib/pq`, `$N` placeholders. |
| `adapter-mysql/` | `.../adapter-mysql` | `AdapterMySQL` — `go-sql-driver/mysql`, `?` placeholders. |
| `adapter-sqlite/` | `.../adapter-sqlite` | `AdapterSQLite` — `mattn/go-sqlite3`, `$N` placeholders. |
| `adapter-mssql/` | `.../adapter-mssql` | `AdapterMSSQL` — `denisenkom/go-mssqldb`, `@pN` named params, `OFFSET/FETCH NEXT`. |
| `orm/` | `.../orm` | The user-facing API: `client` (connection) and `querybuilder` (fluent builder). Also hosts integration tests. |
| `studio/` | `.../studio` | Fiber + templ + htmx + Tailwind/daisyUI admin UI, mountable into any Fiber app. |

Adapter modules depend on `core` and `adapter-base`; `orm` depends on `core`; `studio` depends on `core`, `orm`, and `adapter-postgres`.

```
                 ┌──────────┐
                 │   core   │  DbAdapter, ORMConfig, Model, utils
                 └────┬─────┘
          ┌───────────┼────────────┐
          │           │            │
   ┌──────▼─────┐ ┌───▼───┐  ┌─────▼──────┐
   │adapter-base│ │  orm  │  │   studio   │
   └──────┬─────┘ └───▲───┘  └─────┬──────┘
          │           │            │
   ┌──────▼───────────┴──┐         │
   │ adapter-{pg,mysql,  ├─────────┘
   │   sqlite,mssql}     │
   └─────────────────────┘
```

## Installation

Each module is independently importable. Install the ORM plus the adapter(s) you need:

```bash
go get github.com/nonamecat19/go-orm/orm
go get github.com/nonamecat19/go-orm/core
go get github.com/nonamecat19/go-orm/adapter-postgres
```

You must also import the matching driver for its `database/sql` side effect:

| Adapter | Driver import | Driver name |
|---|---|---|
| Postgres | `_ "github.com/lib/pq"` | `postgres` |
| MySQL | `_ "github.com/go-sql-driver/mysql"` | `mysql` |
| SQLite | `_ "github.com/mattn/go-sqlite3"` | `sqlite3` |
| MSSQL | `_ "github.com/denisenkom/go-mssqldb"` | `sqlserver` |

## Quick start

```go
package main

import (
	"fmt"

	_ "github.com/lib/pq"
	adapterpostgres "github.com/nonamecat19/go-orm/adapter-postgres/lib"
	"github.com/nonamecat19/go-orm/core/lib/config"
	"github.com/nonamecat19/go-orm/orm/lib/client"
	"github.com/nonamecat19/go-orm/orm/lib/querybuilder"
)

func main() {
	db := client.CreateClient(config.ORMConfig{
		Host:     "127.0.0.1",
		Port:     15432,
		User:     "postgres",
		Password: "root",
		DbName:   "orm",
		SSLMode:  false,
	}, adapterpostgres.AdapterPostgres{})

	var users []User
	err := querybuilder.CreateQueryBuilder(db).
		Where("id > ?", 12).
		AndWhere("id < ?", 18).
		Preload("orders").
		Preload("role").
		OrderBy("id DESC").
		Limit(8).
		Offset(2).
		FindMany(&users)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(users))
}
```

### `ORMConfig`

```go
type ORMConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SSLMode  bool
	Path     string // SQLite only: file path, used verbatim as the DSN
}
```

Connection strings per adapter:

- Postgres: `host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`
- MySQL: `%s:%s@tcp(%s:%d)/%s?parseTime=true`
- MSSQL: `sqlserver://%s:%s@%s:%d?database=%s`
- SQLite: `config.Path`

> Note: `SSLMode` is currently present in the struct but not honoured — the Postgres DSN hardcodes `sslmode=disable`.

### `DbClient`

`client.CreateClient(cfg, adapter) DbClient` opens the pool via `sql.Open` (it calls `log.Fatal` on error rather than returning it). Accessors: `GetDb() *sql.DB`, `GetConfig()`, `GetAdapter()`.

## Defining entities

An entity is a struct that:

1. embeds `entities.Model` as its **first field**,
2. tags each column with `db:"column_name"`,
3. implements `Info() string` returning the table name.

```go
import "github.com/nonamecat19/go-orm/core/lib/entities"

type Role struct {
	entities.Model
	Name  string `db:"name"  json:"name,omitempty"`
	Users []User `db:"users" json:"users,omitempty" relation:"role_id"`
}

func (r Role) Info() string { return "roles" }

type User struct {
	entities.Model
	Name   string  `db:"name"    json:"name,omitempty"`
	Email  string  `db:"email"   json:"email,omitempty"`
	Gender string  `db:"gender"  json:"gender,omitempty"`
	Orders []Order `db:"orders"  json:"orders,omitempty" relation:"user_id"`
	RoleId *int64  `db:"role_id" json:"roleId,omitempty"`
	Role   *Role   `db:"role"    relation:"role_id" json:"role,omitempty"`
}

func (u User) Info() string { return "users" }

type Order struct {
	entities.Model
	Count     int64     `db:"count"      json:"count,omitempty"`
	UserId    *int64    `db:"user_id"    json:"userId,omitempty"`
	User      *User     `db:"user"       relation:"user_id" json:"user,omitempty"`
	OrderDate time.Time `db:"order_date" json:"orderDate,omitempty"`
}

func (o Order) Info() string { return "orders" }
```

`entities.Model` supplies the system columns present on every table:

```go
type Model struct {
	ID        int64      `db:"id"         json:"id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
```

### Tag semantics

| Tag | Meaning |
|---|---|
| `db:"col"` **without** `relation` | A real column. Included in SELECT/INSERT/UPDATE field lists. |
| `db:"name"` **with** `relation:"fk"` | A relation field. **Excluded** from column lists; loaded only via `Preload("name")`. |

`relation` names the foreign-key **column**:

- **Many-to-one** (field is `*T`): `relation` is the FK column on *this* entity (`User.Role` → `relation:"role_id"`, matching `User.RoleId`).
- **One-to-many** (field is `[]T`): `relation` is the FK column on the *related* table (`User.Orders` → `relation:"user_id"`, a column of `orders`).

## Query builder API

`querybuilder.CreateQueryBuilder(client) *QueryBuilder` starts a builder with `limit: -1, offset: -1` (meaning "unset"). Chainable methods return `*QueryBuilder`; terminal methods return `error` (or `(*sql.Rows, error)`).

| Chainable | Effect |
|---|---|
| `Select(cols ...string)` | Restrict selected columns. Names are **fully qualified** (`"users.id"`). |
| `Where(cond string, args ...any)` | Sets the WHERE clause to `(cond)`. |
| `AndWhere(cond string, args ...any)` | Appends ` AND (cond)`. |
| `OrWhere(cond string, args ...any)` | Appends ` OR (cond)`. |
| `OrderBy(fields ...string)` | Appends to `ORDER BY` (pass e.g. `"id DESC"`). |
| `Limit(n int)` | Row limit. |
| `Offset(n int)` | Row offset. |
| `LeftJoinAndSelect(table, condition string, selectFields ...string)` | Adds a `LEFT JOIN`. |
| `Preload(cols ...string)` | Eager-load relation fields by their `db` tag. |
| `SetValues(map[string]any)` | Column→value map for `UpdateMany`. |
| `Debug()` | Print the final SQL and args. |

| Terminal | Signature |
|---|---|
| `FindMany` | `FindMany(out any) error` — `out` must be `*[]Entity` |
| `FindOne` | *currently a no-op stub* |
| `InsertOne` | `InsertOne(entity any) error` — value, not pointer |
| `InsertMany` | `InsertMany(entities any) error` — a slice |
| `InsertMap` | `InsertMap(entity any, fields map[string]any) error` |
| `UpdateMany` | `UpdateMany(entity any) error` — pointer; requires `Where` + `SetValues` |
| `UpdateMap` | `UpdateMap(entity any, fields map[string]any) error` — requires `Where` |
| `DeleteMany` | `DeleteMany(entity any) error` — pointer; requires `Where` |
| `ExecuteRaw` | `ExecuteRaw(sql string, args ...any) (*sql.Rows, error)` |
| `ExecuteQuery` | `ExecuteQuery(query string, args ...any) (*sql.Rows, error)` |

### Read

```go
var users []User

err := querybuilder.CreateQueryBuilder(db).
	Select("users.id", "users.name", "users.gender", "users.email").
	Where("email = ?", "test@gmail.com").
	OrderBy("id ASC").
	Offset(0).
	Limit(1).
	FindMany(&users)
```

`FindMany` requires a pointer to a slice. It:

1. reflects the element type to get the table name and column list,
2. builds the query (see [pipeline](#query-generation-pipeline)),
3. scans each row into a fresh struct via field-address pointers,
4. runs any `Preload`s.

`Select` intersects with the full column list, so unknown/unqualified names silently drop columns — always use `table.column` form.

### Insert

```go
// Single struct (note: value, not pointer)
err := querybuilder.CreateQueryBuilder(db).
	InsertOne(User{Name: "testName", Email: "test@gmail.com", Gender: "non_binary", RoleId: &roleId})

// Bulk — one multi-row INSERT statement
err = querybuilder.CreateQueryBuilder(db).InsertMany([]User{u1, u2, u3})

// Partial insert from a map; keys are validated against the entity's db tags
err = querybuilder.CreateQueryBuilder(db).
	InsertMap(User{}, map[string]any{"name": "Ann", "email": "ann@example.com"})
```

Generated form: `INSERT INTO users (name, email, ...) VALUES (?, ?, ...), (?, ?, ...);` with placeholders normalized per dialect. `*T` fields serialize as their pointee or `nil`.

Note that only non-relation, non-system columns are written — `id`/`created_at`/`updated_at`/`deleted_at` are left to the database.

### Update

Both update paths **refuse to run without a `Where`**, guarding against a full-table update.

```go
err := querybuilder.CreateQueryBuilder(db).
	Where("role_id = 3").
	SetValues(map[string]any{"name": "testName"}).
	UpdateMany(&User{})

// or, with map-field validation against the entity:
err = querybuilder.CreateQueryBuilder(db).
	Where("id = ?", 42).
	UpdateMap(User{}, map[string]any{"name": "testName"})
```

### Delete

Also requires a `Where`.

```go
err := querybuilder.CreateQueryBuilder(db).
	Where("role_id = ?", 1).
	OrWhere("gender <> 'female'").
	DeleteMany(&User{})
```

### Raw SQL

```go
rows, err := querybuilder.CreateQueryBuilder(db).
	ExecuteRaw("SELECT id, name FROM users WHERE id > ?", 10)
```

`ExecuteRaw` runs the SQL through the adapter's placeholder normalizer first, so `?` works across dialects. Scanning is up to you.

### Debugging

`.Debug()` prints the fully prepared statement (green, lightly pretty-printed with newlines before `SELECT`/`FROM`/`JOIN`/`WHERE`/…) and the argument slice (yellow) to stdout before execution.

## Relations and preloading

`Preload` avoids N+1 by issuing **one extra query per relation**, then stitching results in memory.

```go
var users []User
err := querybuilder.CreateQueryBuilder(db).
	Preload("orders"). // one-to-many: []Order
	Preload("role").   // many-to-one: *Role
	FindMany(&users)
```

The argument is the relation field's `db` tag. Dispatch is by field kind:

- **`*T` (many-to-one)** — `preloadRelationPointer`: collects the distinct FK values from the parent rows (`relation` column on the parent), runs `SELECT * FROM <table> WHERE id IN (...)`, maps by `id`, and assigns pointers back.
- **`[]T` (one-to-many)** — `preloadRelationSlice`: collects distinct parent `id`s, runs `SELECT <cols> FROM <table> WHERE <relation> IN (...)`, groups rows by the FK value, and assigns each parent its slice.

Anything else returns `"<name> must be relation field"`.

Both paths currently assume the key type is `int64` or `*int64`; other key types return an error.

`LeftJoinAndSelect` is a separate, lower-level mechanism: it appends a real `LEFT JOIN` to the statement and reserves scan slots for the joined columns — but those values are scanned into throwaway `new(any)` placeholders, so joins currently only affect the SQL and row multiplicity, not the materialized structs. Use `Preload` for populating relations.

## Architecture

### Module graph

`core` defines the contract; `adapter-base` implements the default SQL; each dialect adapter is a thin struct whose methods forward to `adapter-base` except where the dialect differs; `orm` composes a client + builder against the `DbAdapter` interface only, so it never imports a dialect package.

### The `DbAdapter` interface

Defined in `core/lib/adapter/adapter.go` — every dialect implements all of it:

```go
type DbAdapter interface {
	// connection
	GetConnString(config config.ORMConfig) string
	GetDbDriver() string

	// statement heads
	DeleteFromTable(tableName string) string
	Update(tableName string) string
	Insert(tableName string, fieldNames []string, values []any, args []any) (string, []any)

	// clause builders
	PrepareOrderBy(query string, orderBy []string) string
	PrepareWhere(query string, where string) string
	PrepareLimit(query string, limit int) string
	PrepareOffset(query string, offset int) string
	PrepareSet(query string, set map[string]any, args []any) (string, []any)
	PrepareJoins(query string, joins []query.JoinClause) string

	// where composition
	Where(condition string, args []any) string
	AndWhere(condition string, where string, args []any) string
	OrWhere(condition string, where string, args []any) string

	// select construction
	GetReadQuery(tableName string, fields []string, fromSubquery string) string
	GetFromSubquery(tableName string, where string, orderBy []string, limit int, offset int) string
	GetSelectQuery(selectValue string, fromValue string) string
	GetSelectWhereIn(tableName string, selectValue string, fieldName string, fieldValues []string) string

	// formatting
	JoinFields(fields []string) string
	JoinFieldsStrictly(fields []string) string
	NormalizeSqlWithArgs(sql string, args []any) string
	PrepareQueryAndArgs(query string, args []any) (string, []any)
}
```

### Query generation pipeline

A `FindMany` builds a **subquery-wrapped** SELECT so that `WHERE`/`ORDER BY`/`LIMIT`/`OFFSET` apply to the base table before joins fan rows out:

1. `GetFromSubquery` → `SELECT * FROM users WHERE (...) ORDER BY id DESC LIMIT 8 OFFSET 2`
2. `GetReadQuery` wraps it and aliases every column strictly:
   `SELECT users.id AS "users.id", users.name AS "users.name", … FROM (<subquery>) AS users`
3. `PrepareJoins` appends any `LEFT JOIN … ON …`.

`JoinFieldsStrictly` is what produces the `col AS "col"` aliasing — it keeps qualified names stable through the subquery wrap and makes the column order deterministic for `Rows.Scan`. `JoinFields` is the plain `, `-join (also mapping the literal string `"<nil>"` to `NULL`).

Because scanning is **positional**, the field-pointer list is assembled in the same order as the SELECT list: system fields first (`id, created_at, updated_at, deleted_at`), then entity fields, then one throwaway pointer per joined select column.

### Placeholder normalization

The builder always writes `?` internally. `NormalizeSqlWithArgs(sql, args)` rewrites them, starting the counter at `len(args)+1` so that fragments appended later (a second `AndWhere`, a `SET` clause) continue the numbering:

| Adapter | Placeholder | Final arg handling |
|---|---|---|
| Postgres | `$1, $2, …` | args passed through unchanged |
| SQLite | `$1, $2, …` | args passed through unchanged |
| MySQL | `?` (no rewrite) | args passed through unchanged |
| MSSQL | `@p1, @p2, …` | `PrepareQueryAndArgs` wraps each arg in `sql.Named("pN", v)` |

### Adapter differences

Everything except the following forwards to `adapter-base`:

- **MySQL** — `NormalizeSqlWithArgs` is the identity function (MySQL uses `?` natively); `Insert`, `Where`/`AndWhere`/`OrWhere`, and `PrepareSet` are reimplemented as methods so they call the MySQL normalizer rather than the base one.
- **MSSQL** — `@pN` placeholders; `PrepareLimit` emits `FETCH NEXT n ROWS ONLY` and `PrepareOffset` emits `OFFSET n ROWS`; `GetFromSubquery` reorders clause application to `WHERE → ORDER BY → OFFSET → LIMIT` (T-SQL requires `OFFSET` before `FETCH`, and both require `ORDER BY`); args become `sql.Named` values.
- **Postgres / SQLite** — pure delegation to `adapter-base`; they differ only in `GetConnString` and `GetDbDriver`.

### Reflection utilities (`core/utils`)

| Function | Purpose |
|---|---|
| `ExtractTableAndFields(entity, prefix)` | Returns `(tableName, entityFields, systemFields, error)`; `prefix=true` qualifies names as `table.col`. Errors unless the entity is a pointer to a struct implementing `Info() string`. |
| `ExtractTableAndFieldsFromType(t, prefix)` | Same, from a `reflect.Type`. |
| `ExtractFields(t)` | All `db` tags **without** a `relation` tag. |
| `GetSystemFields()` | The four `entities.Model` columns. |
| `GetFieldNameByTagValue(t, tag)` | Struct field name for a given `db` tag value. |
| `GetModelFields(model)` | `map[dbTag]*field`. |
| `StringifyReflectValue(v)` | Human-readable rendering used by Studio (`time.Time` → `02/01/2006 15:04:05`, nil → `NULL`). |
| `ToSnakeCase` / `ToHumanCase` | Naming helpers. |
| `If`, `Map`, `MapWithIndex`, `Chunk`, `Contains`, `StringsIntersection`, `AddPrefix`, `GenerateParamsSlice` | Generic slice/string helpers. |

## Studio (web admin UI)

Studio is a database admin panel rendered server-side with [templ](https://templ.guide), driven by [htmx](https://htmx.org), styled with Tailwind + daisyUI, and served by [Fiber](https://gofiber.io). CSS and JS are **inlined into the base template** at build time, so the binary ships without external asset requests.

### Mounting it

```go
import (
	"github.com/gofiber/fiber/v2"
	coreEntities "github.com/nonamecat19/go-orm/core/lib/entities"
	"github.com/nonamecat19/go-orm/studio/lib/app"
)

server := fiber.New()

tables := []coreEntities.Entity{ Order{}, Role{}, User{} }

app.AddStudioFiberGroup(server, tables, dbClient, "/studio")

server.Listen(":8080")
```

`AddStudioFiberGroup(app, tables, client, prefix)` registers a Fiber group and stashes `{tableMap, tables, client, prefix}` in `c.Locals("data")` for handlers. A middleware rewrites emitted URLs (`href="/…"`, `hx-get="/tables/…"`, form actions, …) to be prefix-aware by string replacement on the response body.

### Routes

| Method | Path (relative to prefix) | Handler |
|---|---|---|
| GET | `/` | Table list |
| GET | `/tables/:id` | Table detail — sortable grid, add/edit/delete controls |
| GET | `/settings` | Settings (daisyUI theme picker) |
| POST | `/api/tables/:table_id/records` | Create record |
| PUT | `/api/tables/:table_id/records/:record_id` | Update record |
| DELETE | `/api/tables/:table_id/records/:record_id` | Delete record |
| GET | `/assets/*` | Static assets (packr box) |

Detail pages accept `?sort=<column>&dir=<asc\|desc>` and support htmx partial rendering: an `HX-Request: true` with `HX-Target: table-container` re-renders only the grid. Mutating endpoints respond with an `HX-Redirect` back to the table.

Column metadata (name, Go type, nullability inferred from a leading `*`) is derived reflectively via `buildFieldInfo`, and cell values are rendered through `StringifyReflectValue`.

### Studio source layout

```
studio/
├── cmd/main.go                    # standalone demo server (hardcoded Postgres config)
├── cmd/embed_assets.go            # asset-inlining CLI (also package main — see Known issues)
├── lib/app/app.go                 # AddStudioFiberGroup — the public entry point
├── internal/config/               # StudioConfig (ServerAddr, default ":8080")
├── internal/handlers/             # tablesGroup/*, settingsGroup/*
├── internal/model/                # FieldInfo view model
├── internal/utils/                # SharedData accessors, templ Render helper
├── internal/view/                 # .templ sources + generated *_templ.go
│   ├── layout/  components/  recordForm/  tables/  settings/  icons/
├── web/                           # Tailwind + esbuild sources
│   ├── src/main.css  src/main.js  tailwind.config.js  package.json
│   └── public/assets/             # built main.css / main.js
├── scripts/embed_assets*.sh       # inline built CSS/JS into base.templ / base_templ.go
└── Makefile
```

### Building Studio

```bash
cd studio
make install-deps          # npm install + go mod download
make build                 # css → js → embed → templ generate → go build -o ./bin/myapp
make watch                 # air + templ --watch + tailwind/esbuild watchers (parallel)
make fmt                   # templ fmt internal/view
```

Requires Node/npm and the `templ` CLI. `make watch` proxies through templ's dev proxy on top of the app at `localhost:8080`.

## Demo application

[**go-orm-test-app**](go-orm-test-app/README.md) is a full-stack shopping list application included in this repository. Its backend consumes this library as a published dependency. It is the reference for how the pieces fit together in a real project:

| Concern | Where to look |
|---|---|
| Client setup from env vars ([`ORMConfig`](#ormconfig), `AdapterPostgres`) | [`backend/database/database.go`](go-orm-test-app/backend/database/database.go) |
| [Entity definitions](#defining-entities) with `db:` tags and a one-to-many `relation:` | [`backend/entities/`](go-orm-test-app/backend/entities) |
| [Query builder](#query-builder-api) usage — reads, inserts, updates, deletes, [preloading](#relations-and-preloading) | [`backend/services/`](go-orm-test-app/backend/services) |
| [Mounting Studio](#mounting-it) into a Fiber app (`AddStudioFiberGroup`) | [`backend/main.go`](go-orm-test-app/backend/main.go) |

The app pairs a Go/Fiber REST API with a React + Vite frontend and Postgres, and runs end to end via Docker Compose.

## Development

### Local databases

`docker-compose.yml` brings up all three server databases used by the integration tests (SQLite is a local file):

```bash
docker compose up -d
```

| Service | Image | Host port | Credentials | Database |
|---|---|---|---|---|
| `postgres_db` | `postgres:15-alpine` | 15432 | `postgres` / `root` | `orm` |
| `mysql_db` | `mysql:8.0` | 3306 | `admin` / `root` | `orm` |
| `mssql_db` | `mssql/server:2022-latest` | 1433 | `sa` / `1StrongPwd!!` | `master` |

The compose file also defines an `mssql_configurator` one-shot container that runs `./init/init.sql` against SQL Server once it is healthy. That `init/` directory is not committed — create it (or drop the service) if you need MSSQL bootstrapping beyond what the test dumps do.

### Running tests

Integration tests live in `orm/tests/db/` — one file per dialect, each exercising the same five scenarios defined in `orm/tests/test-utils/queries.go`:

`FindUsersWithOrdersAndRole`, `FindOrdersWithUsers`, `DeleteUsers`, `InsertUser`, `UpdateUsers`.

Each test first calls `PrepareDB`, which reads `orm/tests/dumps/<driver>.sql` (`postgres.sql`, `mysql.sql`, `sqlite3.sql`, `sqlserver.sql` — named after `GetDbDriver()`), splits it on `;`, and executes every statement. The dumps drop and recreate `roles`/`users`/`orders` and seed 4 roles, 20 users, and a set of orders.

Results are serialized to JSON and compared against golden files in `orm/tests/outputs/` with `assert.JSONEq`.

```bash
cd orm
go test ./tests/db/...                    # all dialects (needs docker compose up)
go test ./tests/db/ -run TestPostgres      # one dialect
go test ./tests/db/ -run TestSQLite        # no containers needed
```

These are true integration tests against live databases — there is no unit-test suite for query generation.

### CI

Four GitHub Actions workflows (`.github/workflows/`), all on push/PR to `main`, all running `go work sync` first and then iterating the module list `adapter-base adapter-mysql adapter-mssql adapter-postgres adapter-sqlite core orm`:

| Workflow | What it runs |
|---|---|
| `build.yml` | `go build -v ./...` per module |
| `code-formatting.yml` | `gofmt -l` over every `.go` file, fails if any output |
| `linting.yml` | `golangci-lint run --timeout=5m ./...` per module |
| `static-analysis.yml` | `staticcheck ./...` per module |

Note the module list **excludes `studio`**, so Studio is not built or linted in CI.

## Known issues and limitations

These are current, verifiable characteristics of the codebase — worth knowing before you build on it.

**Build / structure**

- `studio/cmd/` fails to build: both `main.go` and `embed_assets.go` declare `package main` with a `func main`, giving `main redeclared in this block`. `embed_assets.go` is a build-time helper and needs its own directory or a build tag. The `studio/lib/app` package itself builds fine, so mounting Studio into your own app works; only the bundled demo binary is broken.
- `orm/tests/test-utils/queries.go` imports entities from the **published** `github.com/nonamecat19/go-orm/app/entities` module (pinned at `v0.0.0-20250517165903-…`), not from the local `orm/tests/entities/` package — which is therefore dead code, and the two can drift. `studio/cmd/main.go` imports the same remote package. Tests resolve it only because `studio`'s `go.mod` pulls it into the workspace module graph.

**API**

- `FindOne()` is an empty stub — it takes no arguments and does nothing. Use `FindMany` with `Limit(1)`.
- `DbClient.Query(query, args...)` passes `args` as a single slice argument to `db.Query` rather than expanding it, so it will not bind parameters correctly. `QueryBuilder.ExecuteQuery` is the working path.
- `LeftJoinAndSelect` scans joined columns into discarded `new(any)` values; joined data never reaches your structs.
- `config.SSLMode` is ignored — the Postgres DSN hardcodes `sslmode=disable`.
- `CreateClient` calls `log.Fatal` on `sql.Open` failure instead of returning the error.
- There is no transaction API, no context/cancellation support (`QueryContext` is never used), no migrations, no schema introspection, no soft-delete handling (`DeletedAt` is stored and read but never filtered on), and no connection-pool configuration.
- `Where` **replaces** rather than appends: calling `Where` twice discards the first condition. Use `AndWhere`/`OrWhere` for composition.
- Preload key handling assumes `int64` / `*int64` primary and foreign keys.
- `PrepareSet` iterates a `map[string]any` twice — once to build `col = ?` clauses and once to append values. Go map iteration order is randomized per range, so the SET clause order and the argument order can disagree, binding values to the wrong columns. Single-key updates (the tested path) are unaffected.

**Safety**

- `Where`/`OrderBy`/`Select` fragments and join conditions are interpolated into SQL as raw strings; only `?` arguments are bound as parameters. Never build those fragments from untrusted input.
- Studio's `getTableRecords` interpolates the `sort` and `dir` query parameters straight into `ORDER BY` with no allowlist — that is a SQL injection vector on any exposed Studio instance. Studio also ships no authentication or authorization of any kind; mount it behind your own auth middleware and never expose it publicly.
- Studio writes all form values as strings and relies on the database to coerce types.
- Studio handlers call `.Debug()` unconditionally, printing every generated statement and its arguments to stdout.
