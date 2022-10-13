package config

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var initialized = false
var initLock sync.Mutex

func Init() {
	initLock.Lock()
	defer initLock.Unlock()
	if initialized {
		return
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	configPath := getConfigPath()
	for _, env := range append([]string{""}, getEnvFiles()...) {
		viper.SetConfigFile(getConfFilename(configPath, env))
		err := viper.MergeInConfig()
		if err != nil {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}
	validateRequiredVariables()
	initLogs()
	initialized = true
}

func getConfFilename(configPath string, env string) string {
	elems := []string{"config"}
	if env != "" {
		elems = append(elems, env)
	}
	return configPath + strings.Join(elems, "_") + ".yaml"
}

func getEnvFiles() (files []string) {
	if flag.Lookup("test.v") != nil {
		files = append(files, "test")
	}
	return
}

func validateRequiredVariables() {
	for _, key := range []string{
		"arweave.walletJwk",
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
	logrus.SetReportCaller(true)

	switch viper.GetString("log.format") {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	default:
		logrus.Panic("Unsupported log format")
	}
}

func getConfigPath() string {
	p := "./"
	i := 0
out:
	for {
		files, err := os.ReadDir(p)
		if err != nil {
			logrus.Panic(err)
		}
		for _, file := range files {
			if file.Name() == "config.yaml" {
				abs, err := filepath.Abs(p)
				if err != nil {
					logrus.Panic(errors.Wrap(err, "Can't read absolute path for config"))
				}
				p = abs
				break out
			}
		}
		p = "../" + p
		i++
		if i > 10 {
			logrus.Panic("config.yaml missed? Or too deep packages structure?  Feel free to increase if you sure you need it.")
		}
	}
	return p + "/"

}
