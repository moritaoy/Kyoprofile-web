package routes

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func LogIn(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func Main(ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "main.html", gin.H{})
}

func NoRoute(ctx *gin.Context) {
    ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
