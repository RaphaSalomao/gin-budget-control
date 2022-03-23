package database

import (
	"fmt"

	gormPostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

var (
	DB *gorm.DB
	M  *migrate.Migrate

	DbHost     string
	DbUser     string
	DbName     string
	DbPort     string
	DbSslMode  string
	DbPassword string

	MigrationSourcePath string
)

func Connect() error {
	var err error
	config := gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	DB, err = gorm.Open(gormPostgres.Open(DbConnectionString()), &config)
	if err != nil {
		return err
	}
	setupMigration()
	doMigrate()
	return nil
}

func DbConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		DbUser, DbPassword, DbHost, DbPort, DbName, DbSslMode)
}

func setupMigration() {
	db, err := DB.DB()
	if err != nil {
		panic(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		panic(err)
	}
	M, err = migrate.NewWithDatabaseInstance(MigrationSourcePath, "postgres", driver)
	if err != nil {
		panic(err)
	}
}

func doMigrate() {
	err := M.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Println("Migration Up error:", err)
		handleMigrationError(M, err)
	}
}

func handleMigrationError(m *migrate.Migrate, err error) {
	version, dirty, migrateError := m.Version()
	if migrateError != nil {
		panic(migrateError)
	}
	migrateError = migrate.ErrDirty{Version: int(version)}
	if err == migrateError && dirty {
		m.Force(int(version - 1))
		doMigrate()
	} else {
		panic(err)
	}
}
