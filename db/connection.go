package db

import (
	"database/sql"
	"fmt"

	"github.com/WilkerAlves/api-postgresql/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	fmt.Println("abrindo banco")
	fmt.Println("base:", conf.Database)

	sc := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database,
	)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	return conn, err
}
