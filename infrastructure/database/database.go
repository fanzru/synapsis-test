package database

import (
	"synapsis-test/infrastructure/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Connection struct {
	DB *gorm.DB
}

func New(cfg config.Config) (Connection, error) {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.Database.DBUser, cfg.Database.DBPass, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second,   // Slow SQL threshold
	// 		LogLevel:                  logger.Silent, // Log level
	// 		IgnoreRecordNotFoundError: false,         // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      false,         // Don't include params in the SQL log
	// 		Colorful:                  true,          // Disable color
	// 	},
	// )

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: newLogger,
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Cannot open MYSQL")
	}
	log.Println("Databases Connected...")

	return Connection{
		DB: db,
	}, nil
}
