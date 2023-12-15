//go:build mage
// +build mage

package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/Bulut-Bilisimciler/go-ms-boilerplate/pkg/infrastructure"
	"github.com/magefile/mage/mg"
)

type Database mg.Namespace

// SyncDB syncs a table from remote db to local db Usage "mage database:syncDB LOCAL_DB_DSN REMOTE_DB_DSN TO_LOCAL_SCHEMA_NAME TABLES_WITH_COMMA_STRING"
func (Database) SyncDB(localDbDsn string, remoteDbDsn string, toLocalSchema string, syncTables string) error {
	ReadMageConfig()

	localParsed, err := url.Parse(localDbDsn)
	if err != nil {
		fmt.Println("Error parsing local db dsn")
		panic(err.Error())
	}

	// parse remote db dsn
	localUsername := localParsed.User.Username()
	// localPassword, _ := localParsed.User.Password()
	localHost := localParsed.Hostname()
	localPort := localParsed.Port()
	localDbName := strings.TrimPrefix(localParsed.Path, "/")
	localSchema := localParsed.Query().Get("search_path")
	if localSchema == "" {
		localSchema = "public"
	}

	fmt.Printf("LOCAL: u/h/p/d/s %s/%s/%s/%s/%s\n", localUsername, localHost, localPort, localDbName, localSchema)

	// parse remote db dsn
	remoteParsed, err := url.Parse(remoteDbDsn)
	if err != nil {
		fmt.Println("Error parsing remote db dsn")
		panic(err.Error())
	}

	// parse remote db dsn
	remoteUsername := remoteParsed.User.Username()
	remotePassword, _ := remoteParsed.User.Password()
	remoteHost := remoteParsed.Hostname()
	remotePort := remoteParsed.Port()
	remoteDbName := strings.TrimPrefix(remoteParsed.Path, "/")
	remoteSchema := remoteParsed.Query().Get("search_path")
	if remoteSchema == "" {
		remoteSchema = "public"
	}

	fmt.Printf("REMOTE: u/h/p/d/s %s/%s/%s/%s/%s\n", remoteUsername, remoteHost, remotePort, remoteDbName, remoteSchema)

	// *gorm.DB
	dbLocal := infrastructure.NewPostgresDB(localDbDsn)

	// split by comma
	tables := strings.Split(syncTables, ",")

	// Create a temporary schema for Foreign Data Wrapper
	fdwSchema := toLocalSchema + "_fdw"

	fmt.Printf("LOCAL_SCHEMA: %s\n", toLocalSchema)
	fmt.Printf("FDW_SCHEMA: %s\n", fdwSchema)
	fmt.Printf("TABLES: %s\n", syncTables)

	// Step -1: Clean up before starting - Drop temporary schema and user and server information
	dbLocal.Exec("DROP SCHEMA IF EXISTS " + toLocalSchema + " CASCADE;")
	dbLocal.Exec("DROP SCHEMA IF EXISTS " + fdwSchema + " CASCADE;")
	dbLocal.Exec("DROP USER MAPPING IF EXISTS FOR " + localUsername + " SERVER remote_server;")
	dbLocal.Exec("DROP SERVER IF EXISTS remote_server;")
	dbLocal.Exec("DROP EXTENSION IF EXISTS postgres_fdw;")
	// drop user mapping

	// create extension first if err panic
	// Fixed error: no schema has been selected to create in

	// Step 1: Create Foreign Data Wrapper schema and user
	dbLocal.Exec("CREATE SCHEMA IF NOT EXISTS " + localSchema + ";")
	dbLocal.Exec("CREATE SCHEMA IF NOT EXISTS " + toLocalSchema + ";")
	dbLocal.Exec("CREATE SCHEMA IF NOT EXISTS " + fdwSchema + ";")

	// give permission for schemas
	dbLocal.Exec("GRANT USAGE ON SCHEMA " + localSchema + " TO " + localUsername + ";")
	dbLocal.Exec("GRANT USAGE ON SCHEMA " + toLocalSchema + " TO " + localUsername + ";")
	dbLocal.Exec("GRANT USAGE ON SCHEMA " + fdwSchema + " TO " + localUsername + ";")

	// select schema first
	dbLocal.Exec("SET search_path TO " + localSchema + ";")
	dbLocal.Exec("CREATE EXTENSION IF NOT EXISTS postgres_fdw;")

	// Step 2: Create a foreign server
	// execute_sql "CREATE SERVER foreigndb_fdw FOREIGN DATA WRAPPER postgres_fdw OPTIONS (dbname '$REMOTE_DB_DSN');"
	dbLocal.Exec("CREATE SERVER IF NOT EXISTS remote_server FOREIGN DATA WRAPPER postgres_fdw OPTIONS (host '" + remoteHost + "', port '" + remotePort + "', dbname '" + remoteDbName + "');")

	// Step 3: Create a user mapping
	// execute_sql "CREATE USER MAPPING FOR fdw_user SERVER foreigndb_fdw OPTIONS (user '$username', password '$password');"
	dbLocal.Exec("CREATE USER MAPPING FOR " + localUsername + " SERVER remote_server OPTIONS (user '" + remoteUsername + "', password '" + remotePassword + "');")
	dbLocal.Exec(
		"GRANT USAGE ON FOREIGN SERVER remote_server TO " + localUsername + ";" +
			"GRANT USAGE ON SCHEMA " + remoteSchema + " TO " + localUsername + ";" +
			"GRANT SELECT ON ALL TABLES IN SCHEMA " + remoteSchema + " TO " + localUsername + ";",
	)

	// Step 4: Import remote tables into the local schema
	dbLocal.Exec("IMPORT FOREIGN SCHEMA " + remoteSchema + " FROM SERVER remote_server INTO " + fdwSchema + ";")

	for _, table := range tables {
		// fetch remote schema first (dont forget serials and sequences)
		dbLocal.Exec("CREATE TABLE IF NOT EXISTS " + toLocalSchema + "." + table + " (LIKE " + fdwSchema + "." + table + " INCLUDING ALL);")

		// insert into local table from remote table
		dbLocal.Exec("INSERT INTO " + toLocalSchema + "." + table + " SELECT * FROM " + fdwSchema + "." + table + ";")

	}
	// Step 5: Clean up - Drop temporary schema and user and server information
	dbLocal.Exec("DROP SCHEMA IF EXISTS " + fdwSchema + " CASCADE;")
	dbLocal.Exec("DROP SERVER IF EXISTS remote_server CASCADE;")
	dbLocal.Exec("DROP EXTENSION IF EXISTS postgres_fdw CASCADE;")

	fmt.Println("Sync completed successfully.")
	return nil
}
