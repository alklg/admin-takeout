package dishes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AddDish(c *gin.Context) {
	var dish Dish
	var merchant Merchants
	log.Println(merchant)

	err := c.ShouldBindJSON(&dish)
	if err != nil {
		log.Printf("Add dish error = %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"code":   "10000",
	})
}
