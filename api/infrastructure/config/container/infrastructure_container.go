package container

import (
	"database/sql"
	"fmt"
	"os"

	logger "github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

func GetDbConnection() *sql.DB {
	hostDatabase := os.Getenv("HOST_DATABASE")
	db, err := sql.Open("mysql", fmt.Sprintf("root:root@tcp(%s)/movies", hostDatabase))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		logger.Error(err)
	}
	return db
}
