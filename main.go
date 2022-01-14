package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

type User struct {
  gorm.Model
  Name string `json:"Name"`
}

// DBのインスタンスをグローバル変数に格納
var (
	db *gorm.DB
)

func main() {
    // DB接続処理
    var err error
    db, err = gorm.Open("sqlite3", "/tmp/gorm.db")
    if err != nil {
        panic("failed to connect database")
    }
    // サーバーが終了したらDB接続も終了する
    defer db.Close()
		db.AutoMigrate(&User{})
    // サーバーのインスタンス作成
    e := echo.New()
    // ルーティング設定
    e.GET("/contents", getAllContents)
    e.POST("/create", createContent)
    // サーバー起動
    //e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
    e.Logger.Fatal(e.Start(":8080"))
}

func createContent(c echo.Context) error {
		data := new(User)
		err := c.Bind(data)
    // レコード登録
		if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest,"error")
	}
		db.Create(&User{Name:data.Name})
    return c.JSON(http.StatusOK, data)
}

func getAllContents(c echo.Context) error {
    var user User
    // userテーブルのレコードを全件取得
    db.Find(&user)
    // 取得したデータをJSONにして返却
    return c.JSON(http.StatusOK, user)
}