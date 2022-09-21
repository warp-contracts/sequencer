package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
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

	viper.AddConfigPath(getPath(path))
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	inited = true
}

func getPath(path []string) string {
	var p string
	switch len(path) {
	case 0:
		p = ".."
	case 1:
		p = path[0]
	default:
		log.Fatalf("Method don't accept more than 1 argument")
	}
	return p
}
