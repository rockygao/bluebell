package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret = "bluebell"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户或密码错误")
)

func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	user.Password = encryptPassword(user.Password)
	// 执行语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func CheckUserExist(username string) (err error) {
	// 执行语句入库
	sqlStr := `select count(user_id) from user where username = ?`

	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}

	if count > 0 {
		return ErrorUserExist
	}
	return
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	// 用户不存在
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	// 查询报错
	if err != nil {
		return err
	}
	// 判断密码
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}

	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
