package model

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// DBのインスタンスをグローバル変数に格納
var (
	db *gorm.DB
)

func InitDB()error{
	//環境変数読み込み
    godotenv.Load(".env")
    // DB接続処理
    var err error
    db, err = gorm.Open("postgres", os.Getenv("DATABASE_URL") )
    if err != nil {
        panic("failed to connect database")
    }
		db.AutoMigrate(&Post{})
		return err
}