package controllers

import (
	"api-beer/models"
	"api-beer/services"
	"api-beer/utils/error_utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

)

func getBeerId(brIdParam string) (int64, error_utils.MessageErr) {
	brId, brErr := strconv.ParseInt(brIdParam, 10, 64)
	if brErr != nil {
		return 0, error_utils.NewBadRequestError("beer id isn't a number")
	}
	return brId, nil
}

func GetBeer(c *gin.Context) {
	brId, err := getBeerId(c.Param("beer_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	beer, getErr := services.BeerService.GetBeer(brId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, beer)
}

func GetAllBeers(c *gin.Context) {
	beers, getErr := services.BeerService.GetAllBeers()
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, beers)
}

func CreateBeer(c *gin.Context) {
	var beer models.Beer
	if err := c.ShouldBindJSON(&beer); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError("invalid json body")
		c.JSON(theErr.Status(), theErr)
		return
	}
	br, err := services.BeerService.CreateBeer(&beer)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, br)
}

func CalculateBeer(c *gin.Context) {
	var f interface{}

	if err := c.ShouldBindJSON(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quantity := f.(map[string]interface{})["quantity"]
	price := f.(map[string]interface{})["price"]

	var quantityInt int = int(quantity.(float64))
	var priceInt int = int(price.(float64))
	total := quantityInt * priceInt
	
	// c.JSON(http.StatusOK, gin.H{"quantity": quantity})
	// c.JSON(http.StatusOK, gin.H{"price": price})
	c.JSON(http.StatusOK, gin.H{"total": total})
}


