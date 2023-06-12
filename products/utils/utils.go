package utils

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
)

func ExecuteSQLFile(db *sql.DB, filename string) error {
	// Read the SQL file
	sqlBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Failed to read SQL file: %v", err)
	}
	queries := string(sqlBytes)
	// Split the SQL statements using semicolon as the delimiter
	statements := splitStatements(queries)

	for _, statement := range statements {
		_, err = db.Exec(statement)
		if err != nil {
			log.Printf("Error executing statement: %s\n%s", statement, err)
		}
	}

	return nil
}

// splitStatements Split the SQL statements using semicolon as the delimiter
func splitStatements(queries string) []string {
	statements := make([]string, 0)
	currentStatement := ""
	inString := false

	for _, ch := range queries {
		switch ch {
		case ';':
			if !inString {
				statements = append(statements, currentStatement)
				currentStatement = ""
			} else {
				currentStatement += string(ch)
			}
		case '\'':
			inString = !inString
			currentStatement += string(ch)
		default:
			currentStatement += string(ch)
		}
	}

	return statements
}
