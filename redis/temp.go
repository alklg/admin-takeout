package redis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"time"
)

type kv struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func Temp(c *gin.Context) {
	var obj kv
	err := c.ShouldBindJSON(&obj)
	if err != nil {
		log.Printf("bind json error = %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 使用事务同时添加到有序集合和设置过期时间
	pipe := rdb.TxPipeline()
	pipe.ZAdd(ctx, "sorted_set", &redis.Z{
		Score:  float64(obj.Value),
		Member: obj.Key,
	})
	pipe.Set(ctx, "expire:"+obj.Key, "", 15*time.Second)

	_, err = pipe.Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Sorted handles retrieving the sorted kv pairs
func Sorted(c *gin.Context) {
	vals, err := rdb.ZRangeWithScores(ctx, "sorted_set", 0, -1).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var validVals []redis.Z
	for _, val := range vals {
		exists, err := rdb.Exists(ctx, "expire:"+val.Member.(string)).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if exists == 1 {
			validVals = append(validVals, val)
		} else {
			// 如果过期了，从有序集合中移除
			rdb.ZRem(ctx, "sorted_set", val.Member)
		}
	}

	c.JSON(http.StatusOK, validVals)
}
