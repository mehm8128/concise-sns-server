package main

import (
	"net/http"
	"os"

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
    e.GET("/users", getAllUsers)
    e.POST("/create", createUser)
    // サーバー起動
    e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
    //e.Logger.Fatal(e.Start(":8080"))
}

func createUser(c echo.Context) error {
		data := new(User)
		err := c.Bind(data)
		if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest,"error")
	}
    // レコード登録
    var user = User{Name:data.Name}
		db.Create(&user)
    return c.JSON(http.StatusOK, data)
}

func getAllUsers(c echo.Context) error {
    var users []*User
    // userテーブルのレコードを全件取得
    db.Find(&users)
    // 取得したデータをJSONにして返却
    return echo.NewHTTPError(http.StatusOK, users)
}