package database

import (
	"com.jessegrabowski/go-webapp/business/sampling"
	"com.jessegrabowski/go-webapp/config"
	"fmt"
	"strings"
)

var databaseConfig struct {
	ConnectionString string `mapstructure:"url" validate:"required"`
	MigrateEnabled   bool   `mapstructure:"migrate"`
}

type dao interface {
	init(connectionString string) error
	migrate(connectionString string) error
	sampling.Dao
}

var Dao dao

func Init() {
	config.BindPrefix("db", &databaseConfig)

	switch databaseConfig.ConnectionString[:strings.Index(databaseConfig.ConnectionString, ":")] {
	case "postgres":
		Dao = postgres{}
	default:
		panic("no valid database definition found \n")
	}

	err := Dao.init(databaseConfig.ConnectionString)
	if err != nil {
		panic(fmt.Errorf("failed to initialize dao: %w \n", err))
	}

	if databaseConfig.MigrateEnabled {
		err = Dao.migrate(databaseConfig.ConnectionString)
		if err != nil {
			panic(fmt.Errorf("failed to migrate dao: %w \n", err))
		}
	}
}
