package main

import (
	"fmt"

	module2 "github.com/ultrasad/gomodule"

	"github.com/labstack/gommon/color"
)

func main() {
	fmt.Println(color.Green("Test Grren Color"))
	fmt.Println(color.Blue(module2.Hello("Hanajung")))

	/*
		e := echo.New()
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World >> Echo!")
		})

		e.Logger.Fatal(e.Start(":1323"))
	*/
}
