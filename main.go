package main

import (
	"Warden/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "expvar"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	ruzhu := e.Group("/ruzhu")
	ruzhu.POST("/search", controller.SearchCheckIn)
	ruzhu.POST("/add", controller.AddCheckIn)
	ruzhu.POST("/modify", controller.ModifyCheckIn)
	ruzhu.POST("/delete", controller.DeleteCheckIn)
	ruzhu.POST("/getnamebyid", controller.GetNameByID)
	ruzhu.POST("/getidbyname", controller.GetIDByName)
	ruzhu.POST("/searchall", controller.SearchCheckinAll)

	tuizhu := e.Group("/tuizhu")
	tuizhu.POST("/search", controller.SearchCheckOut)
	tuizhu.POST("/add", controller.AddCheckOut)
	tuizhu.POST("/modify", controller.ModifyCheckOut)
	tuizhu.POST("/delete", controller.DeleteCheckOut)
	tuizhu.POST("/approve", controller.ApproveCheckOut)
	tuizhu.POST("/searchall", controller.SearchCheckoutAll)

	waichu := e.Group("/waichu")
	waichu.POST("/search", controller.SearchOut)
	waichu.POST("/add", controller.AddOut)
	waichu.POST("/modify", controller.ModifyOut)
	waichu.POST("/delete", controller.DeleteOut)
	waichu.POST("/approve", controller.ApproveOut)
	waichu.POST("/register", controller.RegisterOut)
	waichu.POST("/searchall", controller.SearchOutAll)

	yonghu := e.Group("/yonghu")
	yonghu.POST("/search", controller.SearchEmployee)
	yonghu.POST("/add", controller.AddEmployee)
	yonghu.POST("/modify", controller.ModifyEmployee)
	yonghu.POST("/searchall", controller.SearchEmployeeAll)

	user := e.Group("/user")
	user.POST("/login", controller.LoginIn)
	user.POST("/info", controller.LoginInfo)
	user.POST("/logout", controller.LoginOut)

	e.Logger.Fatal(e.Start(":8000"))
}