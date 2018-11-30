package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/assets", "public/assets")
	e.File("/", "public/index.html")
	e.Logger.Fatal(e.Start(":8000"))
}
