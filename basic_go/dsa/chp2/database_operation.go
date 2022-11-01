package chp2

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Execute_custom() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println("Customers", customers)
}

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

func GetConnection() (database *sql.DB) {
	databaseDriver := "mysql"
	databaseUser := "root"
	databasePass := "root"
	databaseName := "crm"
	database, err := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	return database
}

func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()
	defer database.Close()

	var err error
	var rows *sql.Rows

	rows, err = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if err != nil {
		panic(err.Error())
	}

	var customer Customer = Customer{}

	var customers []Customer
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		err = rows.Scan(&customerId, &customerName, &ssn)
		if err != nil {
			panic(err.Error())
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}
	return customers
}

func InsertCustomer(customer Customer) {
	var database *sql.DB
	defer database.Close()

	database = GetConnection()

	var err error
	var insert *sql.Stmt
	insert, err = database.Prepare("INSERT INTO CUSTOMER(CustomerName, SSN) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
}
