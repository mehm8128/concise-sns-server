package main

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/mehm8128/concise-sns-server/model"
	"github.com/mehm8128/concise-sns-server/router"
)

func main() {
	err := model.InitDB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	router.SetRouting()
}
