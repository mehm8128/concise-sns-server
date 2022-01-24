package router

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mehm8128/concise-sns-server/model"
)

func SetRouting(){
	// サーバーのインスタンス作成
    e := echo.New()
    //CORSエラー回避
    e.Use(middleware.Logger())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000","https://concise-sns.vercel.app"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
    // ルーティング設定
    e.GET("/get", model.GetAllPosts)
    e.POST("/post", model.PostContent)
    // サーバー起動
    e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}