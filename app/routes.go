package app

import "api-beer/controllers"

func routes() {
	router.GET("/beers/:beer_id", controllers.GetBeer)
	router.GET("/beers", controllers.GetAllBeers)
	router.POST("/beers", controllers.CreateBeer)

}
