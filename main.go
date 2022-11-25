package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	userList := make(map[string]string, 2)
	userList["class1"] = "zl"
	userList["class2"] = "zhangsan"
	v1 := r.Group("/v1")
	v1.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": userList,
		})
	})
	v1.POST("/user", func(c *gin.Context) {
		user := User{}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "body field failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	v1.DELETE("/user", func(c *gin.Context) {
		iphoneNumber, ok := c.GetQuery("iphone")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "phone number is null",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"number":  iphoneNumber,
			"message": "success",
		})
	})
	v1.PUT("/user", func(c *gin.Context) {
		user := User{}
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "json field failed",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	v1.GET("/secret", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"secrets": userSecret,
		})
	})
	v1.DELETE("/secret", func(c *gin.Context) {
		secret, ok := c.GetQueryArray("secret")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "params failed",
			})
			return
		}

		deleteList := make([]string, 0, len(secret))
		for _, v := range secret {
			secret := v
			deleteList = append(deleteList, secret)
		}
		c.JSON(http.StatusOK, gin.H{
			"secret":  deleteList,
			"message": "success",
		})
	})
	v1.POST("/secret", func(c *gin.Context) {
		key, ok := c.GetPostForm("key")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "params failed",
			})
			return
		}
		hash := md5.New()
		data := hash.Sum([]byte(key))
		dataStr := fmt.Sprintf("%x", data)
		c.JSON(http.StatusOK, gin.H{
			"secret":  dataStr,
			"message": "success",
			"info":    "save ok ",
		})
	})
	v1.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "ok",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
