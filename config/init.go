package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

var inited = false
var initLock sync.Mutex

func Init(path ...string) {
	initLock.Lock()
	defer initLock.Unlock()
	if inited {
		return
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.AddConfigPath(getPath(path))
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	validateRequiredVariables()
	initLogs()
	inited = true
}

func validateRequiredVariables() {
	for _, key := range []string{
		"postgres.password",
		"arweave.arConnectKey",
	} {
		if viper.GetString(key) == "" {
			panic(fmt.Sprintf("Key %s can't be empty", key))
		}
	}
}

func initLogs() {
	l, err := logrus.ParseLevel(viper.GetString("log.level"))
	if err != nil {
		panic(err)
	}
	logrus.SetLevel(l)

	switch viper.GetString("log.format") {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	default:
		logrus.Panic("Unsupported log format")
	}
}

func getPath(path []string) string {
	var p string
	switch len(path) {
	case 0:
		p = "./"
	case 1:
		p = path[0]
	default:
		panic("Method don't accept more than 1 argument")
	}
	return p
}
