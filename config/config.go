package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/markbates/pkger"
	"github.com/spf13/viper"
)

type configReader func() error

var validate *validator.Validate

func Init(v *validator.Validate, configLocations []string, profiles []string) {
	validate = v
	viper.SetConfigType("yaml")

	for _, location := range configLocations {
		viper.AddConfigPath(location)
	}

	tryLoadConfig(func() error {
		file, err := pkger.Open("/application.yml")
		if err != nil {
			return err
		}
		return viper.ReadConfig(file)
	})

	for _, profile := range profiles {
		viper.SetConfigName(fmt.Sprintf("application-%s.yml", profile))
		tryLoadConfig(viper.MergeInConfig)
	}
}

func tryLoadConfig(handler configReader) {
	err := handler()
	if err != nil {
		panic(fmt.Errorf("failed to load config file: %w \n", err))
	}
}

func BindPrefix(prefix string, rawValue interface{}) {
	err := viper.UnmarshalKey(prefix, rawValue)
	if err != nil {
		panic(fmt.Errorf("failed to unmarshall configuration: %w \n", err))
	}
	err = validate.Struct(rawValue)
	if err != nil {
		panic(fmt.Errorf("failed to validate configuration: %w \n", err))
	}
}
