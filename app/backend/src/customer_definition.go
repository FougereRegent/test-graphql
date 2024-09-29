package main

import (
	"errors"
	"test-graphql/model"
	"test-graphql/query"

	"github.com/graphql-go/graphql"
)

var customerType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Customer",
	Fields: graphql.Fields{
		"FirstName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if firstName, ok := p.Source.(model.Customer); ok {
					return firstName.FirstName, nil
				} else {
					return nil, nil
				}
			},
		},
		"LastName": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if lastName, ok := p.Source.(model.Customer); ok {
					return lastName.LastName, nil
				} else {
					return nil, nil
				}
			},
		},
		"Email": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if email, ok := p.Source.(model.Customer); ok {
					return email.Email, nil
				} else {
					return nil, nil
				}
			},
		},
		"Address": &graphql.Field{
			Type:    addrType,
			Resolve: getAddress,
		},
	},
})

func getCustomer(p graphql.ResolveParams) (interface{}, error) {
	id, ok := p.Args["id"].(int)
	if !ok {
		return nil, errors.New("Bad format")
	}
	u := query.Use(db).Customer
	user, err := u.WithContext(p.Context).Where(u.CustomerID.Eq(int32(id))).First()
	return *user, err
}

func getCustomers(p graphql.ResolveParams) (interface{}, error) {
	var limit, offset int
	var ok bool

	if limit, ok = p.Args["limit"].(int); !ok {
		return nil, errors.New("Limit args doesn't exist")
	}
	if offset, ok = p.Args["offset"].(int); !ok {
		return nil, errors.New("Limit args doesn't exist")
	}
	u := query.Use(db).Customer

	user, err := u.WithContext(p.Context).Limit(limit).Offset(offset).Find()
	if err != nil {
		return nil, err
	}

	return user, nil
}
