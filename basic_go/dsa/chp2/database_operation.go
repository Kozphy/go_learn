package chp2

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Execute_custom() {
	// create table
	// CreateCustomerTable()

	// select all
	var customers []Customer
	customers = GetCustomers()
	fmt.Println(customers)

	// insert
	// var customer Customer
	// customer.CustomerName = "FHSILA"
	// customer.SSN = "2323343"
	// InsertCustomer(customer)

	// update
	// customer.CustomerName = "Geor Thom"
	// customer.SSN = "123412"
	// customer.CustomerId = 2
	// UpdateCustomer(customer)

	// delete
	// customer.CustomerName = "Geor Thom"
	// customer.SSN = "123412"
	// customer.CustomerId = 2
	// DeleteCustomer(customer)
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
		log.Fatalf("In GetConnection sql.Open failed: %v\n", err)
	}
	return database
}

func CreateCustomerTable() {
	var database *sql.DB
	database = GetConnection()
	defer database.Close()

	var err error
	var res sql.Result
	query := `CREATE TABLE IF NOT EXISTS Customer(CustomerId int primary key auto_increment, CustomerName varchar(255), SSN varchar(255))`
	res, err = database.Exec(query)
	if err != nil {
		log.Fatalf("In CreateCustomerTable Exec failed: %v\n", err.Error())
	}
	resultid, err := res.LastInsertId()
	fmt.Println(resultid)
}

func GetCustomers() []Customer {
	var database *sql.DB
	database = GetConnection()
	defer database.Close()

	var err error
	var rows *sql.Rows

	rows, err = database.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if err != nil {
		log.Fatalf("In GetCustomers query get error: %v\n", err.Error())
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
			log.Fatalf("In GetCustomers Scan error: %v\n", err.Error())
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
	database = GetConnection()

	defer database.Close()

	var err error
	var insert *sql.Stmt
	insert, err = database.Prepare("INSERT INTO Customer(CustomerName, SSN) VALUES(?,?)")
	if err != nil {
		log.Fatalf("In insertCustomer prepare stmt failed: %v\n", err)
	}

	var result sql.Result
	result, err = insert.Exec(customer.CustomerName, customer.SSN)
	if err != nil {
		log.Fatalf("In insertCustomer Exec failed: %v\n", err)
	}

	resid, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("In insertCustomer get failed result: %v\n", err)
	}
	log.Printf("Insert sucess %v\n", resid)
}

func UpdateCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	defer database.Close()

	var err error
	var update *sql.Stmt
	update, err = database.Prepare("UPDATE Customer SET CustomerName=?, SSN=? WHERE CustomerId=?")
	if err != nil {
		log.Fatalf("In UpdateCustomer prepare err: %v\n", err.Error())
	}

	var result sql.Result
	result, err = update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	if err != nil {
		log.Fatalf("In UpdateCustomer Exec err: %v\n", err.Error())
	}

	resid, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("In UpdateCustomer res err: %v\n", err.Error())
	}
	log.Printf("In UpdateCustomer update success: %v\n", resid)
}

func DeleteCustomer(customer Customer) {
	var database *sql.DB
	database = GetConnection()

	defer database.Close()

	var err error
	var delete *sql.Stmt
	delete, err = database.Prepare("DELETE FROM Customer WHERE Customerid=?")
	if err != nil {
		log.Fatalf("In DeleteCustomer Prepare err: %v\n", err.Error())
	}

	var result sql.Result
	result, err = delete.Exec(customer.CustomerId)
	if err != nil {
		log.Fatalf("In DeleteCustomer Exec err: %v\n", err.Error())
	}
	resid, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("In DeleteCustomer get reuslt err: %v\n", err.Error())
	}
	log.Printf("In DeleteCustomer Delete success %v\n", resid)
}
