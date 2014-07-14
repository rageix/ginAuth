package controllers

import (
	"github.com/gin-gonic/gin"
)

func Is404(ctx *gin.Context) {
		ctx.String(404, "This is an invalid route and/or HTTP method.")
}
