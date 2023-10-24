package persons

import (
	"net/http"
	"person-predicator/internal/logger"
	"person-predicator/internal/server/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Summary		Get a person
//	@Description	Get a person from database
//	@Tags			Person
//	@Accept			json
//	@Produce		json
//	@Param			id		query		int		false	"Id"
//	@Param			name	query		string	false	"Name"
//	@Param			surname	query		string	false	"Surname"
//	@Param			age		query		int		false	"Age"
//	@Param			gender	query		string	false	"Gender"
//	@Param			country	query		string	false	"Country"
//	@Param			limit	query		int		false	"Max records count"
//	@Success		200		{object}	domain.Person
//	@Failure		400		{object}	handlers.ErrorResponce
//	@Failure		404
//	@Failure		500	{object}	handlers.ErrorResponce
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

	if n := intParse(ctx, idReq); n > 0 {
		filters["id"] = n
	}
	if n := intParse(ctx, ageReq); n != 0 {
		filters["age"] = n
	}
	if n := intParse(ctx, limitReq); n <= 0 {
		limit = -1
	} else {
		limit = n
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

	per, err := h.service.Get(ctx, filters, int(limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to get a person from database", Error: err})
		return
	}
	if len(per) == 0 {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	logger.Logger.Info("Person get")

	ctx.JSON(http.StatusOK, per)
}

func intParse(ctx *gin.Context, in string) int64 {
	if in == "" {
		return 0
	}
	n, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return 0
	}
	return n
}
