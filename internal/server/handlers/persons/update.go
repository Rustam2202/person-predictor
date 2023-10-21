package persons

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdatePersonRequest struct {
	Id         int64  `json:"id" example:"123456789"`
	Name       string `json:"name" example:"Dmitriy"`
	Surname    string `json:"surname" example:"Ushakov"`
	Patronymic string `json:"patronymic" example:"Vasilevich"`
	Age        int    `json:"age" example:"42"`
	Gender     string `json:"gender" example:"male"`
	Country    string `json:"country" example:"Russia"`
}

// @Summary		Update a person
// @Description	Update a person in database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			request	body		UpdatePersonRequest	true	"Update Person Request"
// @Success		200
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/person [put]
func (h *PersonHandler) Update(ctx *gin.Context) {
	var req UpdatePersonRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		// handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.Update(ctx, req.Id, req.Name, req.Surname, req.Patronymic, req.Age, req.Gender, req.Country)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// handlers.ErrorResponce{Message: "Error with update a person in database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type UpdateNameRequest struct {
	Id   int64  `json:"id" example:"123456789"`
	Name string `json:"name" example:"Dmitriy"`
}

// @Summary		Update a person
// @Description	Update a person in database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateNameRequest	true	"Update Person Request"
// @Success		200
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/person [put]
func (h *PersonHandler) UpdateName(ctx *gin.Context) {
	var req UpdateNameRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		// handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.UpdateName(ctx, req.Id, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// handlers.ErrorResponce{Message: "Error with update a person in database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type UpdateSurnameRequest struct {
	Id      int64  `json:"id" example:"123456789"`
	Surname string `json:"surname" example:"Ushakov"`
}

// @Summary		Update a person
// @Description	Update a person in database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateSurnameRequest	true	"Update Person Request"
// @Success		200
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/person [put]
func (h *PersonHandler) UpdateSurname(ctx *gin.Context) {
	var req UpdateSurnameRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		// handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.UpdateSurname(ctx, req.Id, req.Surname)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// handlers.ErrorResponce{Message: "Error with update a person in database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type UpdatePatronymicRequest struct {
	Id         int64  `json:"id" example:"123456789"`
	Patronymic string `json:"patronymic" example:"Vasilevich"`
}

// @Summary		Update a person
// @Description	Update a person in database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			request	body		UpdatePatronymicRequest	true	"Update Person Request"
// @Success		200		{object}	UpdatePersonRequest
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/person [put]
// func (h *PersonHandler) UpdatePatronymic(ctx *gin.Context) {
// 	var req UpdatePatronymicRequest
// 	err := ctx.ShouldBindJSON(&req)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, nil)
// 		// handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
// 		return
// 	}
// 	err = h.service.UpdatePatronymic(ctx, req.Id, req.Patronymic)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, nil)
// 		// handlers.ErrorResponce{Message: "Error with update a person in database", Error: err})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, nil)
// }

type UpdateAgeRequest struct {
	Id  int64 `json:"id" example:"123456789"`
	Age int   `json:"age" example:"42"`
}

// @Summary		Update a person
// @Description	Update a person in database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateAgeRequest	true	"Update Person Request"
// @Success		200
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/person [put]
func (h *PersonHandler) UpdateAge(ctx *gin.Context) {
	var req UpdateAgeRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		// handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.UpdateAge(ctx, req.Id, req.Age)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// handlers.ErrorResponce{Message: "Error with update a person in database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

type UpdateGenderRequest struct {
	Id     int64  `json:"id" example:"123456789"`
	Gender string `json:"gender" example:"male"`
}

// @Summary		Update a person
// @Description	Update a person in database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateGenderRequest	true	"Update Person Request"
// @Success		200		{object}	UpdatePersonRequest
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/person [put]
// func (h *PersonHandler) UpdateGender(ctx *gin.Context) {
// 	var req UpdateGenderRequest
// 	err := ctx.ShouldBindJSON(&req)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, nil)
// 		// handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
// 		return
// 	}
// 	err = h.service.UpdateGender(ctx, req.Id, req.Gender)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, nil)
// 		// handlers.ErrorResponce{Message: "Error with update a person in database", Error: err})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, nil)
// }

type UpdateCountryRequest struct {
	Id      int64  `json:"id" example:"123456789"`
	Country string `json:"country" example:"Russia"`
}

// @Summary		Update a person
// @Description	Update a person in database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateCountryRequest	true	"Update Person Request"
// @Success		200
// @Failure		400		{object}	handlers.ErrorResponce
// @Failure		500		{object}	handlers.ErrorResponce
// @Router			/person [put]
func (h *PersonHandler) UpdateCountry(ctx *gin.Context) {
	var req UpdateCountryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		// handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	err = h.service.UpdateCountry(ctx, req.Id, req.Country)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// handlers.ErrorResponce{Message: "Error with update a person in database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, req)
}
