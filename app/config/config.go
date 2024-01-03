package config

import (
	"bytes"
	"os"
	"reflect"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/spf13/viper"
)

const (
	envPrefix = "WARP_SEQUENCER_"
)

type Config struct {
	RateLimiter RateLimiter
}

func setDefaults() {
	setRateLimiterDefaults()
}

func IsIndex(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func BindEnv(path []string, val reflect.Value) {
	if val.Kind() == reflect.Slice {
		// Slice of base types
		key := strings.ToLower(strings.Join(path, "."))
		env := envPrefix + strcase.ToScreamingSnake(strings.Join(path, "_"))
		err := viper.BindEnv(key, env)
		if err != nil {
			panic(err)
		}
	} else if val.Kind() != reflect.Struct {
		// Base types
		// key := strings.ToLower(strings.Join(path, "."))
		key := path[0]
		for _, p := range path[1:] {
			if IsIndex(p) {
				key += "[" + p + "]"
				// key += "." + p
			} else {
				key += "." + p
			}
		}

		env := envPrefix + strcase.ToScreamingSnake(strings.Join(path, "_"))
		err := viper.BindEnv(key, env)
		if err != nil {
			panic(err)
		}
	} else {
		// Iterates over struct fields
		for i := 0; i < val.NumField(); i++ {
			newPath := make([]string, len(path))
			copy(newPath, path)
			newPath = append(newPath, val.Type().Field(i).Name)
			BindEnv(newPath, val.Field(i))
		}
	}
}

// Load configuration from file and env
func Load(filename string) (config *Config, err error) {
	viper.SetConfigType("json")

	setDefaults()

	// Visits every field and registers upper snake case ENV name for it
	// Works with embedded structs
	BindEnv([]string{}, reflect.ValueOf(Config{}))

	// Empty filename means we use default values
	if filename != "" {
		var content []byte
		/* #nosec */
		content, err = os.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		err = viper.ReadConfig(bytes.NewBuffer(content))
		if err != nil {
			return nil, err
		}
	}

	config = new(Config)
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return
}
