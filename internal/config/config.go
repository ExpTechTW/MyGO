package config

import "os"

func GetMySQLAddress() string {
	address := os.Getenv("MYSQL_ADDRESS")
	if address == "" {
		address = "127.0.0.1:3306"
	}
	return address
}
