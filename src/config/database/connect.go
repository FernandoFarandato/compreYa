package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect() *sql.DB {
	connectionData := GetConnectionDatBase()

	client, _ := sql.Open(connectionData.Dialect, GetConnectionString(connectionData))
	if err := client.Ping(); err != nil {
		fmt.Println(err)
		return nil
	}

	mysql.SetLogger(log.Default())

	fmt.Println("Database connected successfully.")

	return client
}
