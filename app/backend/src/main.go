package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
)

//go:generate go run ./gen/generate.go
var db *gorm.DB

type connectionDb struct {
	Host     string
	User     string
	Password string
	Port     string
}

func main() {
	if err := GormInit(); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	fieldsQuery := graphql.Fields{
		"Customer": &graphql.Field{
			Type:    customerType,
			Resolve: getCustomer,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 1,
				},
			},
		},
		"Customers": &graphql.Field{
			Type:    graphql.NewList(customerType),
			Resolve: getCustomers,
			Args: graphql.FieldConfigArgument{
				"limit": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 10,
				},
				"offset": &graphql.ArgumentConfig{
					Type:         graphql.Int,
					DefaultValue: 0,
				},
			},
		},
	}

	fieldsMutation := graphql.Fields{
		"createUser": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "User Created", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fieldsQuery,
	}

	rootMutation := graphql.ObjectConfig{
		Name:   "rootMutation",
		Fields: fieldsMutation,
	}

	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/graphql", h)
	http.ListenAndServe(":8080", nil)
}

func GormInit() error {
	var err error
	c := initEnv()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=dvdrental port=%s", c.Host, c.User, c.Password, c.Port)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
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
