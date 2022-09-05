package MySQL

import (
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DBConf struct {
	User     string
	PassWord string
	Host     string
	Port     string
	DBName   string
}

var (
	db  *gorm.DB // 全局数据库
	err error    // 全局错误
)

// InitDB 初始化数据库
func InitDB() {
	var dbConf DBConf
	if _, err = toml.DecodeFile("./conf/mysql.toml", &dbConf); err != nil {
		// TODO 打日志
		fmt.Println(err)
	}
	// data source name
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.PassWord,
		dbConf.Host,
		dbConf.Port,
		dbConf.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`，而不是 `users`
			SingularTable: true,
		},
	})
	if err != nil {
		// TODO 打日志
		fmt.Println("连接数据库失败，请检查参数：", err)
	}

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	// db.AutoMigrate(&User{}, &Category{}, &Article{})

	sqlDB, _ := db.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
