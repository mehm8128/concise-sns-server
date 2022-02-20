package model

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			//Path:でsessionの有効な範囲を指定｡指定無しで全て有効になる｡
			//有効な時間
			MaxAge: 86400 * 7,
			//trueでjsからのアクセス拒否
			HttpOnly: true,
		}
		//ログイン
		sess.Values["auth"] = true
		//状態保存
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		fmt.Println(sess.Values["auth"])
		return c.NoContent(http.StatusOK)
	}
}

func Secret() echo.HandlerFunc {
	return func(c echo.Context) error {
		//sessionを見る
		sess, err := session.Get("session", c)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Error")
		}
		fmt.Println(sess.Values)
		//ログインしているか
		b, ok := sess.Values["auth"]
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		if b != true {
			return c.String(http.StatusUnauthorized, "401")
		} else {
			return c.String(http.StatusOK, "OK")
		}
	}
}

func Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		//ログアウト
		sess.Values["auth"] = false
		//状態を保存
		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}
		return c.NoContent(http.StatusOK)
	}
}
