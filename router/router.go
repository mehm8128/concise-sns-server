package router

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mehm8128/concise-sns-server/model"
)

func SetRouting() {
	// サーバーのインスタンス作成
	e := echo.New()
	port := os.Getenv("PORT")
	fmt.Println(port)
	//CORSエラー回避
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "https://concise-sns.vercel.app"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	// ルーティング設定
	e.GET("/login", model.Login())
	e.GET("/secret", model.Secret())
	e.GET("/logout", model.Logout())
	e.GET("/get", model.GetAllPosts)
	e.POST("/post", model.PostContent)
	e.POST("/delete", model.DeletePost)
	// サーバー起動
	e.Logger.Fatal(e.Start(":" + port))
}
