package users

import (
	"123123/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	log.Println("get user info" + user.Name)

	var userTemp User
	database.DBS.Table("users").Raw("SELECT * FROM users WHERE name = ?", user.Name).Scan(&userTemp)

	log.Printf("total user Info %v", userTemp)

	c.JSON(http.StatusOK, gin.H{
		"name":      userTemp.Name,
		"password":  "***************",
		"email":     userTemp.Email,
		"code":      "10000",
		"status":    userTemp.Status,
		"school":    userTemp.School,
		"studentId": userTemp.Sid,
		"phone":     userTemp.Phone,
	})
}
