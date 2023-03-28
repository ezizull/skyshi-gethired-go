// Package postgres provides the database connection
package mysql

import (
	"gorm.io/gorm"
)

// NewGorm is a function that returns a gorm database connection using  initial configuration
func NewGorm() (gormDB *gorm.DB, err error) {
	var infoPg infoDatabaseMySQL
	err = infoPg.getMysqlConn("Databases.MySQL.Localhost")
	if err != nil {
		return nil, err
	}

	gormDB, err = initMysqlDB(gormDB, infoPg)
	if err != nil {
		return nil, err
	}

	gormDB, err = migrateMysql(gormDB)
	if err != nil {
		return nil, err
	}

	var result int
	// Test the connection by executing a simple query
	if err = gormDB.Raw("SELECT 1").Scan(&result).Error; err != nil {
		return nil, err
	}

	return
}
