package persons

import (
	"encoding/json"
	"fmt"
	"net/http"
	"person-predicator/internal/logger"
	"person-predicator/internal/server/handlers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AddPersonRequest struct {
	Name       string `json:"name" example:"Dmitriy"`
	Surname    string `json:"surname" example:"Ushakov"`
	Patronymic string `json:"patronymic" example:"Vasilevich"`
}

//	@Summary		Add a person
//	@Description	Add a new person to database
//	@Tags			Person
//	@Accept			json
//	@Produce		json
//	@Param			request	body	AddPersonRequest	true	"Add Person Request"
//	@Success		201
//	@Failure		400	{object}	handlers.ErrorResponce
//	@Failure		500	{object}	handlers.ErrorResponce
//	@Failure		502	{object}	handlers.ErrorResponce
//	@Router			/person [post]
func (h *PersonHandler) Add(ctx *gin.Context) {
	var req AddPersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}

	age, err := getAge(req.Name)
	if err != nil {
		logger.Logger.Error("age not founded", zap.Error(err))
		ctx.JSON(http.StatusBadGateway,
			handlers.ErrorResponce{Message: "Failed to get age predict", Error: err})
		return
	}
	gender, err := getGender(req.Name)
	if err != nil {
		logger.Logger.Error("gender not founded", zap.Error(err))
		ctx.JSON(http.StatusBadGateway,
			handlers.ErrorResponce{Message: "Failed to get gender predict", Error: err})
		return
	}
	country, err := getCountry(req.Name)
	if err != nil {
		logger.Logger.Error("country not founded", zap.Error(err))
		ctx.JSON(http.StatusBadGateway,
			handlers.ErrorResponce{Message: "Failed to get country predict", Error: err})
		return
	}

	_, err = h.service.NewPerson(ctx, req.Name, req.Surname, req.Patronymic, age, gender, country)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to add a new person to database", Error: err})
		return
	}
	logger.Logger.Info("Person added")

	ctx.JSON(http.StatusCreated, nil)
}

type ageResponse struct {
	Age int `json:"age"`
}

func getAge(name string) (int, error) {
	response, err := http.Get(fmt.Sprintf("https://api.agify.io/?name=%s", name))
	if err != nil {
		return 0, err
	}
	var age ageResponse
	if err = json.NewDecoder(response.Body).Decode(&age); err != nil {
		return 0, err
	}
	return age.Age, nil
}

type genderResponse struct {
	Gender      string  `json:"gender"`
	Probability float64 `json:"probability"`
}

func getGender(name string) (string, error) {
	response, err := http.Get(fmt.Sprintf("https://api.genderize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	var gender genderResponse
	if err = json.NewDecoder(response.Body).Decode(&gender); err != nil {
		return "", err
	}
	return gender.Gender, nil
}

type country struct {
	Country     string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type countryResponse struct {
	Countries []country `json:"country"`
}

func getCountry(name string) (string, error) {
	response, err := http.Get(fmt.Sprintf("https://api.nationalize.io/?name=%s", name))
	if err != nil {
		return "", err
	}
	var countries countryResponse
	if err = json.NewDecoder(response.Body).Decode(&countries); err != nil {
		return "", err
	}
	var maximumProbability float64
	var countryId string
	for _, country := range countries.Countries {
		if country.Probability > maximumProbability {
			maximumProbability = country.Probability
			countryId = country.Country
		}
	}
	return countryId, nil
}
