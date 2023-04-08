package controller

import (
	"fmt"
	"net/http"

	"github.com/EmeraldLS/quote-generator/helper"
	"github.com/EmeraldLS/quote-generator/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateQuoteFromFrontend(c *gin.Context) {
	var quote model.Quote
	if err := c.BindJSON(quote); err != nil {
		fmt.Println(err)
		c.Abort()
	}
	validate := validator.New()
	if err := validate.Struct(quote); err != nil {
		fmt.Println(err)
		c.Abort()
		return
	}
	insertionID, err := helper.CreateQuote(quote)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"Insertion ID": insertionID,
	})
}

func GetAllQuotesFromFrontend(c *gin.Context) {
	quotes := helper.GetAllQuotes()
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    quotes,
	})
}

func GetRelatedQuoteFromFrontend(c *gin.Context) {
	quotes := helper.GetRelatedQuote(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    quotes,
	})
}

func PopulateDBFromFrontnend(c *gin.Context) {
	message := helper.Populate()
	c.JSON(http.StatusOK, message)
}

func DeleteAllFromFrontend(c *gin.Context) {
	count := helper.DeleteAllQuotes()
	c.JSON(http.StatusOK, gin.H{
		"message":       "success",
		"Deleted Count": count,
	})
}

func DeleteOneFromFrontend(c *gin.Context) {
	count := helper.DeleteQuote(c)
	c.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"Delete Count": count,
	})
}

func DeleteSimilarFromFrontend(c *gin.Context) {
	count := helper.DeleteAllSimilarQuotes(c)
	c.JSON(http.StatusOK, gin.H{
		"message":      "success",
		"Delete Count": count,
	})
}
