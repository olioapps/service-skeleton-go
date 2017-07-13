package dao

import (
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/olioapps/service-skeleton-go/olio/util"
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
		panic("failed to connect database")
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

func NewConnectionManager() *ConnectionManager {
	connectionManager := ConnectionManager{}

	dbConnectionString := util.GetEnv("DB_CONNECTION_STRING", "root:root@/todo?parseTime=true")
	dialect := util.GetEnv("DB_CONNECTION_DIALECT", "mysql")
	connectionManager.db = connectionManager.createDb(dialect, dbConnectionString)

	return &connectionManager
}
