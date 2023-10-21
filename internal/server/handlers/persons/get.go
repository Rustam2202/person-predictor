package persons

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary		Get a person
// @Description	Get a person from database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Id"
// @Success		200	{object}	domain.Person
// @Failure		400	{object}	handlers.ErrorResponce
// @Failure		500	{object}	handlers.ErrorResponce
// @Router			/person/{id} [get]
func (h *PersonHandler) Get(ctx *gin.Context) {
	req := ctx.Param("id")
	id, err := strconv.ParseInt(req, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		// handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	per, err := h.service.Get(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

// @Summary		Get a person
// @Description	Get a person from database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			name	path		string	true	"Name"
// @Success		200	{object}	[]domain.Person
// @Failure		500	{object}	handlers.ErrorResponce
// @Router			/person/{name} [get]
func (h *PersonHandler) GetByName(ctx *gin.Context) {
	req := ctx.Param("name")
	per, err := h.service.GetByName(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

// @Summary		Get a person
// @Description	Get a person from database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			surname	path		string	true	"Surname"
// @Success		200	{object}	[]domain.Person
// @Failure		500	{object}	handlers.ErrorResponce
// @Router			/person/{surname} [get]
func (h *PersonHandler) GetBySurname(ctx *gin.Context) {
	req := ctx.Param("surname")
	per, err := h.service.GetBySurname(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

// @Summary		Get a person
// @Description	Get a person from database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			age	path		int	true	"Age"
// @Success		200	{object}	[]domain.Person
// @Failure		400	{object}	handlers.ErrorResponce
// @Failure		500	{object}	handlers.ErrorResponce
// @Router			/person/{name} [get]
func (h *PersonHandler) GetByAge(ctx *gin.Context) {
	req := ctx.Param("age")
	age, err := strconv.ParseInt(req, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		// handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	per, err := h.service.GetByAge(ctx, int(age))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}

// @Summary		Get a person
// @Description	Get a person from database
// @Tags			Person
// @Accept			json
// @Produce		json
// @Param			country	path		string	true	"Country"
// @Success		200	{object}	[]domain.Person
// @Failure		500	{object}	handlers.ErrorResponce
// @Router			/person/{name} [get]
func (h *PersonHandler) GetByCountry(ctx *gin.Context) {
	req := ctx.Param("country")
	per, err := h.service.GetByCountry(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, nil)
		// handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}
