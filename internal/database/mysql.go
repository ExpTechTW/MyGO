package database

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type MySQLConn struct {
	conn net.Conn
}

func NewMySQLConn(address string) (*MySQLConn, error) {
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)
	if err != nil {
		return nil, fmt.Errorf("無法連接到 MySQL 伺服器: %v", err)
	}
	log.Println("成功連接到 MySQL 伺服器")
	return &MySQLConn{conn: conn}, nil
}

func (db *MySQLConn) SendRawSQL(query string) error {
	if db.conn == nil {
		return errors.New("資料庫未連線")
	}
	_, err := db.conn.Write([]byte(query))
	return err
}

func (db *MySQLConn) Close() {
	if db.conn != nil {
		db.conn.Close()
	}
}
