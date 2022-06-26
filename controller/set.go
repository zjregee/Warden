package controller

import (
	"Warden/model"
	"Warden/tool/context"
	"Warden/tool/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchEmployee(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name string `json:"name" form:"name"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.EmployeeInformation{}

	if err = mysql.DB.First(&model, "name = ?", u.Name).Error; err == nil {
		resData := map[string]interface{}{}
		resData["name"] = model.Name
		resData["sex"] = model.Sex
		resData["age"] = model.Age
		resData["type"] = model.Type
		resData["hire_date"] = model.HireDate
		resData["duty"] = model.Duty
		resData["job_title"] = model.JobTitle
		resData["telephone"] = model.Telephone
		resData["introduction"] = model.Introduction
		resData["remark"] = model.Remark
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
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

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
		mysql.DB.Save(&model)
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

	if err = mysql.DB.Create(&model).Error; err != nil {
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
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.EmployeeInformation{}

	if err = mysql.DB.First(&model, "name = ?", u.Name).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 100,
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
	mysql.DB.Save(&model)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}