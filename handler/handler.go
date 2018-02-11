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
		food := new(data.InnerFood)
		c.Bind(&food)
		result := model.CreateFood(db, food.Name)
		return c.JSON(http.StatusCreated, result)
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
			"id": id,
		}
		return c.JSON(http.StatusOK, response)
	}
}