package department

import (
	"log"

	"github.com/gin-gonic/gin"
	// "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name string `json:"name"`
}

func SetupRoutes(group *gin.RouterGroup, db *gorm.DB) {
	group.GET("/", func(c *gin.Context) {
		getAllDepartments(c, db)
	})
	group.GET("/:id", func(c *gin.Context) {
		getDepartment(c, db)
	})
	group.POST("/", func(c *gin.Context) {
		createDepartment(c, db)
	})
	group.PUT("/:id", func(c *gin.Context) {
		updateDepartment(c, db)
	})
	group.DELETE("/:id", func(c *gin.Context) {
		deleteDepartment(c, db)
	})
	// group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func getAllDepartments(c *gin.Context, db *gorm.DB) {
	log.Println("Handling GET request for all departments")

	var departments []Department
	result := db.Find(&departments)
	if result.Error != nil {
		log.Printf("Failed to fetch departments: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to fetch departments"})
		return
	}

	log.Printf("Returning %d departments", len(departments))
	c.JSON(200, departments)
}

func getDepartment(c *gin.Context, db *gorm.DB) {
	var department Department
	id := c.Param("id")

	result := db.First(&department, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Department not found"})
		return
	}

	c.JSON(200, department)
}

func createDepartment(c *gin.Context, db *gorm.DB) {
	var department Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&department)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create department"})
		return
	}

	c.JSON(201, department)
}

func updateDepartment(c *gin.Context, db *gorm.DB) {
	var department Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	result := db.First(&department, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Department not found"})
		return
	}

	result = db.Save(&department)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update department"})
		return
	}

	c.JSON(200, department)
}

func deleteDepartment(c *gin.Context, db *gorm.DB) {
	var department Department
	id := c.Param("id")

	result := db.First(&department, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Department not found"})
		return
	}

	result = db.Delete(&department)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete department"})
		return
	}

	c.JSON(204, nil)
}
