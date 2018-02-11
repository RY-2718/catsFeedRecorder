package handler

import (
	"database/sql"
	"github.com/RY-2718/catsFoodRecorder/model"
	"github.com/labstack/echo"
	"net/http"
	"github.com/RY-2718/catsFoodRecorder/data"
	"strconv"
)

type H map[string]interface{}

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func GetFoods(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, model.GetFoods(db))
	}
}

func CreateFood(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		food := data.Food{}
		c.Bind(food)
		id, err := model.CreateFood(db, food.Name)
		if err != nil {
			return err
		}
		response := H{
			"created": id,
		}
		return c.JSON(http.StatusCreated, response)
	}
}

func DeleteFood(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := model.DeleteFood(db, id)
		if err != nil {
			return err
		}
		response := H{
			"deleted": id,
		}
		return c.JSON(http.StatusOK, response)
	}
}