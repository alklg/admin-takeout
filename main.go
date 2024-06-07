package main

import (
	"123123/apis/users"
	"123123/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://119.45.145.96:8001"},     // 允许的前端应用地址
		AllowMethods: []string{"GET", "POST", "OPTIONS"},        // 允许的 HTTP 方法
		AllowHeaders: []string{"Authorization", "Content-Type"}, // 允许的请求头字段
	}))

	file, err := os.OpenFile("develop.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("set log error = %v", err)
		return
	}

	defer file.Close()

	log.SetOutput(file)
	log.Println("This is a log message written to the file")

	err = database.InitDatabase()

	if err != nil {
		log.Printf("database init error = %v", err)
		return
	}

	api := r.Group("/api", func(c *gin.Context) {

		c.Next()
	})
	{
		user := api.Group("/user")
		{
			user.POST("/add", users.AddUser)
			user.POST("/find", users.GetUserByName)
			user.POST("/view", users.GetUserInfo)
			user.POST("/delete", users.DeleteUser)
			user.POST("/modify", users.ModifyUser)
		}

		dish := api.Group("/dish")
		{
			dish.POST("/add", dishes.AddDish)
		}
	}

	r.Run(":8082")
}
