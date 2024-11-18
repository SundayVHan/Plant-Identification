package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"plant_identification/internal/common"
	"plant_identification/internal/database"
	"plant_identification/internal/util"
)

func Init() {
	err := database.DB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}

func RegisterUser(userName string, password string, kind int64) error {
	// 将密码哈希后存入
	_, err := getUser(userName)
	if err == nil {
		return common.CustomError{
			Code:    common.ErrUsernameUsed,
			Message: "Username has been used",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("could not hash password: %v", err)
	}

	user := User{
		Username: userName,
		Password: string(hashedPassword),
		Kind:     kind,
	}

	return saveUser(user)
}

func LoginUser(userName string, password string) (int64, error) {
	user, err := getUser(userName)
	if err != nil {
		return 0, common.CustomError{
			Message: "User not registered",
			Code:    common.ErrUserNotRegistered,
		}
	}

	// 验证输入的密码和数据库中的哈希密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, common.CustomError{
				Message: "Password mismatch",
				Code:    common.ErrPasswordMismatch,
			}
		} else {
			return 0, err
		}

	}

	// 密码匹配，认证成功
	return user.Kind, nil
}

func RegisterAndIssueToken(username, password string, kind int64) (string, error) {
	err := RegisterUser(username, password, kind)
	if err != nil {
		return "", err // 返回错误给外层
	}

	// 无错误则说明注册成功，签发JWT
	token, err := util.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func LoginAndIssueToken(username, password string) (string, int64, error) {
	// 首先进行用户验证
	kind, err := LoginUser(username, password)
	if err != nil {
		return "", 0, err // 返回错误给外层
	}

	// 假设用户登录成功，生成JWT
	token, err := util.GenerateToken(username)
	if err != nil {
		return "", 0, err
	}

	return token, kind, nil
}

func GetUser(username string) (User, error) {
	return getUser(username)
}
