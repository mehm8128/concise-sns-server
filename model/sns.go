package model

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Post struct {
  gorm.Model
  Name string `json:"Name"`
  Content string `json:"Content"`
}

func PostContent(c echo.Context) error {
		data := new(Post)
		err := c.Bind(data)
		if err != nil {
        return echo.NewHTTPError(http.StatusBadRequest,"error")
	}
    // レコード登録
    post := Post{Name:data.Name,Content:data.Content}
		db.Create(&post)
    return c.JSON(http.StatusOK, data)
}

func GetAllPosts(c echo.Context) error {
    var posts []Post
    // userテーブルのレコードを全件取得
    db.Order("id desc").Find(&posts)
    // 取得したデータをJSONにして返却
    return echo.NewHTTPError(http.StatusOK, posts)
}