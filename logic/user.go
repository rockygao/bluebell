package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		// 查库报错
		return err
	}
	// 生成UID
	userID := snowflake.GenID()
	// 创建一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	// 创建一个User实例
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 保存进数据库 传递的是一个指针，所以能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	//生成JWT
	return jwt.GenToken(user.UserID, user.Username)

}
