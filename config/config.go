package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	//initialize sqlite
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("failed to initialize sqlite: %v", err)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
