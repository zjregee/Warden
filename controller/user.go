package controller

import (
	"Warden/model"
	"Warden/tool/context"
	"Warden/tool/mysql"
	"io"
	"fmt"
	"time"
	"strconv"
	"net/http"
	"crypto/md5"

	"github.com/labstack/echo/v4"
)

func GenerateMD5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func GetNameFromToken(str string) string {
	model := model.UserInformation{}
	if err := mysql.DB.First(&model, "token = ?", str).Error; err == nil {
		return model.Name
	}
	return ""
}

func CheckPassword(account, password string) bool {
	model := model.UserInformation{}
	if err := mysql.DB.First(&model, "name = ?", account).Error; err == nil {
		return model.Password == password
	}
	return false
}

func LoginIn(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Account  string `json:"account" form:"account"`
		Password string `json:"password" form:"password"`
	} {}

	err := c.Bind(&u)
    if err != nil {
        return err
    }

	account := u.Account
	password := u.Password

	ok := CheckPassword(account, password)
	if !ok {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 100, 
			"token": "",
		})
	}

	token := GenerateMD5(account + strconv.FormatInt(time.Now().Unix(), 10))

	model := model.UserInformation{}
	mysql.DB.Model(&model).Where("name = ?", account).Update("token", token)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200, 
		"token": token,
	})
}

func LoginInfo(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Token  string `json:"token" form:"token"`
	} {}

	err := c.Bind(&u)
    if err != nil {
        return err
    }

	token := u.Token
	name := GetNameFromToken(token)
	model := model.EmployeeInformation{}

	if err = mysql.DB.First(&model, "name = ?", name).Error; err == nil {
		resData := map[string]interface{}{}
		resData["name"] = model.Name
		resData["roles"] = model.Type
		resData["introduction"] = model.Introduction
		resData["avator"] = model.HeadPortrait
		return context.RetData(c, resData)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 100,
	})
}

func LoginOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Token  string `json:"token" form:"token"`
	} {}

	err := c.Bind(&u)
    if err != nil {
        return err
    }

	token := u.Token
	name := GetNameFromToken(token)

	model := model.UserInformation{}
	mysql.DB.Model(&model).Where("name = ?", name).Update("token", "")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}