package controller

import (
	"Warden/model"
	"Warden/tool/context"
	"Warden/tool/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchCheckoutAll(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	informations := []model.CheckOutInformation{}
	_ = mysql.DB.Find(&informations)

	temp := []interface{}{}
	for _, model := range informations {
		resData := map[string]interface{}{}
		resData["name"] = model.Name
		resData["check_out_type"] = model.CheckOutType
		resData["reason"] = model.Reason
		resData["out_time"] = model.OutTime
		resData["apply_time"] = model.ApplyTime
		resData["remark"] = model.Remark
		resData["is_pass"] = model.IsPass
		temp = append(temp, resData)
	}
	
	return context.RetData(c, temp)
}

func SearchCheckOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name string `json:"name" form:"name"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.CheckOutInformation{}

	if err = mysql.DB.First(&model, "name = ?", u.Name).Error; err == nil {
		resData := map[string]interface{}{}
		resData["name"] = model.Name
		resData["check_out_type"] = model.CheckOutType
		resData["reason"] = model.Reason
		resData["out_time"] = model.OutTime
		resData["apply_time"] = model.ApplyTime
		resData["remark"] = model.Remark
		resData["is_pass"] = model.IsPass
		return context.RetData(c, resData)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 100,
	})
}

func AddCheckOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name         string `json:"name" form:"name"`
		CheckOutType string `json:"check_out_type" form:"check_out_type"`
		Reason       string `json:"reason" form:"reason"`
		OutTime      string `json:"out_time" form:"out_time"`
		ApplyTime    string `json:"apply_time" form:"apply_time"`
		Remark       string `json:"remark" form:"remark"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.CheckOutInformation{}

	if err = mysql.DB.First(&model, "name = ?", u.Name).Error; err == nil {
		model.Name = u.Name
		model.CheckOutType = u.CheckOutType
		model.Reason = u.Reason
		model.OutTime = u.OutTime
		model.ApplyTime = u.ApplyTime
		model.Remark = u.Remark
		mysql.DB.Save(&model)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 200,
		})
	}

	model.Name = u.Name
	model.CheckOutType = u.CheckOutType
	model.Reason = u.Reason
	model.OutTime = u.OutTime
	model.ApplyTime = u.ApplyTime
	model.Remark = u.Remark
	model.IsPass = "0"

	if err = mysql.DB.Create(&model).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func ModifyCheckOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name         string `json:"name" form:"name"`
		CheckOutType string `json:"check_out_type" form:"check_out_type"`
		Reason       string `json:"reason" form:"reason"`
		OutTime      string `json:"out_time" form:"out_time"`
		ApplyTime    string `json:"apply_time" form:"apply_time"`
		Remark       string `json:"remark" form:"remark"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.CheckOutInformation{}

	if err = mysql.DB.First(&model, "name = ?", u.Name).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 100,
		})
	}

	model.Name = u.Name
	model.CheckOutType = u.CheckOutType
	model.Reason = u.Reason
	model.OutTime = u.OutTime
	model.ApplyTime = u.ApplyTime
	model.Remark = u.Remark
	mysql.DB.Save(&model)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func DeleteCheckOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name string `json:"name" form:"name"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.CheckOutInformation{}
	mysql.DB.Where("name = ?", u.Name).Delete(&model)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func ApproveCheckOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name   string `json:"name" form:"name"`
		IsPass string `json:"is_pass" form:"is_pass"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.CheckOutInformation{}
	mysql.DB.Model(&model).Where("name = ?", u.Name).Update("is_pass", u.IsPass)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}
