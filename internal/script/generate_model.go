package main

import (
	"context"
	"log"
	"unicode"

	"github.com/bajra-manandhar17/personal-finance-app/internal/config"
	"gorm.io/gen"
)

func main() {
	ctx := context.Background()

	// Connect to PostgreSQL database using NewPostgresDbProvider
	db, err := config.NewPostgresDbProvider(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
	}

	// Get all public tables from PostgreSQL database
	var tables []string
	rows, err := db.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema = 'public'").Rows()
	if err != nil {
		log.Fatalf("Failed to fetch table names: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Fatalf("Error scanning table name: %v", err)
		}
		tables = append(tables, tableName)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error iterating over table names: %v", err)
	}

	// Generate models and queries for each table
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/db/query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)

	for _, table := range tables {
		modelName := capitalize(table)
		g.ApplyBasic(g.GenerateModelAs(table, modelName))
	}

	g.Execute()
}

// Helper function to capitalize the first letter of a string
func capitalize(str string) string {
	if str == "" {
		return ""
	}
	r := []rune(str)
	return string(unicode.ToUpper(r[0])) + string(r[1:])
}
