package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQLDB struct {
	DB *gorm.DB
}

type MySQLConnectionOptions struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func NewMySQLDBConnection() *MySQLDB {
	return &MySQLDB{}
}

func (d *MySQLDB) Connect(options MySQLConnectionOptions) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		options.User,
		options.Password,
		options.Host,
		options.Port,
		options.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return err
	}
	d.DB = db
	return nil
}

func (d *MySQLDB) Close() error {
	db, err := d.DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
