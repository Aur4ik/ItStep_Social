package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc{
	return func (c *gin.Context)  {
		role := c.GetString("role")
		fmt.Println("ROLE : ", role)
		fmt.Println("ADMIN MIDDLEWARE WORKS")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error" : "admin access required",
			})
			c.Abort()
			return
		}
		c.Next()
		
	}
}