package modulehost

import (
	"context"

	ormdialect "github.com/domainry/domainry-orm/dialect"
	ormmigration "github.com/domainry/domainry-orm/migration"
	"github.com/domainry/domainry-orm/query"
	"github.com/domainry/domainry-orm/sqlhost"
)

type Executor = sqlhost.Executor
type Queryer = sqlhost.Queryer
type DBTX = sqlhost.DBTX
type Database = sqlhost.Database

type Dialect interface {
	query.Renderer
	Name() ormdialect.Name
}

type SchemaMigration = ormmigration.Migration
type SchemaBaseline = ormmigration.Baseline
type SchemaTable = ormmigration.Table
type SchemaColumn = ormmigration.Column
type SchemaIndex = ormmigration.Index

type MigrationRegistrar interface {
	ApplyOwnedMigrations(context.Context, string, []SchemaMigration) error
}

type Transactor interface {
	WithinTransaction(context.Context, func(context.Context, DBTX) error) error
}

type Host interface {
	Database() Database
	Dialect() Dialect
	Migrations() MigrationRegistrar
	Transactions() Transactor
}
