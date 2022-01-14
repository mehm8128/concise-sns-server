package main

import (
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Post struct {
  gorm.Model
  Name string `json:"Name"`
  Content string `json:"Content"`
}

// DBのインスタンスをグローバル変数に格納
var (
	db *gorm.DB
)

func main() {
    // DB接続処理
    var err error
    db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL") )
    if err != nil {
        panic("failed to connect database")
    }
    // サーバーが終了したらDB接続も終了する
    defer db.Close()
		db.AutoMigrate(&Post{})
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
    e.GET("/get", getAllPosts)
    e.POST("/post",post)
    // サーバー起動
    e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func post(c echo.Context) error {
		data := new(Post)
		err := c.Bind(data)
		if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest,"error")
	}
    // レコード登録
    var post = Post{Name:data.Name,Content:data.Content}
		db.Create(&post)
    return c.JSON(http.StatusOK, data)
}

func getAllPosts(c echo.Context) error {
    var posts []*Post
    // userテーブルのレコードを全件取得
    db.Find(&posts)
    // 取得したデータをJSONにして返却
    return echo.NewHTTPError(http.StatusOK, posts)
}