package MySQL

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const usersTable = "user"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"userName"` // 用户名
	Password string `gorm:"type:varchar(20);not null" json:"password"` // 密码
	Role     int    `gorm:"type:int" json:"role"`                      // 角色
}

// WriteUser 存储用户
func WriteUser(userName string, password string) error {
	if len(userName) > 20 {
		return errors.New("userName wrongful")
	}
	if len(password) > 20 {
		return errors.New("password wrongful")
	}

	user := &User{
		UserName: userName,
		Password: password,
	}
	db.Table(usersTable).Create(user)
	return db.Error
}

// ReadUser 读取用户
func ReadUser(userName string) {
	var user User
	db.Table(usersTable).First(&user, "user_name=?", userName)
	fmt.Println(user)
	fmt.Println(user.ID, user.UserName, user.Password)
}
