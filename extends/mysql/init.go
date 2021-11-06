package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go-Gateway/models"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func (db *Database) InitDB() error {
	d, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=%t&loc=%s",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.addr"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.database"),
		true,
		"Local",
	))
	if err != nil {
		return fmt.Errorf("Database connection failed. Database name: %s", viper.GetString("mysql.database"))
	}
	//db.DB().SetMaxOpenConns(20000) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	d.DB().SetMaxIdleConns(0) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	DB = &Database{
		Self: d,
	}
	// 自动迁移
	DB.autoMigrate()
	return nil
}

func (db *Database) autoMigrate() {
	db.Self.AutoMigrate(
		&models.UserTables{},
	)
}
