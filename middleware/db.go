package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetDBConnection puts db connection to gin context
func SetDBConnection(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}