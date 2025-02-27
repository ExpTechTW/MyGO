package main

import (
	"fmt"
	"log"

	"github.com/yourusername/mymysqlpkg/internal/config"
	"github.com/yourusername/mymysqlpkg/internal/database"
)

func main() {
	mysqlAddr := config.GetMySQLAddress()

	db, err := database.NewMySQLConn(mysqlAddr)
	if err != nil {
		log.Fatal("連接 MySQL 失敗:", err)
	}
	defer db.Close()

	query := "SELECT * FROM users;"
	err = db.SendRawSQL(query)
	if err != nil {
		log.Fatal("發送 SQL 失敗:", err)
	}

	fmt.Println("成功發送 SQL!")
}
