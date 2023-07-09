package customer

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SetupRoutes(group *gin.RouterGroup, db *gorm.DB) {

	group.GET("/", func(c *gin.Context) {
		getAllCustomers(c, db)
	})
	group.GET("/:id", func(c *gin.Context) {
		getCustomer(c, db)
	})
	group.POST("/", func(c *gin.Context) {
		createCustomer(c, db)
	})
	group.PUT("/:id", func(c *gin.Context) {
		updateCustomer(c, db)
	})
	group.DELETE("/:id", func(c *gin.Context) {
		deleteCustomer(c, db)
	})
	// group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func getAllCustomers(c *gin.Context, db *gorm.DB) {
	log.Println("Handling GET request for all customers")

	var customers []Customer
	result := db.Find(&customers)
	if result.Error != nil {
		log.Printf("Failed to fetch customers: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to fetch customers"})
		return
	}

	log.Printf("Returning %d customers", len(customers))
	c.JSON(200, customers)
}

func getCustomer(c *gin.Context, db *gorm.DB) {
	var customer Customer
	id := c.Param("id")

	result := db.First(&customer, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Customer not found"})
		return
	}

	c.JSON(200, customer)
}

func createCustomer(c *gin.Context, db *gorm.DB) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&customer)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create customer"})
		return
	}

	c.JSON(201, customer)
}

func updateCustomer(c *gin.Context, db *gorm.DB) {
	var customer Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	result := db.First(&customer, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Customer not found"})
		return
	}

	result = db.Save(&customer)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update customer"})
		return
	}

	c.JSON(200, customer)
}

func deleteCustomer(c *gin.Context, db *gorm.DB) {
	var customer Customer
	id := c.Param("id")

	result := db.First(&customer, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Department not found"})
		return
	}

	result = db.Delete(&customer)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete customer"})
		return
	}

	c.JSON(204, nil)
}
