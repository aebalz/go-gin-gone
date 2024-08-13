package database

import (
	"fmt"
	"log"
	"time"

	"github.com/aebalz/go-gin-gone/configs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	logGorm "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type MySQLDatabase struct {
	db *gorm.DB
}

func NewMySQLDatabase() Database {
	return &MySQLDatabase{}
}

func setupLoggerMode(isDebug bool) logGorm.Interface {
	if isDebug {
		return logGorm.Default.LogMode(logGorm.Info)
	}
	return logGorm.Default.LogMode(logGorm.Silent)
}

func isTraceResolverMode(isDebug bool) bool {
	return isDebug
}

func (m *MySQLDatabase) Connect(config configs.Config, isDebug bool) error {
	configDb := &gorm.Config{
		Logger:  setupLoggerMode(isDebug),
		NowFunc: func() time.Time { return time.Now().UTC() },
	}
	dsn := m.buildDBConnection(configs.AppConfig)

	db, err := gorm.Open(mysql.Open(dsn), configDb)
	if err != nil {
		return err
	}

	err = db.Use(dbresolver.Register(
		dbresolver.Config{
			Replicas:          []gorm.Dialector{mysql.Open(dsn)},
			Sources:           []gorm.Dialector{mysql.Open(dsn)},
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: isTraceResolverMode(isDebug),
		}).
		SetConnMaxLifetime(5 * time.Minute).
		SetMaxIdleConns(10).
		SetMaxOpenConns(100))

	if err != nil {
		return err
	}

	// Assign the connected database instance to the 'db' field
	m.db = db

	return nil
}

func (m *MySQLDatabase) Close() error {
	if m.db != nil {
		sqlDB, _ := m.db.DB()
		log.Println("Shutting down database...")
		sqlDB.Close()
		log.Println("Database exiting...")
	}

	return nil
}

func (m *MySQLDatabase) GetDB() *gorm.DB {
	return m.db
}

func (m *MySQLDatabase) GetConfig() gorm.Dialector {
	return m.db.Dialector
}

func (m *MySQLDatabase) buildDBConnection(db configs.Config) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		db.MysqlUser, db.MysqlPassword, db.MysqlHost, db.MysqlPort, db.MysqlDb,
	)
}
