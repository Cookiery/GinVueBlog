package MySQL

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type DataBase struct {
	Host     string
	Port     int
	User     string
	PassWord string
}

func DB() {
	var dataBase map[string]DataBase
	if _, err := toml.DecodeFile("./conf/db.toml", &dataBase); err != nil {
		fmt.Println(err)
	}
	fmt.Println(dataBase)
}
