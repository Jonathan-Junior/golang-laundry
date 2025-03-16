package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

//Database Connection
const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password ="Postgrejo123"
	dbname = "enigma_laundry"
)

//Struct Model
type Customer struct{
	customer_id int
	name string
	phone string
	address string
	created_at time.Time
	updated_at time.Time
}

type Order struct{
	order_id int
	customer_id int
	order_date time.Time
	completion_date sql.NullTime
	received_by string
	created_at time.Time
	updated_at time.Time
}

type OrderDetail struct{
	order_id int
	service_id int
	qty int
}

type Service struct {
	service_id   int
	service_name string
	unit         string
	price        int
	created_at   time.Time
	updated_at   time.Time
}

var db *sql.DB
var scanner *bufio.Scanner



func main() {
scanner = bufio.NewScanner(os.Stdin)

	//Connect to Database
	psqlinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Test Connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfuly Connected")

//Menu
// Main menu
for {
	clearScreen()
	fmt.Println("====== ENIGMA LAUNDRY MANAGEMENT SYSTEM ======")
	fmt.Println("1. Customer")
	fmt.Println("2. Service")
	fmt.Println("3. Order")
	fmt.Println("4. Exit")
	fmt.Print("\nChoose menu: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		customerMenu()
	case 2:
		serviceMenu()
	case 3:
		orderMenu()
	case 4:
		fmt.Println("Thank you for using Enigma Laundry Management System!")
		return
	default:
		fmt.Println("Invalid choice! Press Enter to continue...")
		fmt.Scanln()
	}
}
}

// Customer Menu
func customerMenu() {
for {
	clearScreen()
	fmt.Println("====== CUSTOMER MENU ======")
	fmt.Println("1. Create Customer")
	fmt.Println("2. View Of List Customer")
	fmt.Println("3. View Details Customer By ID")
	fmt.Println("4. Update Customer")
	fmt.Println("5. Delete Customer")
	fmt.Println("6. Back to Main Menu")
	fmt.Print("\nChoose menu: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		createCustomer()
	case 2:
		viewListCustomer()
	case 3:
		viewCustomerByID()
	case 4:
		updateCustomer()
	case 5:
		deleteCustomer()
	case 6:
		return
	default:
		fmt.Println("Invalid choice! Press Enter to continue...")
		fmt.Scanln()
	}
}
}

// Service Menu
func serviceMenu() {
for {
	clearScreen()
	fmt.Println("====== SERVICE MENU ======")
	fmt.Println("1. Create Service")
	fmt.Println("2. View Of List Service")
	fmt.Println("3. View Details Service By ID")
	fmt.Println("4. Update Service")
	fmt.Println("5. Delete Service")
	fmt.Println("6. Back to Main Menu")
	fmt.Print("\nChoose menu: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		createService()
	case 2:
		viewListService()
	case 3:
		viewServiceByID()
	case 4:
		updateService()
	case 5:
		deleteService()
	case 6:
		return
	default:
		fmt.Println("Invalid choice! Press Enter to continue...")
		fmt.Scanln()
	}
}
}

// Order Menu
func orderMenu() {
for {
	clearScreen()
	fmt.Println("====== ORDER MENU ======")
	fmt.Println("1. Create Order")
	fmt.Println("2. Complete Order")
	fmt.Println("3. View Of List Order")
	fmt.Println("4. View Order Details By ID")
	fmt.Println("5. Back to Main Menu")
	fmt.Print("\nChoose menu: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		createOrder()
	case 2:
		completeOrder()
	case 3:
		viewListOrder()
	case 4:
		viewOrderByID()
	case 5:
		return
	default:
		fmt.Println("Invalid choice! Press Enter to continue...")
		fmt.Scanln()
	}
}
}

// Helper function to read input line
func readLine() string {
scanner.Scan()
return strings.TrimSpace(scanner.Text())
}

// Helper function to clear screen
func clearScreen() {
fmt.Print("\033[H\033[2J")
}

