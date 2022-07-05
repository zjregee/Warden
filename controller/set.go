package controller

import (
	"Warden/model"
	"Warden/tool/context"
	"Warden/tool/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchEmployeeAll(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	
	informations := []model.EmployeeInformation{}
	_ = mysql.DB.Find(&informations)

	temp := []interface{}{}
	for _, information := range informations {
		resData := map[string]interface{}{}
		resData["name"] = information.Name
		resData["sex"] = information.Sex
		resData["age"] = information.Age
		resData["type"] = information.Type
		resData["hire_date"] = information.HireDate
		resData["duty"] = information.Duty
		resData["job_title"] = information.JobTitle
		resData["telephone"] = information.Telephone
		resData["introduction"] = information.Introduction
		resData["remark"] = information.Remark

		user := model.UserInformation{}
		if err := mysql.DB.First(&user, "name = ?", information.Name).Error; err == nil {
			resData["password"] = user.Password
		} else {
			resData["password"] = ""
		}

		temp = append(temp, resData)
	}
	
	return context.RetData(c, temp)
}

func SearchEmployee(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name string `json:"name" form:"name"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	information := model.EmployeeInformation{}

	if err = mysql.DB.First(&information, "name = ?", u.Name).Error; err == nil {
		resData := map[string]interface{}{}
		resData["name"] = information.Name
		resData["sex"] = information.Sex
		resData["age"] = information.Age
		resData["type"] = information.Type
		resData["hire_date"] = information.HireDate
		resData["duty"] = information.Duty
		resData["job_title"] = information.JobTitle
		resData["telephone"] = information.Telephone
		resData["introduction"] = information.Introduction
		resData["remark"] = information.Remark

		user := model.UserInformation{}
		if err := mysql.DB.First(&user, "name = ?", information.Name).Error; err == nil {
			resData["password"] = user.Password
		} else {
			resData["password"] = ""
		}

		return context.RetData(c, resData)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 100,
	})
}

func AddEmployee(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name         string `json:"name" form:"name"`
		Sex          string `json:"sex" form:"sex"`
		Age          string `json:"age" form:"age"`
		Type         string `json:"type" form:"type"`
		HireDate     string `json:"hire_date" form:"hire_date"`
		Duty         string `json:"duty" form:"duty"`
		JobTitle     string `json:"job_title" form:"job_title"`
		Telephone    string `json:"telephone" form:"telephone"`
		Introduction string `json:"introduction" form:"introduction"`
		Remark       string `json:"remark" form:"remark"`
		Password     string `json:"password" form:"password"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	user := model.UserInformation{}
	model := model.EmployeeInformation{}

	if err = mysql.DB.First(&model, "name = ?", u.Name).Error; err == nil {
		model.Name = u.Name
		model.Sex = u.Sex
		model.Age = u.Age
		model.Type = u.Type
		model.HireDate = u.HireDate
		model.Duty = u.Duty
		model.JobTitle = u.JobTitle
		model.Telephone = u.Telephone
		model.Introduction = u.Introduction
		model.Remark = u.Remark
		user.Name = u.Name;
		user.Password = u.Password;
		user.Token = "";
		mysql.DB.Save(&model)
		mysql.DB.Save(&user)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 200,
		})
	}

	model.Name = u.Name
	model.Sex = u.Sex
	model.Age = u.Age
	model.Type = u.Type
	model.HireDate = u.HireDate
	model.Duty = u.Duty
	model.JobTitle = u.JobTitle
	model.Telephone = u.Telephone
	model.Introduction = u.Introduction
	model.Remark = u.Remark
	user.Name = u.Name;
	user.Password = u.Password;
	user.Token = "";

	if err = mysql.DB.Create(&model).Error; err != nil {
		return err
	}

	if err = mysql.DB.Create(&user).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func ModifyEmployee(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name         string `json:"name" form:"name"`
		Sex          string `json:"sex" form:"sex"`
		Age          string `json:"age" form:"age"`
		Type         string `json:"type" form:"type"`
		HireDate     string `json:"hire_date" form:"hire_date"`
		Duty         string `json:"duty" form:"duty"`
		JobTitle     string `json:"job_title" form:"job_title"`
		Telephone    string `json:"telephone" form:"telephone"`
		Introduction string `json:"introduction" form:"introduction"`
		Remark       string `json:"remark" form:"remark"`
		Password     string `json:"password" form:"password"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	user := model.UserInformation{}
	model := model.EmployeeInformation{}

	if err = mysql.DB.First(&model, "name = ?", u.Name).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 100,
		})
	}

	if err = mysql.DB.First(&user, "name = ?", u.Name).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 100,
		})
	}

	user.Name = u.Name;
	user.Password = u.Password;
	user.Token = "";

	model.Name = u.Name
	model.Sex = u.Sex
	model.Age = u.Age
	model.Type = u.Type
	model.HireDate = u.HireDate
	model.Duty = u.Duty
	model.JobTitle = u.JobTitle
	model.Telephone = u.Telephone
	model.Introduction = u.Introduction
	model.Remark = u.Remark
	mysql.DB.Save(&model)
	mysql.DB.Save(&user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}