package db

import (
	"vnet-mysql/config"
	"vnet-mysql/pkg/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := config.GetConfig().Database.MY_SQL
	// tlsConfig := &tls.Config{]
	//     InsecureSkipVerify: false,
	// }
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.User{})
	return db, nil
}
