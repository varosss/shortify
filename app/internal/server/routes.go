package server

import (
	"net/http"
	"shortify/internal/auth"
	"shortify/internal/links"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	authHandler *auth.AuthHandler,
	linksHandler *links.LinksHandler,
	jwtManager *auth.JWTManager,
) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.POST("/sign-up", authHandler.SignUp)
	r.POST("/login", authHandler.Login)

	api := r.Group("/api", auth.AuthMiddleware(jwtManager))

	api.POST("/links", linksHandler.AddLinks)
	api.GET("/links", linksHandler.GetLinks)
	api.GET("/links/:short", linksHandler.GetLink)
	api.DELETE("/links/:short", linksHandler.DeleteLink)

	r.GET("/:short", linksHandler.Redirect)

	return r
}
