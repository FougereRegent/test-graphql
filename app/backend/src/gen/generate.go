package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var db *gorm.DB

type connectionDb struct {
	Host     string
	User     string
	Password string
	Port     string
}

func init() {
	var err error
	c := initEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=dvdrental port=%s", c.Host, c.User, c.Password, c.Port)
	if db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		os.Exit(1)
	}
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../.",
	})

	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateModel("customer"),
		g.GenerateModel("address"),
		g.GenerateModel("country"),
		g.GenerateModel("city"),
	)
}

func initEnv() connectionDb {
	return connectionDb{
		Host:     loadDelaultEnv("APP_HOST", "localhost"),
		User:     loadDelaultEnv("APP_USER", "gorm"),
		Password: loadDelaultEnv("APP_PASSWORD", "gorm"),
		Port:     loadDelaultEnv("APP_PORT", "5432"),
	}
}

func loadDelaultEnv(envName string, defaultValue string) string {
	if r := os.Getenv(envName); r != "" {
		return r
	}
	return defaultValue
}
