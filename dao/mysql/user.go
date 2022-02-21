package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secret = "bluebell"

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
		return errors.New("用户已存在")
	}
	return
}

func SignUp() {
	// 执行语句入库
}
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
