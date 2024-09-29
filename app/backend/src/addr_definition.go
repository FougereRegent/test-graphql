package main

import (
	"test-graphql/model"
	"test-graphql/query"

	"github.com/graphql-go/graphql"
)

var addrType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Address",
	Fields: graphql.Fields{
		"Address": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if addr, ok := p.Source.(model.Address); ok {
					return addr.Address, nil
				}
				return nil, nil
			},
		},
		"Address2": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if addr, ok := p.Source.(model.Address); ok {
					return addr.Address2, nil
				}
				return nil, nil
			},
		},
		"City": &graphql.Field{
			Type:    graphql.String,
			Resolve: getCity,
		},
		"Country": &graphql.Field{
			Type:    graphql.String,
			Resolve: getCountry,
		},
	},
})

func getAddress(p graphql.ResolveParams) (interface{}, error) {
	customer, _ := p.Source.(model.Customer)
	a := query.Use(db).Address
	addr, err := a.WithContext(p.Context).Where(a.AddressID.Eq(int32(customer.AddressID))).First()
	return *addr, err
}

func getCity(p graphql.ResolveParams) (interface{}, error) {
	addr, _ := p.Source.(model.Address)
	c := query.Use(db).City

	city, err := c.WithContext(p.Context).Where(c.CityID.Eq((int32(addr.CityID)))).First()
	return city.City, err
}

func getCountry(p graphql.ResolveParams) (interface{}, error) {
	query := query.Use(db)
	addr, _ := p.Source.(model.Address)
	c := query.City
	co := query.Country

	request := co.WithContext(p.Context).Join(c, c.CountryID.EqCol(co.CountryID)).Where(c.CityID.Eq(int32(addr.CityID)))
	country, err := request.First()

	return country.Country, err
}
