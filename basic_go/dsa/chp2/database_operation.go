package chp2

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Execute_custom() {
	var customers []Customer
	customers = GetCustomers()
	fmt.Println(customers)

	var customer Customer
	customer.CustomerName = "FHSILA"
	customer.SSN = "2323343"

	InsertCustomer(customer)

	// customer.CustomerName = "Geor Thom"
	// customer.SSN = "123412"
	// customer.CustomerId = 2
	// UpdateCustomer(customer)
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
	// Opening a driver typically will not attempt to connect to the database.
	database, err := sql.Open(databaseDriver, databaseUser+":"+databasePass+"@/"+databaseName)
	if err != nil {
		panic(err.Error())
	}
	return database
}

// TODO: complete
func CreateCustomerTable() {
	var database *sql.DB
	database = GetConnection()
	defer database.Close()

	var err error
	var res sql.Result
	query := `CREATE TABLE IF NOT EXISTS Customer(CustomerId int primary key, CustomerName varchar(255), SSN varchar(255))`
	res, err = database.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	res.LastInsertId()
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
		// copies database value to dest
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
	insert, err = database.Prepare("INSERT INTO Customer(CustomerName, SSN) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)
}

func UpdateCustomer(customer Customer) {
	var database *sql.DB

	defer database.Close()
	database = GetConnection()

	var err error
	var update *sql.Stmt
	update, err = database.Prepare("UPDATE CUSTOMER SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if err != nil {
		panic(err.Error())
	}
	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
}

func DeleteCustomer(customer Customer) {
	var database *sql.DB

	defer database.Close()
	database = GetConnection()

	var err error
	var delete *sql.Stmt
	delete, err = database.Prepare("DELETE FROM Customer WHERE Customerid=?")
	if err != nil {
		panic(err.Error())
	}
	delete.Exec(customer.CustomerId)
}
