package main

import (
	"net/http"

	"github.com/alyumi/music_searcher/app"
	"github.com/labstack/echo/v4"
)

func main() {
	a := app.InitApp()
	e := echo.New()
	l := e.StdLogger
	e.GET("/", func(c echo.Context) error {
		l.Print("Hello!")
		return c.File("web/index.html")
	})

	e.GET("/getLinks", func(c echo.Context) error {

		URL := c.FormValue("link")
		l.Println("URL: ", URL)
		links := a.FindLinks(URL)
		ULTRA := links[0] + "\n" + links[1] + "\n" + links[2] + "\n" + links[3] + "\n" + links[4]

		return c.String(http.StatusOK, ULTRA)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
