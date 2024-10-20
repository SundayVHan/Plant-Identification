package user

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"plant_identification/internal/database"
	"plant_identification/internal/util"
)

func Init() {
	err := database.DB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}

func RegisterUser(userName string, password string) error {
	// 将密码哈希后存入
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("could not hash password: %v", err)
	}

	user := User{
		Username: userName,
		Password: string(hashedPassword),
		Kind:     0,
	}

	return saveUser(user)
}

func LoginUser(userName string, password string) (bool, error) {
	user, err := getUser(userName)
	if err != nil {
		return false, fmt.Errorf("could not find user: %v", err)
	}

	// 验证输入的密码和数据库中的哈希密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			// 密码不匹配
			return false, nil
		}
		return false, fmt.Errorf("could not compare password: %v", err)
	}

	// 密码匹配，认证成功
	return true, nil
}

func RegisterAndIssueToken(username, password string) (string, error) {
	err := RegisterUser(username, password)
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

func LoginAndIssueToken(username, password string) (string, error) {
	// 首先进行用户验证
	isAuthenticated, err := LoginUser(username, password)
	if err != nil {
		return "", err // 返回错误给外层
	}

	if !isAuthenticated {
		return "", errors.New("invalid username or password")
	}

	// 假设用户登录成功，生成JWT
	token, err := util.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(username string) (User, error) {
	return getUser(username)
}
