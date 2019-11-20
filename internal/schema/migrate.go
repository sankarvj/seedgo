package schema

import (
	"github.com/dimiro1/darwin"
	"github.com/jmoiron/sqlx"
)

// Migrate attempts to bring the schema for db up to date with the migrations
// defined in this package.
func Migrate(db *sqlx.DB) error {
	driver := darwin.NewGenericDriver(db.DB, darwin.PostgresDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}

// migrations contains the queries needed to construct the database schema.
// Entries should never be removed from this slice once they have been ran in
// production.
//
// Using constants in a .go file is an easy way to ensure the queries are part
// of the compiled executable and avoids pathing issues with the working
// directory. It has the downside that it lacks syntax highlighting and may be
// harder to read for some cases compared to using .sql files. You may also
// consider a combined approach using a tool like packr or go-bindata.
var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Add accounts",
		Script: `
		CREATE TABLE accounts (
			account_id    UUID,
			name          TEXT,
			domain        TEXT UNIQUE,
			avatar        TEXT,
			plan          INTEGER DEFAULT 0,
			mode          INTEGER DEFAULT 0,
			timezone      TEXT,
			language      TEXT,
			country       TEXT,
			issued_at     TIMESTAMP,
			expiry        TIMESTAMP,
			created_at    TIMESTAMP,
			updated_at    BIGINT,
			PRIMARY KEY (account_id)
		);
		`,
	},
	{
		Version:     2,
		Description: "Add users",
		Script: `
		CREATE TABLE users (
			user_id       UUID,
			account_id    UUID REFERENCES accounts ON DELETE CASCADE,
			name          TEXT,
			avatar 		  TEXT,
			email         TEXT,
			phone         TEXT,
			verified      BOOLEAN DEFAULT FALSE,
			roles         TEXT[],
			password_hash TEXT,
			provider      TEXT,
			issued_at     TIMESTAMP,
			created_at    TIMESTAMP,
			updated_at    BIGINT,
			PRIMARY KEY (user_id),
			UNIQUE (account_id, email)
		);
		`,
	},
}
