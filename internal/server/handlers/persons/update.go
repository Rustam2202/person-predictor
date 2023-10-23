package persons

import (
	"net/http"
	"person-predicator/internal/logger"
	"person-predicator/internal/server/handlers"

	"github.com/gin-gonic/gin"
)

type UpdatePersonRequest struct {
	Id         int64  `json:"id" example:"123456789"`
	Name       string `json:"name" example:"Dmitriy"`
	Surname    string `json:"surname" example:"Ushakov"`
	Patronymic string `json:"patronymic" example:"Vasilevich"`
	Age        int    `json:"age" example:"42"`
	Gender     string `json:"gender" example:"male"`
	Country    string `json:"country" example:"RU"`
}

// @Summary		Update a person
// @Description	Update a person in database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			request	body	UpdatePersonRequest	true	"Update Person Request"
// @Success		200
// @Failure		400	{object}	handlers.ErrorResponce
// @Failure		500	{object}	handlers.ErrorResponce
// @Router			/person [put]
func (h *PersonHandler) Update(ctx *gin.Context) {
	var req UpdatePersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	if req.Id <= 0 {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Incorrect Id", Error: err})
		return
	}

	err = h.service.Update(ctx, req.Id, req.Name, req.Surname, req.Patronymic, req.Age, req.Gender, req.Country)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Error with update a person in database", Error: err})
		return
	}
	logger.Logger.Debug("Person updated")

	ctx.JSON(http.StatusOK, nil)
}
