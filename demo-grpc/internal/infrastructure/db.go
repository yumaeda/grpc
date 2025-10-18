package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DatabaseConfig is a configuration of the Database connection.
type DatabaseConfig struct {
	Password string `json:"db.password"`
	Host     string `json:"db.host"`
	Name     string `json:"db.name"`
	User     string `json:"db.user"`
}

// ConnectToDB connects to the TiDB based on the configuration and returns pointer to the connection.
func ConnectToDB() (*gorm.DB, func(), error) {
	secretManagerJSON := os.Getenv("TIDB_CONFIG_JSON")
	dbConfig := DatabaseConfig{}
	json.Unmarshal([]byte(secretManagerJSON), &dbConfig)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:4000)/%s?tls=true&charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		return nil, nil, sqlDBErr
	}
	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(6)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, func() {
		sqlDB.Close()
	}, nil
}
