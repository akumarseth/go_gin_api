package order

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	// "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Product string  `json:"product"`
	Price   float64 `json:"price"`
}

func SetupRoutes(group *gin.RouterGroup, db *gorm.DB) {

	group.GET("/", func(c *gin.Context) {
		getAllOrders(c, db)
	})
	group.GET("/:id", func(c *gin.Context) {
		getOrder(c, db)
	})
	group.POST("/", func(c *gin.Context) {
		createOrder(c, db)
	})
	group.PUT("/:id", func(c *gin.Context) {
		updateOrder(c, db)
	})
	group.DELETE("/:id", func(c *gin.Context) {
		deleteOrder(c, db)
	})
	// group.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func getAllOrders(c *gin.Context, db *gorm.DB) {
	log.Println("Handling GET request for all orders")

	var orders []Order
	result := db.Find(&orders)
	if result.Error != nil {
		log.Printf("Failed to fetch orders: %v", result.Error)
		c.JSON(500, gin.H{"error": "Failed to fetch orders"})
		return
	}

	log.Printf("Returning %d orders", len(orders))
	c.JSON(200, orders)
}

func getOrder(c *gin.Context, db *gorm.DB) {
	var order Order
	id := c.Param("id")

	result := db.First(&order, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(200, order)
}

func createOrder(c *gin.Context, db *gorm.DB) {
	var order Order
	fmt.Println(order)
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := db.Create(&order)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(201, order)
}

func updateOrder(c *gin.Context, db *gorm.DB) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	result := db.First(&order, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "order not found"})
		return
	}

	result = db.Save(&order)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(200, order)
}

func deleteOrder(c *gin.Context, db *gorm.DB) {
	var order Order
	id := c.Param("id")

	result := db.First(&order, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Order not found"})
		return
	}

	result = db.Delete(&order)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(204, nil)
}
