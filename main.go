package main

import (
	"net/http"

	t "website-buider/templates"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	renderer := t.NewTemplate()

	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	pb := e.Group("pb") // page builder
	pb.GET("/edit", BuilderPage)

	e.Logger.Fatal(e.Start(":1323"))
}

func BuilderPage(c echo.Context) error {
	return c.Render(http.StatusOK, "builder", "World")
}
