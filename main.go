package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gostudio_app/common"
	"gostudio_app/customer"
	"gostudio_app/department"
	"gostudio_app/employee"
	"gostudio_app/order"
)

var db *gorm.DB // Declare a global variable for the database connection

func main() {

	db := common.SetupDatabase()

	r := gin.Default()

	customerGroup := r.Group("/customers")
	customer.SetupRoutes(customerGroup, db)

	orderGroup := r.Group("/orders")
	order.SetupRoutes(orderGroup, db)

	employeeGroup := r.Group("/employees")
	employee.SetupRoutes(employeeGroup, db)

	departmentGroup := r.Group("/departments")
	department.SetupRoutes(departmentGroup, db)

	r.Run(":8080")
}

func SetupDatabase() {
	panic("unimplemented")
}
