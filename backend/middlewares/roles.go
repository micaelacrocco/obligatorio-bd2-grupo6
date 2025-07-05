package middlewares

import (
    "net/http"
    "EleccionesUcu/models"
    "github.com/gin-gonic/gin"
)

func RequireRoles(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRaw, exists := c.Get("user")
        if !exists {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
            return
        }

        claims := userRaw.(*models.JwtCustomClaims)
        for _, role := range roles {
            if claims.UserType == role {
                c.Next()
                return
            }
        }

        c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
    }
}
