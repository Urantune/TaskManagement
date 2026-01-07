package middleware

import (
	"ManageTask/utils"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func MidwareAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		claim, err := utils.ParseToken(token)
		if err != nil {

			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		c.Set("user", claim)
		c.Next()
	}
}

func CheckRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, ok := c.Get("user")

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		claim, ok := user.(*utils.Claims)

		if !ok {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		if claim.UserRole != role {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}

type userLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	mu       sync.Mutex
	limiters = make(map[string]*userLimiter)
)

func RateLimitByUser(connect rate.Limit, burst int) gin.HandlerFunc {

	go func() {
		for {
			time.Sleep(time.Minute)
			mu.Lock()
			for k, v := range limiters {
				if time.Since(v.lastSeen) > 3*time.Minute {
					delete(limiters, k)
				}
			}
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		user, ok := c.Get("user")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		claim := user.(*utils.Claims)
		key := claim.UserEmail

		mu.Lock()
		ul, exists := limiters[key]
		if !exists {
			ul = &userLimiter{
				limiter: rate.NewLimiter(connect, burst),
			}
			limiters[key] = ul
		}
		ul.lastSeen = time.Now()
		lim := ul.limiter
		mu.Unlock()

		if !lim.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many requests, slow down",
			})
			return
		}

		c.Next()
	}
}
