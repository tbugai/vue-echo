package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	_ = godotenv.Load()

	e := echo.New()
	e.Static("/assets", "public/assets")
	e.File("/", "public/index.html")

	serverAddr := os.ExpandEnv(":$BACKEND_PORT")
	e.Logger.Fatal(e.Start(serverAddr))
}
