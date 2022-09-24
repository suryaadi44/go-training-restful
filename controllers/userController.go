package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/suryaadi44/go-training-restful/lib/database"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   users,
	})
}
