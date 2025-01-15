package challenge

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context) error {
	var reqBody requestNumberSummary

	// Bind JSON body struct
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	response := SumValue(reqBody.ListNumber)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"result": response,
	})
}
