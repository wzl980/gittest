package main

import (
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
		secret, ok := c.GetQuery("secret")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "params failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"secret":  secret,
			"message": "success",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
