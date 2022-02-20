package model

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type PostRequest struct {
	Name     string `json:"name"`
	Content  string `json:"content"`
	Password string `json:"password"`
}
type Post struct {
	gorm.Model
	Name           string `json:"name"`
	Content        string `json:"content"`
	HashedPassword string `json:"password"`
}
type DeleteRequest struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

func GetAllPosts(c echo.Context) error {
	var posts []*Post
	// postsテーブルのレコードを全件取得
	db.Order("id desc").Find(&posts)
	// 取得したデータをJSONにして返却
	return echo.NewHTTPError(http.StatusOK, posts)
}

func PostContent(c echo.Context) error {
	data := new(PostRequest)
	err := c.Bind(data)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error")
	}
	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "hash error")
	}
	// レコード登録
	post := Post{Name: data.Name, Content: data.Content, HashedPassword: string(hashedPassword)}
	db.Create(&post)
	return c.JSON(http.StatusOK, &data)
}

func DeletePost(c echo.Context) error {
	var post Post
	var posts []*Post

	data := new(DeleteRequest)
	err := c.Bind(data)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error")
	}
	// postsテーブルのIDがdata.IDのものを取得
	db.First(&post, data.ID)
	// passwordが一致するか確認
	err = bcrypt.CompareHashAndPassword([]byte(post.HashedPassword), []byte(data.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid password")
	}
	db.Delete(&posts, data.ID)
	// postsテーブルのレコードを全件取得
	db.Order("id desc").Find(&posts)
	// 削除後のリストを返却
	return echo.NewHTTPError(http.StatusOK, posts)
}
