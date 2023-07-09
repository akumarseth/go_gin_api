package employee

import (
	"log"

	"github.com/gin-gonic/gin"
	// "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name     string `json:"name"`
	Position string `json:"position"`
}

func SetupRoutes(group *gin.RouterGroup, db *gorm.DB) {
	group.GET("/", func(c *gin.Context) {
		getAllEmployees(c, db)
	})
	group.GET("/:id", func(c *gin.Context) {
		getEmployee(c, db)
	})
	group.POST("/", func(c *gin.Context) {
		createEmployee(c, db)
	})
	group.PUT("/:id", func(c *gin.Context) {
		updateEmployee(c, db)
	})
	group.DELETE("/:id", func(c *gin.Context) {
		deleteEmployee(c, db)
	})
	// group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func getAllEmployees(c *gin.Context, db *gorm.DB) {
	log.Println("Handling GET request for all employees")

	var employees []Employee
	result := db.Find(&employees)
	if result.Error != nil {
		log.Printf("Failed to fetch employees: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to fetch employees"})
		return
	}

	log.Printf("Returning %d employees", len(employees))
	c.JSON(200, employees)
}

func getEmployee(c *gin.Context, db *gorm.DB) {
	var employee Employee
	id := c.Param("id")

	result := db.First(&employee, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(200, employee)
}

func createEmployee(c *gin.Context, db *gorm.DB) {
	var employee Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&employee)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create employee"})
		return
	}

	c.JSON(201, employee)
}

func updateEmployee(c *gin.Context, db *gorm.DB) {
	var employee Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	result := db.First(&employee, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Employee not found"})
		return
	}

	result = db.Save(&employee)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update employee"})
		return
	}

	c.JSON(200, employee)
}

func deleteEmployee(c *gin.Context, db *gorm.DB) {
	var employee Employee
	id := c.Param("id")

	result := db.First(&employee, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Employee not found"})
		return
	}

	result = db.Delete(&employee)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete employee"})
		return
	}

	c.JSON(204, nil)
}
