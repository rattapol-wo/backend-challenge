package challenge2

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(c echo.Context) error {

	var reqBody requestKeyBoardEncoded

	// Bind JSON body struct
	if err := c.Bind(&reqBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	var response []responseKeyBoardDecoded
	for _, v := range reqBody.Encoded {
		resp := ProcessKeyboardDecode(v)
		response = append(response, resp)
	}

	return c.JSON(http.StatusOK, response)
}