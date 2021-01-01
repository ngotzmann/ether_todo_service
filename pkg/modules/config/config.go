package config

import (
	"flag"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"log"
	"reflect"
	"strings"
)

var CustomCfgLocation string

func ReadConfig(cfgStruct interface{}) interface{} {
	viper.AddConfigPath(getPath())
	viper.SetConfigName(getStructType(cfgStruct))
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = mapstructure.Decode(getCfg(cfgStruct), &cfgStruct)
	if err != nil {
		log.Fatal(err)
	}
	return cfgStruct
}

func getCfg(cfgStruct interface{}) interface{} {
	var i interface{}
	err := viper.Unmarshal(&i)
	if err != nil {
		log.Fatal(err)
	}
	cfgMap := i.(map[string]interface{})
	normalizedCfgMap := cfgMap[strings.ToLower(getStructType(cfgStruct))]
	return normalizedCfgMap
}

func getPath() string {
	if CustomCfgLocation != "" {
		return CustomCfgLocation
	} else {
		return "./config/" + *getRunningEnv()
	}
}

func getStructType(i interface{}) string {
	if t := reflect.TypeOf(i); t.Kind() == reflect.Ptr {
		return strings.ToLower(t.Elem().Name())
	} else {
		return strings.ToLower(t.Name())
	}
}

func getRunningEnv() *string {
	if flag.Lookup("env") == nil {
		return flag.String("env", "local", "this is foo")
	}
	var result = "local"
	return &result
}

type Server struct {
	Address 			 string
	Port                 int
	SessionSecret        string
	SessionName          string
	DefaultLang 		 string
}

type Database struct {
	Dialect            string
	Address            string
	Port               int
	Database           string
	User               string
	Password           string
	MaxIdleConnections int
	Logging            bool
	Audit              bool
	ShouldLog 		   bool
	SSLMode 		   string
}

type Log struct {
	LogFile              string
	LogLevel             string
	LogTimestampFormat   string
}