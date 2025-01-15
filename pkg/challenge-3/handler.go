package challenge3

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context) error {

	// Bind JSON body struct
	// meat-and-filler
	value := c.QueryParam("typeMeat")
	if value == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "message": "Missing typeMeat query parameter",
        })
    }

	response, err := CountMeat(value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Process count meat failed",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"beef": response,
	})
}
