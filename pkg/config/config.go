package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var tokens = []string{"SEPTA_URL", "SEPTA_TOKEN", "SEPTA_BUCKET", "SEPTA_ORG"}

func SetEnv() {
	for _, j := range tokens {
		IfEnvNot(j)
	}
}

func IfEnvNot(j string)  {
	r := os.Getenv(j)
	if r == "" {
		os.Setenv(j, viperEnvVariable(j))
	}
}



func viperEnvVariable(key string) string {

	viper.SetConfigName(".septa")
	viper.AddConfigPath("$HOME") // call multiple times to add many search paths
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/septa")

	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Printf("Error while reading config file %s", err)
		return ""
	}

	// viper.Get() returns an empty interface{}
	// to get the underlying type of the key,
	// we have to do the type assertion, we know the underlying value is string
	// if we type assert to other type it will throw an error
	value, ok := viper.Get(key).(string)

	// If the type is a string then ok will be true
	// ok will make sure the program not break
	if !ok {
		log.Printf("Invalid type assertion")
		return ""
	}

	return value
}
