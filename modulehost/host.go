package modulehost

import (
	"context"

	auditcontract "github.com/domainry/domainry-audit-sdk/contract"
	sharedartifact "github.com/domainry/domainry-foundation/artifact"
	lifecyclecontract "github.com/domainry/domainry-lifecycle-sdk/contract"
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
	Insert(string, []string) string
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

// AuditStoreHost supplies the installation-wide immutable Audit append ports.
// Lifecycle emits compliance facts through these ports instead of owning a
// second audit/evidence table.
type AuditStoreHost interface {
	AuditAppender() auditcontract.Appender
	AuditTransactionalAppender() auditcontract.TransactionalAppender
}

// ArtifactStoreHost supplies the shared governed metadata store and the
// deployment-owned byte boundary used by Lifecycle archives and uploads.
type ArtifactStoreHost interface {
	ArtifactStore() sharedartifact.ManagedStore
	ArtifactContentStore() lifecyclecontract.ArtifactContentStore
	ArtifactContentWriter() lifecyclecontract.ArtifactContentWriter
}

type Host interface {
	Database() Database
	Dialect() Dialect
	Migrations() MigrationRegistrar
	Transactions() Transactor
}
