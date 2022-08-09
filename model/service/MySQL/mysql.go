package MySQL

import (
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
)

type DBConf struct {
	User     string
	PassWord string
	Host     string
	Port     string
}

// var db *sql.DB

func InitDB() *sql.DB {
	var (
		dbConf DBConf
		err    error
		dsn    string // Data Source Name
	)
	if _, err = toml.DecodeFile("./conf/MySQL.toml", &dbConf); err != nil {
		// 这里打个日志
		fmt.Println(err)
	}
	dsn = dbConf.User + ":" + dbConf.PassWord + "@tcp(" + dbConf.Host + ":" + dbConf.Port + ")/user"
	// dsn := "root:123456@tcp(127.0.0.1:3306)/user"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("123")
		fmt.Println(err)
	}
	return db
}

func InsertData(db *sql.DB) {
	sqlStr := "insert into user(user, password) value (?,?)"
	r, err := db.Exec(sqlStr, "yangxuqi", "123")
	if err != nil {
		fmt.Println(err)
	}
	_, err = r.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
}
