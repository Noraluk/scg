package routes

import (
	"scg/api/controllers/doscg"

	"github.com/labstack/echo"
)

func Run() {
	e := echo.New()

	e.GET("/find-xyz", doscg.FindXYZ)
	e.GET("/find-bc", doscg.FindBC)
	e.GET("/find-direction", doscg.FindDirection)
	e.GET("/webhook", doscg.ReceiveLineMessage)
	e.Logger.Fatal(e.Start(":8080"))
}
