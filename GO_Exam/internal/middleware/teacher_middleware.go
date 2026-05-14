package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TeacherOrAdmin() gin.HandlerFunc {

	return func(c *gin.Context) {

		role := c.GetString("role")

		if role != "teacher" && role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "teacher or admin access required",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}