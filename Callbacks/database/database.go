package database

import (
	"fmt"
	"os"
	"strconv"
)

func GetConnectionString() string {
	mysql_user := os.Getenv("MYSQL_USER")
	mysql_password := os.Getenv("MYSQL_PASS")
	mysql_server := os.Getenv("MYSQL_SERVER")
	mysql_db := os.Getenv("MYSQL_DB")

	mysql_port, err := strconv.ParseInt(os.Getenv("MYSQL_PORT"), 0, 64)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysql_user, mysql_password, mysql_server, mysql_port, mysql_db)
}
