package persons

import (
	"net/http"
	"person-predicator/internal/logger"
	"person-predicator/internal/server/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	@Summary		Delete a person
//	@Description	Delete a person from database
//	@Tags			Person
//	@Accept			json
//	@Produce		json
//	@Param			id	query	int	true	"Person Id"
//	@Success		200
//	@Failure		400	{object}	handlers.ErrorResponce
//	@Failure		500	{object}	handlers.ErrorResponce
//	@Router			/person/{id} [delete]
func (h *PersonHandler) Delete(ctx *gin.Context) {
	req := ctx.Query("id")
	id, err := strconv.ParseInt(req, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Failed to parse request", Error: err})
		return
	}
	if id <= 0 {
		ctx.JSON(http.StatusBadRequest,
			handlers.ErrorResponce{Message: "Incorrect ID", Error: err})
		return
	}
	err = h.service.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			handlers.ErrorResponce{Message: "Failed to delete a person from database", Error: err})
		return
	}
	logger.Logger.Info("Person deleted")
	ctx.Status(http.StatusOK)
}
