package db

import (
	"database/sql"
	"errors"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/olioapps/service-skeleton-go/olio/dao"
	"github.com/olioapps/service-skeleton-go/olio/util"
)

type MigrationManager struct {
	connectionManager *dao.ConnectionManager
	migrations        []Migration
}

type Migration func() error

func (self *MigrationManager) getRequiredSchemaVersion() int {
	return len(self.migrations)
}

func (self *MigrationManager) getCurrentVersion() (int, error) {
	db := self.connectionManager.GetDb()
	var version int
	row := db.Table("migrations").Select("version").Row()
	row.Scan(&version)
	return version, db.Error
}

func (self *MigrationManager) perequisites() error {
	db := self.connectionManager.GetDb()
	if db.Exec(`
		CREATE TABLE IF NOT EXISTS
		migrations (
			version int
		) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci
	`).Error != nil {
		return db.Error
	}

	rows, error := db.Table("migrations").Select("version").Rows()
	if error != nil {
		return error
	}
	if rows == nil || rows.Err() == sql.ErrNoRows || !rows.Next() {
		return db.Exec("INSERT INTO migrations values (?)", 0).Error
	}
	return nil
}

func (self *MigrationManager) incrementVersion(targetVersion int) error {
	db := self.connectionManager.GetDb()
	return db.Exec("UPDATE migrations SET version = ?", targetVersion).Error
}

func (self *MigrationManager) Migrate() error {
	err := self.perequisites()
	if err != nil {
		return err
	}

	currentVersion, err := self.getCurrentVersion()
	if err != nil {
		return err
	}

	requiredVersion := self.getRequiredSchemaVersion()
	log.Info("Required schema version is ", requiredVersion)
	log.Info("Database is at schema version ", currentVersion)

	if currentVersion == requiredVersion {
		log.Info("Database schema is up to date.")
		return nil
	}

	if currentVersion > requiredVersion {
		return errors.New("Current schema cannot be ahead of code!")
	}

	targetVersion := currentVersion + 1

	if targetVersion > len(self.migrations) {
		return errors.New("Invalid target version " + util.Int64ToString(int64(targetVersion)) + "; no migration exists.")
	}

	err = self.migrations[targetVersion-1]()

	if err != nil {
		return errors.New("Failed to migrate from " + strconv.Itoa(currentVersion) + " to " + strconv.Itoa(targetVersion) + ": " + err.Error())
	}

	if err := self.incrementVersion(targetVersion); err != nil {
		panic(err)
	}

	return self.Migrate()
}

func NewMigrationManager(connectionManager *dao.ConnectionManager, migrations []Migration) *MigrationManager {
	migrationManager := MigrationManager{}
	migrationManager.connectionManager = connectionManager
	migrationManager.migrations = migrations

	return &migrationManager
}
