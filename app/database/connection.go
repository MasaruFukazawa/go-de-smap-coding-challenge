package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	connection *gorm.DB
}

func (db *DB) Close() error {
	sqlDB, err := db.connection.DB()

	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (db *DB) GetConnection() *gorm.DB {
    return db.connection
}

func NewDB(path string) (*DB, error) {

	connection, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DB{connection: connection}, nil
}