// Customer functions
func createCustomer() {
clearScreen()
fmt.Println("====== CREATE CUSTOMER ======")

var customer Customer

fmt.Print("Enter Customer ID: ")
fmt.Scanln(&customer.customer_id)

// Check if customer ID already exists
var exists bool
err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1)", customer.customer_id).Scan(&exists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if exists {
	fmt.Println("Customer ID already exists. Please enter a different ID.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Print("Enter Name: ")
fmt.Scanln() // Clear buffer
customer.name = readLine()

fmt.Print("Enter Phone: ")
customer.phone = readLine()

fmt.Print("Enter Address: ")
customer.address = readLine()

// Insert new customer
_, err = db.Exec("INSERT INTO customer (customer_id, name, phone, address) VALUES ($1, $2, $3, $4)",
	customer.customer_id, customer.name, customer.phone, customer.address)

if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred while creating customer. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Customer created successfully!")
fmt.Println("Press Enter to continue...")
fmt.Scanln()
}

func viewListCustomer() {
clearScreen()
fmt.Println("====== LIST OF CUSTOMERS ======")

rows, err := db.Query("SELECT customer_id, name, phone FROM customer ORDER BY customer_id")
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}
defer rows.Close()

fmt.Printf("%-10s %-30s %-15s\n", "ID", "Name", "Phone")
fmt.Println(strings.Repeat("-", 60))

for rows.Next() {
	var id int
	var name, phone string
	err = rows.Scan(&id, &name, &phone)
	if err != nil {
		log.Printf("Row scan error: %v", err)
		continue
	}
	fmt.Printf("%-10d %-30s %-15s\n", id, name, phone)
}

fmt.Println("\nPress Enter to continue...")
fmt.Scanln()
}

