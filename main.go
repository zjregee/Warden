package main

import (
	"Warden/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	ruzhu := e.Group("/ruzhu")
	ruzhu.POST("/search", controller.SearchCheckIn)
	ruzhu.POST("/add", controller.AddCheckIn)
	ruzhu.POST("/modify", controller.ModifyCheckIn)
	ruzhu.POST("/delete", controller.DeleteCheckIn)

	tuizhu := e.Group("/tuizhu")
	tuizhu.POST("/search", controller.SearchCheckOut)
	tuizhu.POST("/add", controller.AddCheckOut)
	tuizhu.POST("/modify", controller.ModifyCheckOut)
	tuizhu.POST("/delete", controller.DeleteCheckOut)
	tuizhu.POST("/approve", controller.ApproveCheckOut)

	waichu := e.Group("/waichu")
	waichu.POST("/search", controller.SearchOut)
	waichu.POST("/add", controller.AddOut)
	waichu.POST("/modify", controller.ModifyOut)
	waichu.POST("/delete", controller.DeleteOut)
	waichu.POST("/approve", controller.ApproveOut)
	waichu.POST("/register", controller.RegisterOut)

	yonghu := e.Group("/yonghu")
	yonghu.POST("/search", controller.SearchEmployee)
	yonghu.POST("/add", controller.AddEmployee)
	yonghu.POST("/modify", controller.ModifyEmployee)

	user := e.Group("/user")
	user.POST("/login", controller.LoginIn)
	user.POST("/info", controller.LoginInfo)
	user.POST("/logout", controller.LoginOut)

	e.Logger.Fatal(e.Start(":8000"))
}