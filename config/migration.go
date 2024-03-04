package config

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

const migrationDirectory = "file://./internal/infra/database/migrations"

func RunMigrations(configs *conf) {
	// Starting run migrations
	log.Println("Running migrations...")
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName)
	db, err := sql.Open("mysql", dbUrl)
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationDirectory,
		configs.DBName,
		driver,
	)

	defer db.Close()
	if err != nil {
		log.Panic(err)
	}

	// Run migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	} else if errors.Is(err, migrate.ErrNoChange) {
		log.Println("No migrations to run!")
	} else {
		log.Println("Migrations ran successfully!")
	}
}
