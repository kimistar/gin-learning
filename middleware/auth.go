package middleware

import "core"

func Auth(c *core.Context) {
	if false {
		c.Fail(401, "unauthorized", nil)
	}
	c.Next()
}
