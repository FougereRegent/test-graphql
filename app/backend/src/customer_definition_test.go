package main

import (
	"fmt"
	"test-graphql/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/graphql-go/graphql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDb() (func(), sqlmock.Sqlmock) {
	conn, mock, _ := sqlmock.New()

	connection, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}))

	db = connection

	def := func() {
		conn.Close()
	}

	return def, mock
}

func prepareData(mock sqlmock.Sqlmock, id int) {
	rows := sqlmock.NewRows([]string{"customer_id", "first_name", "last_name", "email"}).
		AddRow(id, "damien", "venant-valéry", "damien.venant@outlook.com")
	mock.ExpectQuery("SELECT").
		WillReturnRows(rows)
}

func prepareDatas(mock sqlmock.Sqlmock) {
	rows := sqlmock.NewRows([]string{"customer_id", "first_name", "last_name", "email"}).
		AddRow(1, "damien1", "venant-valéry1", "damien.venant@outlook.com1").
		AddRow(2, "damien2", "venant-valéry2", "damien.venant@outlook.com2").
		AddRow(3, "damien3", "venant-valéry3", "damien.venant@outlook.com3").
		AddRow(4, "damien4", "venant-valéry4", "damien.venant@outlook.com4").
		AddRow(5, "damien5", "venant-valéry5", "damien.venant@outlook.com5").
		AddRow(6, "damien6", "venant-valéry6", "damien.venant@outlook.com6").
		AddRow(7, "damien7", "venant-valéry7", "damien.venant@outlook.com7").
		AddRow(8, "damien8", "venant-valéry8", "damien.venant@outlook.com8").
		AddRow(9, "damien9", "venant-valéry9", "damien.venant@outlook.com9").
		AddRow(10, "damien10", "venant-valéry10", "damien.venant@outlook.com10")
	mock.ExpectQuery("SELECT").
		WillReturnRows(rows)
}

func TestGetCustomer(t *testing.T) {
	for i := 0; i < 100; i++ {
		setup, mock := setupDb()
		defer setup()

		/*Init Mock*/
		prepareData(mock, i)
		args := make(map[string]interface{})
		args["id"] = i

		res, _ := getCustomer(graphql.ResolveParams{
			Args: args,
		})
		if res == nil {
			t.Fatalf("The result shouldn't be null")
		}

		if user, ok := res.(model.Customer); !ok {
			t.Fatalf("The result should be a Customer type")
		} else if user.CustomerID != int32(i) {
			t.Fatalf("Customer ID : %d\nId : %d", user.CustomerID, i)
		}
	}
}

func TestGetCustomers(t *testing.T) {
	setup, mock := setupDb()
	defer setup()

	/*Init Data*/
	prepareDatas(mock)

	args := make(map[string]interface{})
	args["limit"] = 1
	args["offset"] = 2

	res, err := getCustomers(graphql.ResolveParams{
		Args: args,
	})
	if err != nil {
		t.Fatalf(err.Error())
	}

	if users, ok := res.([]*model.Customer); !ok {
		fmt.Println(res)
		t.FailNow()
	} else if len(users) <= 0 {
		t.FailNow()
	}
}
