package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Calculator struct {
	Num1 int `json:"num1" validate:"required"`
	Num2 int `json:"num2" validate:"required"`
}

func Tambah() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var calculator Calculator

		if err := c.BindJSON(&calculator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}

		result := calculator.Num1 + calculator.Num2

		defer cancel()

		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

func Kurang() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var calculator Calculator

		if err := c.BindJSON(&calculator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}

		result := calculator.Num1 - calculator.Num2

		defer cancel()

		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

func Kali() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var calculator Calculator

		if err := c.BindJSON(&calculator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}

		result := calculator.Num1 * calculator.Num2

		defer cancel()

		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

func Bagi() gin.HandlerFunc {
	return func(c *gin.Context) {
		var _, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var calculator Calculator

		if err := c.BindJSON(&calculator); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		}

		result := calculator.Num1 / calculator.Num2

		defer cancel()

		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}
