package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("recomendation", func(ctx *gin.Context) {
		var input struct {
			Budget float64 `json:"budget"`
			Purpose string `json:"purpose"`
		}

		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid input"})
			return
		}

		// Berdasarkan purpose, tentukan rekomendasi laptop
		var recomendation string
		if input.Purpose == "gaming" {
			recomendation = fmt.Sprintf("With a budget of Rp %.0f, you can get a high-performance gaming laptop that can handle the latest games and provide a smooth gaming experience. Here's a recommendation for a gaming laptop within your budget: ...", input.Budget)
		} else {
			recomendation = "Sorry, we currently don't have recommendations for the specified purpose."
		}

		response := gin.H{
			"status": "success",
			"data": recomendation,
		}

		ctx.JSON(http.StatusOK, response)
	})

	router.Run(":8080")
}
