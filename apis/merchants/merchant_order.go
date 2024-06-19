package merchants

import (
	"123123/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func ViewOrderDaily(c *gin.Context) {
	var order_merchant OrderMerchant
	var orderId []int64
	//var orders []orders2.Order
	err := c.ShouldBindJSON(&order_merchant)

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tomorrow := today.AddDate(0, 0, 1)

	if err != nil {
		log.Printf("bind json mid error = %v\n", err)

		c.JSON(http.StatusOK, gin.H{
			"code":   "20000",
			"status": "bind json error",
		})
		return
	}

	log.Printf("today is %v\n mid is %v\n", today, order_merchant.Mid)

	database.DBS.Table("order_merchant").Raw("SELECT oid FROM order_merchant WHERE mid = ? AND time >= ? AND time < ?", order_merchant.Mid, today, tomorrow).Scan(&orderId)
	log.Printf("orders = %v\n", orderId)
	// after Raw operate , get orders array, which contains order match with merchant where mid = merchant.Mid

	database.DBS.Table("orders").Raw("SELECT ")

	c.JSON(http.StatusOK, gin.H{
		"code":   "OK",
		"status": "success",
	})
}
