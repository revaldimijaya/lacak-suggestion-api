package httphandler

import (
	"context"
	"github/revaldimijaya/lacak-api/app/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HttpHandler) GetCitySuggestions(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*30)
	defer cancel()

	query := c.Query("q")
	if query == "" {
		c.JSON(
			http.StatusBadRequest,
			HttpResponse{
				ResultStatus: ResultStatus{
					Status:  "Error Validation",
					Message: "parameter q must be filled",
				},
				Data: nil,
			},
		)
		return
	}

	latStr := c.Query("latitude")
	lonStr := c.Query("longitude")

	resp, err := h.Usecase.GetCitySuggestions(ctx, usecase.RequestScoredCity{
		Query:  query,
		LatStr: latStr,
		LonStr: lonStr,
	})
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			HttpResponse{
				ResultStatus: ResultStatus{
					Status:  "Failed",
					Message: err.Error(),
				},
				Data: nil,
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		HttpResponse{
			ResultStatus: ResultStatus{
				Status:  "Success",
				Message: "Success",
			},
			Data: resp,
		},
	)
}
