package routes

import (
	"myapp/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	type Message struct {
		Message string `json:"hello world"`
	}

	e.GET("/", func(c echo.Context) error {
		res := &Message{
			Message: "hello world",
		}
		return c.JSON(200, res)
	})

	e.GET("/pegawai", controllers.Index)
	e.POST("/pegawai", controllers.Store)
	e.PUT("/pegawai", controllers.Update)
	e.DELETE("/pegawai", controllers.Delete)

	return e
}
