package captcha

import (
	"123123/redis"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"log"
	"net/http"
)

func GenerateCaptchaHandler(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(40, 120, 5, 0.7, 80)
	captcha := base64Captcha.NewCaptcha(driver, redis.Store)
	id, b64s, _, err := captcha.Generate()

	if err != nil {
		log.Printf("generate captcha error = %v\n", err)

		c.JSON(http.StatusOK, gin.H{
			"code":   "20000",
			"status": "generate captcha error",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}

func VerifyCaptchaHandler(c *gin.Context) {
	var requestBody RequestBody

	err := c.ShouldBindJSON(&requestBody)
	log.Printf("this is requestBody %v\n", requestBody)

	if err != nil {
		log.Printf("bind json requestBody error = %v\n", err)
		return
	}

	ans := redis.Store.Verify(requestBody.CaptchaId, requestBody.CaptchaSolution, true)

	if ans == true {
		c.JSON(http.StatusOK, gin.H{
			"code":   "10000",
			"status": "verify success",
		})
	} else {
		log.Printf("this is solution %v, and this is v %v 11", requestBody.CaptchaSolution, redis.V)
		c.JSON(http.StatusOK, gin.H{
			"code":   "20000",
			"status": "Incorrect Captcha",
		})
	}

	return
}
