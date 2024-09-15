package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"test/pkg/setting"
)

//初始化数据库

type Model struct {
	ID       int `gorm:"primary_key" json:"id"`
	CreateOn int `json:"create_on"`
	UpdateOn int `json:"update_on"`
}

func init() {
	sec, _ := setting.Cfg.GetSection("database")
	user := sec.Key("USER")
	password := sec.Key("PASSWORD")
	host := sec.Key("HOST")
	dbName := sec.Key("NAME")
	tablePrefix := sec.Key("TABLE_PREFIX").String()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 表前缀，例如 "blog_"
			SingularTable: true,        // 不使用复数表名
		},
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}
