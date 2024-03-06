package routes

import (
	"L0/cache"
	"L0/dbconnection/entity"
	"L0/dbconnection/repo"
	"encoding/json"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cache *cache.Cache) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	router.Static("/web", "./web")
	router.GET("/", func(c *gin.Context) {
		c.File("./web/index.html")
	})

	router.GET("/orders/:order_uid", func(ctx *gin.Context) {
		id := ctx.Param("order_uid")
		if order, exists := cache.Get(id); exists {
			ctx.JSON(200, order)
		} else {
			orderData, err := repo.GetOrderById(id)
			if err != nil {
				ctx.JSON(404, gin.H{"error": "Order not found"})
				return
			}
			var order entity.Order
			orderJSON, err := json.Marshal(orderData)
			if err != nil {
				ctx.JSON(500, gin.H{"error": "Failed to process order data"})
				return
			}
			if err := json.Unmarshal(orderJSON, &order); err == nil {
				cache.Set(id, order)
			}
			ctx.JSON(200, orderData)
		}
	})
	return router
}
