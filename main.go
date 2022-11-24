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

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