func viewCustomerByID() {
clearScreen()
fmt.Println("====== VIEW CUSTOMER DETAILS ======")

var customer_id int
fmt.Print("Enter Customer ID: ")
fmt.Scanln(&customer_id)

var customer Customer
err := db.QueryRow("SELECT customer_id, name, phone, address, created_at, updated_at FROM customer WHERE customer_id = $1",
	customer_id).Scan(&customer.customer_id, &customer.name, &customer.phone, &customer.address, &customer.created_at, &customer.updated_at)

if err == sql.ErrNoRows {
	fmt.Println("Customer not found.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
} else if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Customer Details:")
fmt.Printf("ID          : %d\n", customer.customer_id)
fmt.Printf("Name        : %s\n", customer.name)
fmt.Printf("Phone       : %s\n", customer.phone)
fmt.Printf("Address     : %s\n", customer.address)
fmt.Printf("Created At  : %s\n", customer.created_at.Format("2006-01-02 15:04:05"))
fmt.Printf("Updated At  : %s\n", customer.updated_at.Format("2006-01-02 15:04:05"))

fmt.Println("\nPress Enter to continue...")
fmt.Scanln()
}

func updateCustomer() {
clearScreen()
fmt.Println("====== UPDATE CUSTOMER ======")

var customer_id int
fmt.Print("Enter Customer ID: ")
fmt.Scanln(&customer_id)

// Check if customer exists
var exists bool
err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1)", customer_id).Scan(&exists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if !exists {
	fmt.Println("Customer not found.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

var customer Customer
customer.customer_id = customer_id

fmt.Print("Enter new Name: ")
fmt.Scanln() // Clear buffer
customer.name = readLine()

fmt.Print("Enter new Phone: ")
customer.phone = readLine()

fmt.Print("Enter new Address: ")
customer.address = readLine()

// Update customer
_, err = db.Exec("UPDATE customer SET name = $1, phone = $2, address = $3, updated_at = CURRENT_TIMESTAMP WHERE customer_id = $4",
	customer.name, customer.phone, customer.address, customer.customer_id)

if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred while updating customer. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Customer updated successfully!")
fmt.Println("Press Enter to continue...")
fmt.Scanln()
}

func deleteCustomer() {
clearScreen()
fmt.Println("====== DELETE CUSTOMER ======")

var customer_id int
fmt.Print("Enter Customer ID: ")
fmt.Scanln(&customer_id)

// Check if customer exists
var exists bool
err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1)", customer_id).Scan(&exists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if !exists {
	fmt.Println("Customer ID not found. Please enter a different ID.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

// Check if customer is being used in orders
var orderExists bool
err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM \"order\" WHERE customer_id = $1)", customer_id).Scan(&orderExists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if orderExists {
	fmt.Println("Customer ID is being used in orders. Please delete the order first.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

// Delete customer
_, err = db.Exec("DELETE FROM customer WHERE customer_id = $1", customer_id)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred while deleting customer. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Customer deleted successfully!")
fmt.Println("Press Enter to continue...")
fmt.Scanln()
}

// Service functions
func createService() {
clearScreen()
fmt.Println("====== CREATE SERVICE ======")

var service Service

fmt.Print("Enter Service ID: ")
fmt.Scanln(&service.service_id)

// Check if service ID already exists
var exists bool
err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM service WHERE service_id = $1)", service.service_id).Scan(&exists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if exists {
	fmt.Println("Service ID already exists. Please enter a different ID.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Print("Enter Service Name: ")
fmt.Scanln() // Clear buffer
service.service_name = readLine()

fmt.Print("Enter Unit: ")
service.unit = readLine()

fmt.Print("Enter Price: ")
fmt.Scanln(&service.price)

// Insert new service
_, err = db.Exec("INSERT INTO service (service_id, service_name, unit, price) VALUES ($1, $2, $3, $4)",
	service.service_id, service.service_name, service.unit, service.price)

if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred while creating service. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Service created successfully!")
fmt.Println("Press Enter to continue...")
fmt.Scanln()
}

func viewListService() {
clearScreen()
fmt.Println("====== LIST OF SERVICES ======")

rows, err := db.Query("SELECT service_id, service_name, unit, price FROM service ORDER BY service_id")
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}
defer rows.Close()

fmt.Printf("%-10s %-30s %-15s %-10s\n", "ID", "Service Name", "Unit", "Price")
fmt.Println(strings.Repeat("-", 70))

for rows.Next() {
	var id, price int
	var name, unit string
	err = rows.Scan(&id, &name, &unit, &price)
	if err != nil {
		log.Printf("Row scan error: %v", err)
		continue
	}
	fmt.Printf("%-10d %-30s %-15s %-10d\n", id, name, unit, price)
}

fmt.Println("\nPress Enter to continue...")
fmt.Scanln()
}

func viewServiceByID() {
clearScreen()
fmt.Println("====== VIEW SERVICE DETAILS ======")

var service_id int
fmt.Print("Enter Service ID: ")
fmt.Scanln(&service_id)

var service Service
err := db.QueryRow("SELECT service_id, service_name, unit, price, created_at, updated_at FROM service WHERE service_id = $1",
	service_id).Scan(&service.service_id, &service.service_name, &service.unit, &service.price, &service.created_at, &service.updated_at)

if err == sql.ErrNoRows {
	fmt.Println("Service not found.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
} else if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Service Details:")
fmt.Printf("ID          : %d\n", service.service_id)
fmt.Printf("Name        : %s\n", service.service_name)
fmt.Printf("Unit        : %s\n", service.unit)
fmt.Printf("Price       : %d\n", service.price)
fmt.Printf("Created At  : %s\n", service.created_at.Format("2006-01-02 15:04:05"))
fmt.Printf("Updated At  : %s\n", service.updated_at.Format("2006-01-02 15:04:05"))

fmt.Println("\nPress Enter to continue...")
fmt.Scanln()
}

func updateService() {
clearScreen()
fmt.Println("====== UPDATE SERVICE ======")

var service_id int
fmt.Print("Enter Service ID: ")
fmt.Scanln(&service_id)

// Check if service exists
var exists bool
err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM service WHERE service_id = $1)", service_id).Scan(&exists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if !exists {
	fmt.Println("Service not found.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

var service Service
service.service_id = service_id

fmt.Print("Enter new Service Name: ")
fmt.Scanln() // Clear buffer
service.service_name = readLine()

fmt.Print("Enter new Unit: ")
service.unit = readLine()

fmt.Print("Enter new Price: ")
fmt.Scanln(&service.price)

// Update service
_, err = db.Exec("UPDATE service SET service_name = $1, unit = $2, price = $3, updated_at = CURRENT_TIMESTAMP WHERE service_id = $4",
	service.service_name, service.unit, service.price, service.service_id)

if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred while updating service. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Service updated successfully!")
fmt.Println("Press Enter to continue...")
fmt.Scanln()
}

func deleteService() {
clearScreen()
fmt.Println("====== DELETE SERVICE ======")

var service_id int
fmt.Print("Enter Service ID: ")
fmt.Scanln(&service_id)

// Check if service exists
var exists bool
err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM service WHERE service_id = $1)", service_id).Scan(&exists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if !exists {
	fmt.Println("Service ID not found. Please enter a different ID.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

// Check if service is being used in orders
var orderDetailExists bool
err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM order_detail WHERE service_id = $1)", service_id).Scan(&orderDetailExists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if orderDetailExists {
	fmt.Println("Service ID is being used in orders. Please delete the order first.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

// Delete service
_, err = db.Exec("DELETE FROM service WHERE service_id = $1", service_id)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred while deleting service. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Service deleted successfully!")
fmt.Println("Press Enter to continue...")
fmt.Scanln()
}

// Order functions
func createOrder() {
clearScreen()
fmt.Println("====== CREATE ORDER ======")

var order Order

fmt.Print("Enter Order ID: ")
fmt.Scanln(&order.order_id)

// Check if order ID already exists
var orderExists bool
err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM \"order\" WHERE order_id = $1)", order.order_id).Scan(&orderExists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if orderExists {
	fmt.Println("Order ID already exists. Please enter a different ID.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Print("Enter Customer ID: ")
fmt.Scanln(&order.customer_id)

// Check if customer exists
var customerExists bool
err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM customer WHERE customer_id = $1)", order.customer_id).Scan(&customerExists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if !customerExists {
	fmt.Println("Customer not found.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

// Set order date to current date
order.order_date = time.Now()

fmt.Print("Enter Received By: ")
fmt.Scanln() // Clear buffer
order.received_by = readLine()

// Start a transaction
tx, err := db.Begin()
if err != nil {
	log.Printf("Transaction error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

// Insert new order
_, err = tx.Exec("INSERT INTO \"order\" (order_id, customer_id, order_date, received_by) VALUES ($1, $2, $3, $4)",
	order.order_id, order.customer_id, order.order_date, order.received_by)

if err != nil {
	tx.Rollback()
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred while creating order. Press Enter to continue...")
	fmt.Scanln()
	return
}

// Add order details
for {
	fmt.Println("\nAdd Order Detail (Enter 0 for Service ID to finish):")

	var order_detail OrderDetail
	order_detail.order_id = order.order_id

	fmt.Print("Enter Service ID (0 to finish): ")
	fmt.Scanln(&order_detail.service_id)

	if order_detail.service_id == 0 {
		break
	}

	// Check if service exists
	var serviceExists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM service WHERE service_id = $1)", order_detail.service_id).Scan(&serviceExists)
	if err != nil {
		log.Printf("Database error: %v", err)
		continue
	}

	if !serviceExists {
		fmt.Println("Service not found. Try again.")
		continue
	}

	fmt.Print("Enter Quantity: ")
	fmt.Scanln(&order_detail.qty)

	// Insert order detail
	_, err = tx.Exec("INSERT INTO order_detail (order_id, service_id, qty) VALUES ($1, $2, $3)",
		order_detail.order_id, order_detail.service_id, order_detail.qty)

	if err != nil {
		log.Printf("Database error: %v", err)
		fmt.Println("Error adding order detail. Try again.")
		continue
	}

	fmt.Println("Service added to order successfully!")
}

// Commit the transaction
err = tx.Commit()
if err != nil {
	log.Printf("Transaction commit error: %v", err)
	fmt.Println("An error occurred while finalizing order. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Order created successfully!")
fmt.Println("Press Enter to continue...")
fmt.Scanln()
}

func completeOrder() {
clearScreen()
fmt.Println("====== COMPLETE ORDER ======")

var order_id int
fmt.Print("Enter Order ID: ")
fmt.Scanln(&order_id)

// Check if order exists
var exists bool
err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM \"order\" WHERE order_id = $1)", order_id).Scan(&exists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if !exists {
	fmt.Println("Order not found.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Print("Enter Completion Date (YYYY-MM-DD): ")
var date_str string
fmt.Scanln(&date_str)

completion_date, err := time.Parse("2006-01-02", date_str)
if err != nil {
	fmt.Println("Invalid date format. Please use YYYY-MM-DD.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

// Update order
_, err = db.Exec("UPDATE \"order\" SET completion_date = $1, updated_at = CURRENT_TIMESTAMP WHERE order_id = $2",
	completion_date, order_id)

if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred while completing order. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Order completed successfully!")
fmt.Println("Press Enter to continue...")
fmt.Scanln()
}

func viewListOrder() {
clearScreen()
fmt.Println("====== LIST OF ORDERS ======")

rows, err := db.Query(`
	SELECT o.order_id, c.name, o.order_date, o.completion_date, o.received_by
	FROM "order" o
	JOIN customer c ON o.customer_id = c.customer_id
	ORDER BY o.order_id
`)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}
defer rows.Close()

fmt.Printf("%-10s %-25s %-15s %-15s %-20s\n", "ID", "Customer", "Order Date", "Completion Date", "Received By")
fmt.Println(strings.Repeat("-", 90))

for rows.Next() {
	var id int
	var customer, received_by string
	var order_date time.Time
	var completion_date sql.NullTime

	err = rows.Scan(&id, &customer, &order_date, &completion_date, &received_by)
	if err != nil {
		log.Printf("Row scan error: %v", err)
		continue
	}

	completion_str := "Not completed"
	if completion_date.Valid {
		completion_str = completion_date.Time.Format("2006-01-02")
	}

	fmt.Printf("%-10d %-25s %-15s %-15s %-20s\n",
		id, customer, order_date.Format("2006-01-02"), completion_str, received_by)
}

fmt.Println("\nPress Enter to continue...")
fmt.Scanln()
}

func viewOrderByID() {
clearScreen()
fmt.Println("====== VIEW ORDER DETAILS ======")

var order_id int
fmt.Print("Enter Order ID: ")
fmt.Scanln(&order_id)

// Check if order exists
var exists bool
err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM \"order\" WHERE order_id = $1)", order_id).Scan(&exists)
if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred. Press Enter to continue...")
	fmt.Scanln()
	return
}

if !exists {
	fmt.Println("Order not found.")
	fmt.Println("Press Enter to continue...")
	fmt.Scanln()
	return
}

// Get order header information
var order Order
var customer_name string

err = db.QueryRow(`
	SELECT o.order_id, o.customer_id, c.name, o.order_date, o.completion_date, o.received_by, o.created_at, o.updated_at
	FROM "order" o
	JOIN customer c ON o.customer_id = c.customer_id
	WHERE o.order_id = $1
`, order_id).Scan(
	&order.order_id, &order.customer_id, &customer_name, &order.order_date,
	&order.completion_date, &order.received_by, &order.created_at, &order.updated_at)

if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred while retrieving order. Press Enter to continue...")
	fmt.Scanln()
	return
}

fmt.Println("Order Details:")
fmt.Printf("Order ID        : %d\n", order.order_id)
fmt.Printf("Customer        : %s (ID: %d)\n", customer_name, order.customer_id)
fmt.Printf("Order Date      : %s\n", order.order_date.Format("2006-01-02"))

if order.completion_date.Valid {
	fmt.Printf("Completion Date : %s\n", order.completion_date.Time.Format("2006-01-02"))
} else {
	fmt.Println("Completion Date : Not completed")
}

fmt.Printf("Received By     : %s\n", order.received_by)
fmt.Printf("Created At      : %s\n", order.created_at.Format("2006-01-02 15:04:05"))
fmt.Printf("Updated At      : %s\n", order.updated_at.Format("2006-01-02 15:04:05"))

// Get order details
rows, err := db.Query(`
	SELECT od.order_detail_id, s.service_name, s.unit, s.price, od.qty, (s.price * od.qty) as subtotal
	FROM order_detail od
	JOIN service s ON od.service_id = s.service_id
	WHERE od.order_id = $1
	ORDER BY od.order_detail_id
`, order_id)

if err != nil {
	log.Printf("Database error: %v", err)
	fmt.Println("An error occurred while retrieving order details. Press Enter to continue...")
	fmt.Scanln()
	return
}
defer rows.Close()

fmt.Println("\nOrder Items:")
fmt.Printf("%-5s %-25s %-10s %-10s %-10s %-10s\n", "ID", "Service", "Unit", "Price", "Qty", "Subtotal")
fmt.Println(strings.Repeat("-", 80))

var total_amount int = 0

for rows.Next() {
	var detail_id, price, qty, subtotal int
	var service_name, unit string

	err = rows.Scan(&detail_id, &service_name, &unit, &price, &qty, &subtotal)
	if err != nil {
		log.Printf("Row scan error: %v", err)
		continue
	}

	fmt.Printf("%-5d %-25s %-10s %-10d %-10d %-10d\n",
		detail_id, service_name, unit, price, qty, subtotal)

	total_amount += subtotal
}

fmt.Printf("\nTotal Amount: %d\n", total_amount)

fmt.Println("\nPress Enter to continue...")
fmt.Scanln()
}

