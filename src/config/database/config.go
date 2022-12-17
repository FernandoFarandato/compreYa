package database

import (
	"fmt"
	"os"
)

type ConnectionData struct {
	Host     string
	Schema   string
	Username string
	Password string
	Dialect  string
}

func GetConnectionDatBase() *ConnectionData {
	connectionData := ConnectionData{}

	connectionData.Host = os.Getenv("HOST_DB")
	connectionData.Schema = "compreYa"
	connectionData.Username = os.Getenv("USERNAME_DB")
	connectionData.Password = os.Getenv("PASSWORD_DB")
	connectionData.Dialect = "mysql"

	return &connectionData
}

func GetConnectionString(cd *ConnectionData) string {
	//This is for test purposes
	//if cd.Dialect == "sqlite3" {
	//	return cd.Host
	//}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		cd.Username, cd.Password, cd.Host, cd.Schema)
}
