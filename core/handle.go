package core

import "github.com/gin-gonic/gin"

type HandlerFunc func(c *Context)

const ContextKey = "__context__"

func Handle(h HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := getContext(c)
		h(ctx)
	}
}

func getContext(c *gin.Context) *Context {
	ctx, ok := c.Get(ContextKey)
	if ok {
		return ctx.(*Context)
	}

	context := &Context{
		c,
	}
	c.Set(ContextKey, context)
	return context
}
