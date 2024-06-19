package login

import (
	"123123/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func PostLoginHandler(c *gin.Context) {
	var user User

	err := c.ShouldBindJSON(&user)
	log.Printf("login user = %v\n", user)

	if err != nil {
		log.Printf("login bind json error = %v\n", err)
		return
	}

	var password string
	database.DBS.Table("employee").Raw("select password from employee where username = ?", user.Username).Scan(&password)

	if password != user.Password {
		log.Printf("login password does not match")

		c.JSON(http.StatusOK, gin.H{
			"code":   "20000",
			"status": "password does not match",
		})

		return
	} else {
		jwtToken, err := NewJwt()

		if err != nil {
			log.Printf("new jwtToken error = %v\n", err)
			c.JSON(http.StatusOK, gin.H{
				"code":   "20000",
				"status": "login failed",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"code":   "10000",
			"status": "login success",
			"token":  jwtToken,
		})
	}
}
