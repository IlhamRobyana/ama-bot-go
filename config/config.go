package config

import (
	"fmt"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	err := loadENV()

	if err != nil {
		fmt.Println(err.Error())
	}
}
