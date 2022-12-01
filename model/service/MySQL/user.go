package mysql

import (
	"encoding/base64"
	"errors"
	"log"
	"main/commond/errmsg"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

const usersTable = "user"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);not null" json:"userName"` // 用户名
	Password string `gorm:"type:varchar(20);not null" json:"password"` // 密码
	Role     int    `gorm:"type:int" json:"role"`                      // 角色
}

// UserExist 查询用户是否存在
func UserExist(userName string) int {
	var users User
	db.Table(usersTable).Where("user_name = ?", userName).First(&users)
	// 用户名已存在
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// AddUser 添加用户
func AddUser(userData *User) int {
	userData.Password = ScryptPassword(userData.Password)
	err := db.Table(usersTable).Create(&userData).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// WriteUser 存储用户
// 和上面的添加用户合并
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

// GetUsers 查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	// 分页
	err = db.Table(usersTable).Limit(pageSize).Offset(offset).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// hook 函数
// 参考：https://gorm.io/docs/hooks.html
// func (u *User) BeforeSave(_ *gorm.DB) (err error) {
// 	u.Password = ScryptPassword(u.Password)
// 	return nil
// }

// ScryptPassword 密码加密
// 参考：https://learnku.com/docs/build-web-application-with-golang/95-storage-password/3213
// 还有一个可以考虑的方案：bcrypt
// TODO: 还没有研究性能开销
func ScryptPassword(password string) string {
	const kenLen = 10
	salt := []byte{11, 12, 13, 14, 15, 16, 17, 18}
	hashPassword, err := scrypt.Key([]byte(password), salt, 4, 8, 1, kenLen)
	if err != nil {
		log.Fatal(err) // TODO 这里需要处理下错误
	}

	return base64.StdEncoding.EncodeToString(hashPassword)
}

// DeleteUser 删除用户（软删除）
func DeleteUser(id int) int {
	err = db.Table(usersTable).Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// EditUser 编辑用户
func EditUser(id int, userData *User) int {
	var maps = make(map[string]any)
	maps["user_name"] = userData.UserName
	maps["role"] = userData.Role
	err = db.Table(usersTable).Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 修改密码
func ChangePassword() {

}

// VerifyLogin
func VerifyLogin(userName string, password string) int {
	var user User
	db.Table(usersTable).Where("user_name = ?", userName).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPassword(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCSE
}
