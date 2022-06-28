package controller

import (
	"Warden/model"
	"Warden/tool/context"
	"Warden/tool/mysql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SearchCheckinAll(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	informations := []model.OldInformation{}
	_ = mysql.DB.Find(&informations)

	temp := []interface{}{}
	for _, information := range informations {
		resData := map[string]interface{}{}
		resData["id"] = information.ID
		resData["name"] = information.Name
		resData["age"] = information.Age
		resData["sex"] = information.Sex
		resData["identification"] = information.Identification
		resData["room"] = information.Room
		resData["building"] = information.Building
		resData["file_id"] = information.FileId
		resData["old_type"] = information.OldType
		resData["check_in_time"] = information.CheckInTime
		resData["expiration_time"] = information.ExpirationTime
		resData["telephone"] = information.Telephone
		resData["birthday"] = information.Birthday
		resData["height"] = information.Height
		resData["marriage"] = information.Marriage
		resData["weight"] = information.Weight
		resData["blood_type"] = information.BloodType
		resData["head_portrait"] = information.HeadPortrait
		resData["remark"] = information.Remark
		temp = append(temp, resData)
	}
	return context.RetData(c, temp)
}

func SearchCheckIn(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name string `json:"name" form:"name"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	information := model.OldInformation{}

	if err = mysql.DB.First(&information, "name = ?", u.Name).Error; err == nil {
		resData := map[string]interface{}{}
		resData["id"] = information.ID
		resData["name"] = information.Name
		resData["age"] = information.Age
		resData["sex"] = information.Sex
		resData["identification"] = information.Identification
		resData["room"] = information.Room
		resData["building"] = information.Building
		resData["file_id"] = information.FileId
		resData["old_type"] = information.OldType
		resData["check_in_time"] = information.CheckInTime
		resData["expiration_time"] = information.ExpirationTime
		resData["telephone"] = information.Telephone
		resData["birthday"] = information.Birthday
		resData["height"] = information.Height
		resData["marriage"] = information.Marriage
		resData["weight"] = information.Weight
		resData["blood_type"] = information.BloodType
		resData["head_portrait"] = information.HeadPortrait
		resData["remark"] = information.Remark
		return context.RetData(c, resData)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 100,
	})
}

func AddCheckIn(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name           string `json:"name" form:"name"`
		Age            string `json:"age" form:"age"`
		Sex            string `json:"sex" form:"sex"`
		Identification string `json:"identification" form:"identification"`
		Room           string `json:"room" form:"room"`
		Building       string `json:"building" form:"building"`
		FileId         string `json:"file_id" form:"file_id"`
		OldType        string `json:"old_type" form:"old_type"`
		CheckInTime    string `json:"check_in_time" form:"check_in_time"`
		ExpirationTime string `json:"expirtion_time" form:"expirtion_time"`
		Telephone      string `json:"telephone" form:"telephone"`
		Birthday       string `json:"birthday" form:"birthday"`
		Height         string `json:"height" form:"height"`
		Marriage       string `json:"marriage" form:"marriage"`
		Weight         string `json:"weight" form:"weight"`
		BloodType      string `json:"blood_type" form:"blood_type"`
		HeadPortrait   string `json:"head_portrait" form:"head_portrait"`
		Remark         string `json:"remark" form:"remark"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	information := model.OldInformation{}

	if err = mysql.DB.First(&information, "name = ?", u.Name).Error; err == nil {
		information.Name = u.Name
		information.Age = u.Age
		information.Sex = u.Sex
		information.Identification = u.Identification
		information.Room = u.Room
		information.Building = u.Building
		information.FileId = u.FileId
		information.OldType = u.OldType
		information.CheckInTime = u.CheckInTime
		information.ExpirationTime = u.ExpirationTime
		information.Telephone = u.Telephone
		information.Birthday = u.Birthday
		information.Height = u.Height
		information.Marriage = u.Marriage
		information.Weight = u.Weight
		information.BloodType = u.BloodType
		information.HeadPortrait = u.HeadPortrait
		information.Remark = u.Remark
		mysql.DB.Save(&information)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 200,
		})
	}

	information.Name = u.Name
	information.Age = u.Age
	information.Sex = u.Sex
	information.Identification = u.Identification
	information.Room = u.Room
	information.Building = u.Building
	information.FileId = u.FileId
	information.OldType = u.OldType
	information.CheckInTime = u.CheckInTime
	information.ExpirationTime = u.ExpirationTime
	information.Telephone = u.Telephone
	information.Birthday = u.Birthday
	information.Height = u.Height
	information.Marriage = u.Marriage
	information.Weight = u.Weight
	information.BloodType = u.BloodType
	information.HeadPortrait = u.HeadPortrait
	information.Remark = u.Remark

	if err = mysql.DB.Create(&information).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func ModifyCheckIn(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name           string `json:"name" form:"name"`
		Age            string `json:"age" form:"age"`
		Sex            string `json:"sex" form:"sex"`
		Identification string `json:"identification" form:"identification"`
		Room           string `json:"room" form:"room"`
		Building       string `json:"building" form:"building"`
		FileId         string `json:"file_id" form:"file_id"`
		OldType        string `json:"old_type" form:"old_type"`
		CheckInTime    string `json:"check_in_time" form:"check_in_time"`
		ExpirationTime string `json:"expirtion_time" form:"expirtion_time"`
		Telephone      string `json:"telephone" form:"telephone"`
		Birthday       string `json:"birthday" form:"birthday"`
		Height         string `json:"height" form:"height"`
		Marriage       string `json:"marriage" form:"marriage"`
		Weight         string `json:"weight" form:"weight"`
		BloodType      string `json:"blood_type" form:"blood_type"`
		HeadPortrait   string `json:"head_portrait" form:"head_portrait"`
		Remark         string `json:"remark" form:"remark"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	information := model.OldInformation{}

	if err = mysql.DB.First(&information, "name = ?", u.Name).Error; err != nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 100,
		})
	}

	information.Name = u.Name
	information.Age = u.Age
	information.Sex = u.Sex
	information.Identification = u.Identification
	information.Room = u.Room
	information.Building = u.Building
	information.FileId = u.FileId
	information.OldType = u.OldType
	information.CheckInTime = u.CheckInTime
	information.ExpirationTime = u.ExpirationTime
	information.Telephone = u.Telephone
	information.Birthday = u.Birthday
	information.Height = u.Height
	information.Marriage = u.Marriage
	information.Weight = u.Weight
	information.BloodType = u.BloodType
	information.HeadPortrait = u.HeadPortrait
	information.Remark = u.Remark
	mysql.DB.Save(&information)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func DeleteCheckIn(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name string `json:"name" form:"name"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	information := model.OldInformation{}
	mysql.DB.Where("name = ?", u.Name).Delete(&information)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 200,
	})
}

func GetIDByName(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		Name string `json:"name" form:"name"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	information := model.OldInformation{}

	if err = mysql.DB.First(&information, "name = ?", u.Name).Error; err == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 200,
			"id": information.ID,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 100,
	})
}

func GetNameByID(c echo.Context) error {
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")

	var u = struct {
		ID string `json:"id" form:"id"`
	}{}

	err := c.Bind(&u)
	if err != nil {
		return err
	}

	information := model.OldInformation{}

	if err = mysql.DB.First(&information, "ID = ?", u.ID).Error; err == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 200,
			"name": information.Name,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": 100,
	})
}