package config

import (
	"fmt"
	viper "github.com/spf13/viper"
	"reflect"
)

var CustomCfgLocation string

func ReadConfig(cfgStruct interface{}) interface{} {
	viper.New()
	/*if CustomCfgLocation != "" {
		viper.AddConfigPath(CustomCfgLocation)
	} else {*/
		viper.AddConfigPath("./config/" + *getRunningEnv() + "/")
	//}

	viper.SetConfigName("Service")
	_ = viper.ReadInConfig()

	var c Service
	_ = viper.Unmarshal(&c)

	fmt.Println(c)
	return nil

	/*rammel := getType(cfgStruct)
	viper.SetConfigName(rammel)
	viper.SetConfigType("yaml")*/
//	viper.AutomaticEnv()

/*	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	//err = viper.Unmarshal(&cfgStruct)


	//No values are loaded from
	var s *Service
	err = viper.Unmarshal(&s)
	fmt.Println(s)
	return cfgStruct*/
}

func getType(i interface{}) string {
	if t := reflect.TypeOf(i); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func getRunningEnv() *string {
	//return flag.String("env", "local", "The environment this service is deployed")
	//var blubb *string
	var result = "local"
	return &result
}
