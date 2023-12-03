package db

import (
	"fmt"
	"gogintemplate/logger"
	"gorm.io/driver/postgres"
	"strconv"

	"gogintemplate/env"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	logger := logger.GetLogger()

	dbPort, err := strconv.Atoi(env.Env.DbPort)
	if err != nil {
		logger.Error(err.Error())
	}

	pgDB, err := GetClient(env.Env.DbHost, env.Env.DbUsername, env.Env.DbPassword, env.Env.DbName, dbPort, "disable")

	if err != nil {
		logger.Error(err.Error())
	}

	db = pgDB
}

func GetConnection() *gorm.DB {
	return db
}

func GetClient(host string, user string, password string, dbName string, port int, sslmode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		host, user, password, dbName, port, sslmode)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func CloseConnection(db *gorm.DB) {
	logger := logger.GetLogger()

	dbSQL, err := db.DB()

	if err != nil {
		logger.Error(err.Error())

		return
	}

	if err = dbSQL.Close(); err != nil {
		logger.Error(err.Error())

		return
	}
}
