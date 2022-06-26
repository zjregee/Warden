package controller

import (
	"Warden/model"
	"Warden/tool/context"
	"Warden/tool/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name string `json:"name" form:"name"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.OutInformation{}

	if err = mysql.DB.First(&model, "name = ?", u.Name).Error; err == nil {
		resData := map[string]interface{}{}
		resData["name"] = model.Name
		resData["reason"] = model.Reason
		resData["out_time"] = model.OutTime
		resData["predict_back_time"] = model.PredictBackTime
		resData["people"] = model.People
		resData["relationship"] = model.RelationShip
		resData["telephone"] = model.Telephone
		resData["remark"] = model.Remark
		resData["fact_back_time"] = model.FactBackTime
		resData["is_pass"] = model.IsPass
		return context.RetData(c, resData)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 100,
	})
}

func AddOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name            string `json:"name" form:"name"`
		Reason          string `json:"reason" form:"reason"`
		OutTime         string `json:"out_time" form:"out_time"`
		PredictBackTime string `json:"predict_back_time" form:"predict_back_time"`
		People          string `json:"people" form:"people"`
		RelationShip    string `json:"relationship" form:"relationship"`
		Telephone       string `json:"telephone" form:"telephone"`
		Remark          string `json:"remark" form:"remark"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.OutInformation{}

	if err = mysql.DB.First(&model, "name = ?", u.Name).Error; err == nil {
		model.Name = u.Name
		model.Reason = u.Reason
		model.OutTime = u.OutTime
		model.PredictBackTime = u.PredictBackTime
		model.People = u.People
		model.RelationShip = u.RelationShip
		model.Telephone = u.Telephone
		model.Remark = u.Remark
		mysql.DB.Save(&model)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 200,
		})
	}

	model.Name = u.Name
	model.Reason = u.Reason
	model.OutTime = u.OutTime
	model.PredictBackTime = u.PredictBackTime
	model.People = u.People
	model.RelationShip = u.RelationShip
	model.Telephone = u.Telephone
	model.Remark = u.Remark

	if err = mysql.DB.Create(&model).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func ModifyOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name            string `json:"name" form:"name"`
		Reason          string `json:"reason" form:"reason"`
		OutTime         string `json:"out_time" form:"out_time"`
		PredictBackTime string `json:"predict_back_time" form:"predict_back_time"`
		People          string `json:"people" form:"people"`
		RelationShip    string `json:"relationship" form:"relationship"`
		Telephone       string `json:"telephone" form:"telephone"`
		Remark          string `json:"remark" form:"remark"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.OutInformation{}

	if err = mysql.DB.First(&model, "name = ?", u.Name).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 100,
		})
	}

	model.Name = u.Name
	model.Reason = u.Reason
	model.OutTime = u.OutTime
	model.PredictBackTime = u.PredictBackTime
	model.People = u.People
	model.RelationShip = u.RelationShip
	model.Telephone = u.Telephone
	model.Remark = u.Remark
	mysql.DB.Save(&model)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func DeleteOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name string `json:"name" form:"name"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.OutInformation{}
	mysql.DB.Where("name = ?", u.Name).Delete(&model)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func ApproveOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name   string `json:"name" form:"name"`
		IsPass string `json:"is_pass" form:"is_pass"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.OutInformation{}
	mysql.DB.Model(&model).Where("name = ?", u.Name).Update("is_pass", u.IsPass)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func RegisterOut(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name         string `json:"name" form:"name"`
		FactBackTime string `json:"fact_back_time" form:"fact_back_time"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	model := model.OutInformation{}
	mysql.DB.Model(&model).Where("name = ?", u.Name).Update("fact_back_time", u.FactBackTime)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}
