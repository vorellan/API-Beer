package services

import (
	"api-beer/models"
	"api-beer/utils/error_utils"

)

var (
	BeerService beerServiceInterface = &beerService{}
)

type beerService struct{}

type beerServiceInterface interface {
	GetBeer(int64) (*models.Beer, error_utils.MessageErr)
	CreateBeer(*models.Beer) (*models.Beer, error_utils.MessageErr)
	GetAllBeers() ([]models.Beer, error_utils.MessageErr)
}

func (m *beerService) GetBeer(msgId int64) (*models.Beer, error_utils.MessageErr) {
	beer, err := models.Server.Get(msgId)
	if err != nil {
		return nil, err
	}
	return beer, nil
}

func (m *beerService) GetAllBeers() ([]models.Beer, error_utils.MessageErr) {
	beer, err := models.Server.GetAll()
	if err != nil {
		return nil, err
	}
	return beer, nil
}

func (m *beerService) CreateBeer(beer *models.Beer) (*models.Beer, error_utils.MessageErr) {
	if err := beer.Validate(); err != nil {
		return nil, err
	}
	
	beer, err := models.Server.Create(beer)
	if err != nil {
		return nil, err
	}
	return beer, nil
}

