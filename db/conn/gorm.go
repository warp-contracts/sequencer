package conn

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	gormConnectOnce sync.Once
	gormConn        *gorm.DB
)

func GetConnection() *gorm.DB {
	gormConnectOnce.Do(func() {
		initGormConnect()
	})
	return gormConn
}

func initGormConnect() {
	host := viper.GetString("postgres.host")
	port := viper.GetInt("postgres.port")
	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.password")
	database := viper.GetString("postgres.database")
	sslmode := viper.GetString("postgres.sslmode")
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, database, sslmode)

	var err error
	gormConn, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		logrus.Panic(err)
	}
}
