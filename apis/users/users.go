package users

import (
	"123123/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var user User

func GetUserByName(c *gin.Context) {
	var userByName ByName

	err := c.ShouldBindJSON(&userByName)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":   "30000",
			"status": "Unknown Error",
		})
		log.Println("error type 1")
		return
	}

	user.Name = userByName.Name
	var status int

	database.DBS.Table("users").Raw("SELECT status FROM users where name = ?", user.Name).Scan(&status)
	if status == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":   "20000",
			"status": "userNotFound",
		})

		log.Println("user not found")
		return
	}

	log.Printf("Get user by name %v\n", user)

	c.JSON(http.StatusOK, gin.H{
		"code":   "10000", // means backend successfully find user
		"status": "success",
	})
}

func GetUserByUsername(c *gin.Context) {
	var user ByUsername

	err := c.ShouldBindJSON(&user)

	if err != nil {
		log.Println("error type 1")
		return
	}
}

func AddUser(c *gin.Context) {
	var userByName ByName
	err := c.ShouldBindJSON(&userByName)

	if err != nil {
		log.Println("Parse Json Error In Add User")
	}

	var userTemp User
	database.DBS.Table("users").Raw("SELECT status FROM users WHERE NAME = ?", userByName.Name).Scan(&userTemp)

	if userTemp.Status != 0 {

		c.JSON(http.StatusOK, gin.H{
			"code":   "20000",
			"status": "User already exists",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":   "10000",
		"status": "Jump To Add Detail Page",
	})
}

func ModifyUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {
	var userByName ByName
	err := c.ShouldBindJSON(&userByName)

	if err != nil {
		log.Println("Parse Json Error In Delete User")
	}

	database.DBM.Table("users").Exec("DELETE FROM users WHERE name = ? IF EXISTS", userByName.Name)
	c.JSON(http.StatusOK, gin.H{
		"code":   "10000",
		"status": "successful",
	})
}
