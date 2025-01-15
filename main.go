package main

import (
	challenge2 "backend-challenge/pkg/challenge-2"
	challenge3 "backend-challenge/pkg/challenge-3"
	"backend-challenge/pkg/challenge"

	"github.com/labstack/echo/v4"
)

func main() {

	// สร้าง instance ของ Echo
	e := echo.New()

	// สร้าง route และเชื่อมต่อกับ handler function
	e.POST("/number/summary", challenge.Handler)

	// สร้าง route และเชื่อมต่อกับ handler function
	e.POST("/keyboard-encoded", challenge2.Handler)

	// สร้าง route และเชื่อมต่อกับ handler function
	e.GET("/beef/summary", challenge3.Handler)

	// เริ่ม server
	e.Logger.Fatal(e.Start(":8080"))
}
