package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!\n")
	})
	e.GET("/omikuji", func(c echo.Context) error {
		rand.Seed(time.Now().UnixNano())
    omikujiNum:=rand.Intn(6)
		omikujiResult:=""
		switch omikujiNum{
		case 0:
			omikujiResult="大吉"
		case 1:
			omikujiResult="中吉"
		case 2:
			omikujiResult="小吉"
		case 3:
			omikujiResult="吉"
		case 4:
			omikujiResult="凶"
		case 5:
			omikujiResult="大凶"
		}

		return c.String(http.StatusOK, omikujiResult+"\n")
	})
    e.Logger.Fatal(e.Start(":"+os.Getenv("PORT")))
}