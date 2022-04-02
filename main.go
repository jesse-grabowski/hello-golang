package main

//go:generate pkger

import (
	"com.jessegrabowski/go-webapp/config"
	"com.jessegrabowski/go-webapp/database"
	"com.jessegrabowski/go-webapp/server"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jessevdk/go-flags"
)

var validate *validator.Validate

var opts struct {
	Port            string   `short:"p" long:"port" description:"Application port" required:"true"`
	Profiles        []string `short:"P" long:"profile" description:"Active configuration profile(s)" required:"false"`
	ConfigLocations []string `short:"l" long:"config-location" description:"Directory to load configuration from" default:"."`
}

func main() {
	validate = validator.New()

	_, err := flags.Parse(&opts)
	if err != nil {
		panic(fmt.Errorf("failed to parse args: %w \n", err))
	}

	config.Init(validate, opts.ConfigLocations, opts.Profiles)
	database.Init()
	server.Start(opts.Port)
}
