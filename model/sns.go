package model

import (
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Name     string `json:"name"`
	Content  string `json:"content"`
	Password string `json:"password"`
}
type DeleteRequest struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

func GetAllPosts(c echo.Context) error {
	var posts []Post
	// postsテーブルのレコードを全件取得
	db.Order("id desc").Find(&posts)
	// 取得したデータをJSONにして返却
	return echo.NewHTTPError(http.StatusOK, posts)
}

func PostContent(c echo.Context) error {
	data := new(Post)
	err := c.Bind(data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error")
	}
	// レコード登録
	post := Post{Name: data.Name, Content: data.Content, Password: data.Password}
	db.Create(&post)
	return c.JSON(http.StatusOK, data)
}

func DeletePost(c echo.Context) error {
	var post Post
	var posts []Post

	data := new(Post)
	err := c.Bind(data)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error")
	}
	// postsテーブルのIDがdata.IDのものを取得
	db.First(&post, data.ID)
	// passwordが一致するか確認
	if post.Password != data.Password {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid password")
	}
	db.Delete(&posts, data.ID)
	// postsテーブルのレコードを全件取得
	db.Order("id desc").Find(&posts)
	// 削除後のリストを返却
	return echo.NewHTTPError(http.StatusOK, posts)
}
