package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"sync"
)

const (
	PgMock = "postgres.mock"
)

var connection *sql.DB
var lock = &sync.Mutex{}

func GetPostgresConnection() *sql.DB {
	if connection == nil {
		lock.Lock()
		defer lock.Unlock()
		if connection == nil {
			host := viper.GetString("postgres.host")
			port := viper.GetInt("postgres.port")
			user := viper.GetString("postgres.user")
			password := viper.GetString("postgres.password")
			database := viper.GetString("postgres.database")
			sslmode := viper.GetString("postgres.sslmode")
			psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
				"password=%s dbname=%s sslmode=%s",
				host, port, user, password, database, sslmode)
			conn, err := sql.Open("postgres", psqlInfo)
			if err != nil {
				panic(err)
			}
			connection = conn
		}
	}
	return connection
}
