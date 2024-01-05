package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	var fileConfig string
	switch os.Getenv("APP_ENV") {
	case "DEV":
		fileConfig = "env-dev.json"
	case "STG":
		fileConfig = "env-stg.json"
	case "PROD":
		fileConfig = "env-prod.json"
	default:
		fileConfig = "env.json"
	}
	logrus.Info("APP ENV : " + os.Getenv("APP_ENV"))

	viper.SetConfigName(fileConfig)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Warn(fmt.Sprintf("Error while reading config file %s", err))
	}
}

func String(key string, def string) string {
	value, ok := viper.Get(key).(string)
	if !ok {
		logrus.Warn(fmt.Sprintf("Error while reading config file %s", "Invalid type assertion - from = "+key))
		return def
	}
	return value
}

func Int(key string, def int) int {
	if value, ok := viper.Get(key).(string); ok {
		intVal, err := strconv.Atoi(value)
		if err != nil {
			logrus.Warn(fmt.Sprintf("Error while convert to int %s", key))
		}
		return intVal
	}
	logrus.Warn(fmt.Sprintf("Error while reading config file %s", "Invalid type assertion - from = "+key))
	return def

}

func Bool(key string, defaultValue bool) bool {
	if value, ok := viper.Get(key).(string); ok {
		boolVal, err := strconv.ParseBool(value)
		if err != nil {
			logrus.Warn(fmt.Sprintf("Error while convert to bool %s", key))
		}
		return boolVal
	}
	logrus.Warn(fmt.Sprintf("Error while reading config file %s", "Invalid type assertion - from = "+key))

	return defaultValue
}

func Interface(key string, def interface{}) interface{} {
	value := viper.Get(key)
	return value
}
