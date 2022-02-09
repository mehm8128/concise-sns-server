package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mehm8128/concise-sns-server/model"
	"github.com/mehm8128/concise-sns-server/router"
)

type Post struct {
	gorm.Model
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

func main() {
	err := model.InitDB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	router.SetRouting()
}
