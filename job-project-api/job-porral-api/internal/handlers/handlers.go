package handlers

import (
	"golang/internal/auth"
	"golang/internal/middlewares"
	"golang/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func API(a *auth.Auth, svc service.UserService) *gin.Engine {
	router := gin.New()

	m, err := middlewares.NewMid(a)
	if err != nil {
		log.Panic().Msg("Middleware not set up")

	}
	h := handler{
		service: svc,
	}

	router.Use(m.Log(), gin.Recovery())

	router.GET("/check", m.Authenticate(check))
	router.POST("/signup", h.Signup)
	router.POST("/login", h.Signin)
	router.POST("/companies", m.Authenticate(h.AddCompany))
	router.GET("/companies", m.Authenticate(h.ViewAllCompanies))
	router.GET("/companies/:id", m.Authenticate(h.ViewCompany))
	return router
}

func check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Msg": "Good"})

}
