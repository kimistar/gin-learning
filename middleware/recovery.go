package middleware

import (
	"gin-learning/core"
	"net/http"
)

func Recovery(c *core.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.HTML(http.StatusOK, "error/500", nil)
		}
	}()
	c.Next()
}
