# Golang Project (Laundry)
sebuah aplikasi sederhana berbasis console untuk mencatat transaksi di toko

#Features
-Customer management
-Service management
-Order processing and tracking

#Requirement
-Golang new version
-PostgreSQL database
-github.com/lib/pq(PostgreSQL driver for Go)

#Installation
-Clone the repository or download the source code
-Install dependencies: "go get github.com/lib/pq"
-Set up the PostgreSQL database: Create database, execute DDL script to create the required tables
-Configure the database connection in file "main.go":
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "postgres"
    dbname   = "enigma_laundry"
)
-Build and run the application

#Usage Guide
>Main Menu
The application provides four main options:
-Customer - Manage customer data
-Service - Manage service offerings
-Order - Process and track orders
-Exit - Exit the application

CUSTOMER MENU
-->Create Customer
-Enter a unique customer ID
-Enter customer details (name, phone, address)
-The system validates that the ID is not already in use

-->View Of List Customer
-Displays all registered customers with their IDs, names, and phone numbers

-->View Details Customer By ID
-Enter a customer ID to view complete customer information
-Includes created and updated timestamps

-->Update Customer
-Enter a customer ID to update
-Enter new information for name, phone, and address
-The system validates that the customer exists

-->Delete Customer
-Enter a customer ID to delete
-The system checks if the customer exists
-The system prevents deletion if the customer has associated orders

-->Back to Main Menu
-Return to the main menu



SERVICE MENU
-->Create Service
-Enter a unique service ID
-Enter service details (name, unit, price)
-The system validates that the ID is not already in use

-->View Of List Service
-Displays all registered services with their IDs, names, units, and prices

-->View Details Service By ID
-Enter a service ID to view complete service information
-Includes created and updated timestamps

-->Update Service
-Enter a service ID to update
-Enter new information for name, unit, and price
-The system validates that the service exists

-->Delete Service
-Enter a service ID to delete
-The system checks if the service exists
-The system prevents deletion if the service is used in any orders

-->Back to Main Menu
-Return to the main menu



ORDER MENU

-->Create Order
-Enter a unique order ID
-Enter customer ID (validated to ensure it exists)
-Enter who received the order
-Add multiple services with quantities to the order
-Enter 0 for Service ID to finish adding services

-->Complete Order
-Enter an order ID to mark as completed
-Enter the completion date in YYYY-MM-DD format
-The system validates that the order exists

-->View Of List Order
-Displays all orders with their basic information
-Shows customer name, order date, completion status, and who received the order


-->View Order Details By ID
-Enter an order ID to view complete order information
-Shows order header information and all ordered services
-Calculates subtotals for each service and the total amount

-->Back to Main Menu
-Return to the main menu
