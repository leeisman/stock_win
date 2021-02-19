package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connections for gorm v2 ...
type Connections struct {
	ReadDB  *gorm.DB
	WriteDB *gorm.DB
}

func InitDatabases(cfg *Config) (*Connections, error) {
	var (
		conn Connections
		err  error
	)
	conn.ReadDB, err = SetupDatabase(cfg.Read)
	if err != nil {
		return nil, err
	}

	conn.WriteDB, err = SetupDatabase(cfg.Write)
	if err != nil {
		return nil, err
	}
	return &conn, err
}

func SetupDatabase(database *Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.Username,
		database.Password,
		database.Host,
		database.Port,
		database.DBName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
