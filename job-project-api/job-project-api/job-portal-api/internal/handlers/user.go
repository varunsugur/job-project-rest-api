package handlers

import (
	"encoding/json"
	"golang/internal/middlewares"
	"golang/internal/models"
	"golang/internal/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type handler struct {
	service service.UserService
}

func (h *handler) Signup(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if !ok {
		log.Error().Msg("missing context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return

	}
	var nu models.NewUser

	err := json.NewDecoder(c.Request.Body).Decode(&nu)

	if err != nil {
		log.Error().Err(err).Str("Trace is ", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": http.StatusInternalServerError})
		return
	}
	validate := validator.New()
	err = validate.Struct(nu)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": "please provide valid username, email and password"})
		return
	}

	userDetails, err := h.service.UserSignup(ctx, nu)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": "please provide valid username, email and password"})
		return
	}
	c.JSON(http.StatusOK, userDetails)
}

func (h *handler) Signin(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewares.TraceIdKey).(string)
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
	if !ok {
		log.Error().Msg("missing context")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusInternalServerError})
		return
	}

	var u models.UserLogin

	err := json.NewDecoder(c.Request.Body).Decode(&u)
	if err != nil {
		log.Error().Err(err).Str("Trace is ", traceId)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Msg": http.StatusInternalServerError})
		return
	}

	token, err := h.service.UserSignin(ctx, u)
	if err != nil {
		log.Error().Err(err).Str("trace id", traceId)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}
