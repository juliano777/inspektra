package db

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

// PgParams holds the configuration for the database connection.
type PgParams struct {
	Host            string
	Port            int
        Database        string
	User            string
	Password        string
	SSLMode         string
	ApplicationName string
}

// PgConnStrGen creates a connection string with default values.
func PgConnStrGen(p PgParams) string {
	if p.Port == 0 {
		p.Port = 5432
	}
	if p.SSLMode == "" {
		p.SSLMode = "disable"
	}
	if p.ApplicationName == "" {
		p.ApplicationName = "inspektra"
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s sslmode=%s fallback_application_name=%s",
		p.Host, p.Port, p.User, p.SSLMode, p.ApplicationName,
	)

	// Appends password only if it is provided (not empty).
	if p.Password != "" {
		connStr += fmt.Sprintf(" password=%s", p.Password)
	}

	return connStr
}

// PgConnect opens the database connection and returns the DB object.
func PgConnect(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		// Validating the non-anonymous import of lib/pq
		if pqErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf("postgres error: %v", pqErr)
		}
		return nil, err
	}

	return db, nil
}

// PgExec runs a query and returns a dynamic list of maps (rows and cols).
func PgExec(db *sql.DB, query string) ([]map[string]any, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var result []map[string]any

	for rows.Next() {
		// Create dynamic arrays to hold unknown column types
		vals := make([]any, len(cols))
		ptrs := make([]any, len(cols))
		for i := range vals {
			ptrs[i] = &vals[i]
		}

		if err := rows.Scan(ptrs...); err != nil {
			return nil, err
		}

		// Map each column name to its scanned value
		row := make(map[string]any)
		for i, col := range cols {
			val := vals[i]
			// Postgres drivers often return bytes, so we convert to string
			if b, ok := val.([]byte); ok {
				row[col] = string(b)
			} else {
				row[col] = val
			}
		}
		result = append(result, row)
	}

	return result, nil
}


