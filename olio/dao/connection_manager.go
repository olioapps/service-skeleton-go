package dao

import (
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
	"github.com/olioapps/service-skeleton-go/olio/extractors"
	log "github.com/sirupsen/logrus"
)

type ConnectionProvider interface {
	GetDb() *gorm.DB
}

type ConnectionManager struct {
	db *gorm.DB
}

func init() {
	gorm.NowFunc = func() time.Time {
		return time.Now().UTC()
	}
}

func NewGormProvider(db *gorm.DB) *ConnectionManager {
	connectionManager := ConnectionManager{}
	connectionManager.db = db
	return &connectionManager
}

func (self *ConnectionManager) createDb(dbDialect string, dbConnectionString string) *gorm.DB {
	db, err := gorm.Open(dbDialect, dbConnectionString)
	if err != nil {
		log.Error("failed to connect database")
	}

	env := os.Getenv("GIN_ENV")
	if env == "development" || env == "test" {
		db.LogMode(true)
	}

	return db
}

func (self *ConnectionManager) GetDb() *gorm.DB {
	return self.db
}

func (self *ConnectionManager) Ping() error {
	return self.db.DB().Ping()
}

func (self *ConnectionManager) Close() error {
	return self.db.DB().Close()
}

func NewConnectionManager(dbExtractor extractors.DbExtractor) *ConnectionManager {
	connectionManager := ConnectionManager{}

	dbConnectionString := dbExtractor.ExtractConnectionString()
	dialect := dbExtractor.ExtractDialect()

	log.Info("Connecting to [", dbConnectionString, "], a [", dialect, "] database")
	connectionManager.db = connectionManager.createDb(string(dialect), dbConnectionString)

	return &connectionManager
}
