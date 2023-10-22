package persons

import (
	"net/http"
	"person-predicator/internal/server/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Summary		Get a person
//	@Description	Get a person from database
//	@Tags			Person
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int		true	"Id"
//	@Param			name	path		string	true	"Name"
//	@Param			surname	path		string	true	"Surname"
//	@Param			age		path		int		true	"Age"
//	@Param			gender	path		int		true	"Gender"
//	@Param			country	path		int		true	"Country"
//	@Param			limit	path		int		true	"Max records count"
//	@Success		200		{object}	domain.Person
//	@Failure		400		{object}	handlers.ErrorResponce
//	@Failure		500		{object}	handlers.ErrorResponce
//	@Router			/person [get]
func (h *PersonHandler) Get(ctx *gin.Context) {
	filters := make(map[string]interface{})
	var err error
	idReq := ctx.Query("id")
	name := ctx.Query("name")
	surname := ctx.Query("surname")
	ageReq := ctx.Query("age")
	gender := ctx.Query("gender")
	country := ctx.Query("country")
	limitReq := ctx.Query("limit")
	var limit int64

	if idReq != "" {
		id, err := strconv.ParseInt(idReq, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest,
				handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
			return
		}
		filters["id"] = id
	}
	if ageReq != "" {
		age, err := strconv.ParseInt(ageReq, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest,
				handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
			return
		}
		filters["age"] = age
	}
	if limitReq != "" {
		limit, err = strconv.ParseInt(limitReq, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusBadRequest,
				handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
			return
		}
	} else {
		limit = 10
	}
	if name != "" {
		filters["name"] = name
	}
	if surname != "" {
		filters["surname"] = surname
	}
	if gender != "" {
		filters["gender"] = gender
	}
	if country != "" {
		filters["country"] = country
	}

	per, err := h.service.Get(filters, int(limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	ctx.JSON(http.StatusOK, per)
}
